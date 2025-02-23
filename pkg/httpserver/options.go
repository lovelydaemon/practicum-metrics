package httpserver

import (
	"net"
	"time"
)

type Option func(s *Server)

func Port(port string) Option {
	return func(s *Server) {
		s.server.Addr = net.JoinHostPort("", port)
	}
}

func ReadTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.server.ReadTimeout = timeout
	}
}

func ReadHeaderTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.server.ReadHeaderTimeout = timeout
	}
}

func WriteTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.server.WriteTimeout = timeout
	}
}

func IdleTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.server.IdleTimeout = timeout
	}
}

func MaxHeaderBytes(maxHeaderBytes int) Option {
	return func(s *Server) {
		s.server.MaxHeaderBytes = maxHeaderBytes
	}
}

func ShutdownTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.shutdownTimeout = timeout
	}
}
