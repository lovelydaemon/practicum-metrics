package main

import "github.com/lovelydaemon/practicum-metrics/internal/server/app"

func main() {
	if err := app.Run(); err != nil {
		panic(err)
	}
}
