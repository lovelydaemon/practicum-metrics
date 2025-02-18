package httpserver

import (
	"context"
	"net/http"
	"time"
)

const (
	_defaultAddr              = ":8080"
	_defaultReadTimeout       = 10 * time.Second
	_defaultReadHeaderTimeout = 1 * time.Second
	_defaultWriteTimeout      = 10 * time.Second
	_defaultIdleTimeout       = 10 * time.Second
	_defaultMaxHeaderBytes    = 1 << 20 // 1MB
	_defaultShutdownTimeout   = 3 * time.Second
)

type Server struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

func New(handler http.Handler) *Server {
	httpServer := &http.Server{
		Addr:              _defaultAddr,
		Handler:           handler,
		ReadTimeout:       _defaultReadTimeout,
		ReadHeaderTimeout: _defaultReadHeaderTimeout,
		WriteTimeout:      _defaultWriteTimeout,
		IdleTimeout:       _defaultIdleTimeout,
		MaxHeaderBytes:    _defaultMaxHeaderBytes,
	}

	s := &Server{
		server:          httpServer,
		notify:          make(chan error, 1),
		shutdownTimeout: _defaultShutdownTimeout,
	}

	s.start()

	return s
}

func (s *Server) start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
