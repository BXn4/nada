package server

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type SocketType int

const (
	Socket SocketType = iota
	WebSocket
)

type ServerConfig struct {
	Host   string
	Port   string
	Socket SocketType
}

type Server struct {
	config  *ServerConfig
	maxConn int
}

func New(config *ServerConfig) (*Server, error) {
	return &Server{
		config: config,
	}, nil
}

func (s *Server) Run() error {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	address := fmt.Sprintf("%s:%s", s.config.Host, s.config.Port)
	switch s.config.Socket {
	case Socket:
		return s.ListenSocket(address)
	case WebSocket:
		return s.ListenWebSocket(address)
	default:
		return fmt.Errorf("Unknown socket type: %d", s.config.Socket)
	}
}

func (s *Server) ListenSocket(address string) error {
	return nil
}

func (s *Server) ListenWebSocket(address string) error {
	return nil
}
