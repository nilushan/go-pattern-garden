package options

import (
	"fmt"
	"time"
)

type Server struct {
	host     string
	port     int
	timeout  time.Duration
	maxConns int
	tls      bool
}

func (s *Server) String() string {
	return fmt.Sprintf(
		"Server Configuration:\n"+
			"\tHost: %s\n"+
			"\tPort: %d\n"+
			"\tTLS Enabled: %v\n"+
			"\tTimeout: %s\n"+
			"\tMax Connections: %d",
		s.host, s.port, s.tls, s.timeout, s.maxConns,
	)
}

// Option function type
type Option func(*Server)

// Individual option functions
func WithPort(port int) Option {
	return func(s *Server) {
		s.port = port
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.timeout = timeout
	}
}

func WithTLS(enabled bool) Option {
	return func(s *Server) {
		s.tls = enabled
	}
}

func WithMaxConnections(max int) Option {
	return func(s *Server) {
		s.maxConns = max
	}
}

// Constructor accepts variadic options
func NewServer(host string, opts ...Option) *Server {
	server := &Server{
		host:     host,
		port:     8080, // defaults
		timeout:  30 * time.Second,
		maxConns: 100,
		tls:      false,
	}

	// Apply all options
	for _, opt := range opts {
		opt(server)
	}

	return server
}

func Run() {
	server1 := NewServer("api.example.com",
		WithMaxConnections(20),
		WithPort(9000),
		WithTimeout(60*time.Second),
		WithTLS(true),
	)
	fmt.Printf("Server1: %v\n", server1)

	server2 := NewServer("api.example.com", WithMaxConnections(10), WithPort(3000), WithTLS(true))
	fmt.Printf("Server2: %v\n", server2)

	server3 := NewServer("api.example.com", WithMaxConnections(10), WithPort(3000), WithTLS(true), WithTimeout(10*time.Second))
	fmt.Printf("Server3: %v\n", server3)

}
