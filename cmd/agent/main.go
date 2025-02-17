package main

import (
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"
	"reflect"
	"runtime"
	"time"
)

const (
	pollInterval   = 2 * time.Second
	reportInterval = 10 * time.Second
)

type gauge float64
type counter int64

type Metrics struct {
	Alloc         gauge
	BuckHashSys   gauge
	Frees         gauge
	GCCPUFraction gauge
	GCSys         gauge
	HeapAlloc     gauge
	HeapIdle      gauge
	HeapInuse     gauge
	HeapObjects   gauge
	HeapReleased  gauge
	HeapSys       gauge
	LastGC        gauge
	Lookups       gauge
	MCacheInuse   gauge
	MCacheSys     gauge
	MSpanInuse    gauge
	MSpanSys      gauge
	Mallocs       gauge
	NextGC        gauge
	NumForcedGC   gauge
	NumGC         gauge
	OtherSys      gauge
	PauseTotalNs  gauge
	StackInuse    gauge
	StackSys      gauge
	Sys           gauge
	TotalAlloc    gauge
	PollCount     counter
	RandomValue   gauge
}

func newMetrics() *Metrics {
	return &Metrics{}
}

func (m *Metrics) update() {
	mem := &runtime.MemStats{}
	runtime.ReadMemStats(mem)

	m.Alloc = gauge(mem.Alloc)
	m.BuckHashSys = gauge(mem.BuckHashSys)
	m.Frees = gauge(mem.Frees)
	m.GCCPUFraction = gauge(mem.GCCPUFraction)
	m.GCSys = gauge(mem.GCSys)
	m.HeapAlloc = gauge(mem.HeapAlloc)
	m.HeapIdle = gauge(mem.HeapIdle)
	m.HeapInuse = gauge(mem.HeapInuse)
	m.HeapObjects = gauge(mem.HeapObjects)
	m.HeapReleased = gauge(mem.HeapReleased)
	m.HeapSys = gauge(mem.HeapSys)
	m.LastGC = gauge(mem.LastGC)
	m.Lookups = gauge(mem.Lookups)
	m.MCacheInuse = gauge(mem.MCacheInuse)
	m.MCacheSys = gauge(mem.MCacheSys)
	m.MSpanInuse = gauge(mem.MSpanInuse)
	m.MSpanSys = gauge(mem.MSpanSys)
	m.Mallocs = gauge(mem.Mallocs)
	m.NextGC = gauge(mem.NextGC)
	m.NumForcedGC = gauge(mem.NumForcedGC)
	m.NumGC = gauge(mem.NumGC)
	m.OtherSys = gauge(mem.OtherSys)
	m.PauseTotalNs = gauge(mem.PauseTotalNs)
	m.StackInuse = gauge(mem.StackInuse)
	m.StackSys = gauge(mem.StackSys)
	m.Sys = gauge(mem.Sys)
	m.TotalAlloc = gauge(mem.TotalAlloc)
	m.PollCount++
	m.RandomValue = gauge(rand.Float64())
}

func prepapeForReports(metrics *Metrics) {
	value := reflect.ValueOf(metrics).Elem()
	typeOfMetrics := reflect.TypeOf(*metrics)

	for i := 0; i < value.NumField(); i++ {
		metricType := value.Field(i).Type().Name()
		metricName := typeOfMetrics.Field(i).Name
		metricValue := fmt.Sprint(value.Field(i))

		sendToServer(metricType, metricName, metricValue)
	}
}

func sendToServer(metricType string, metricName string, metricValue string) error {
	url := fmt.Sprintf("http://localhost:8080/update/%s/%s/%s", metricType, metricName, metricValue)
	res, err := http.Post(url, "text/plain", nil)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	_, err = io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	metrics := newMetrics()

	pollTicker := time.Tick(pollInterval)
	reportTicker := time.Tick(reportInterval)

	for {
		select {
		case <-pollTicker:
			metrics.update()
		case <-reportTicker:
			prepapeForReports(metrics)
		default:
		}
	}
}
