package server

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/dgethings/chunter/internal/document"
	"github.com/dgethings/chunter/internal/features"
)

type serverState int

const (
	stateCreated serverState = iota
	stateInitialized
	stateShutDown
)

type Server struct {
	documents *document.Store
	features  *features.Router
	logger    *log.Logger

	stateMu sync.Mutex
	state   serverState
	version string
}

func New(version string) *Server {
	return &Server{
		documents: document.NewStore(),
		features:  features.NewRouter(),
		logger:    log.New(os.Stderr, "[chunter] ", log.Ldate|log.Ltime|log.Lshortfile),
		version:   version,
	}
}

func (s *Server) RegisterFeature(f features.Feature) {
	s.features.Register(f)
}

func (s *Server) Logger() *log.Logger {
	return s.logger
}

func (s *Server) setState(newState serverState) error {
	s.stateMu.Lock()
	defer s.stateMu.Unlock()
	if newState == stateShutDown && s.state != stateInitialized {
		return fmt.Errorf("server not initialized")
	}
	s.state = newState
	return nil
}

func (s *Server) getState() serverState {
	s.stateMu.Lock()
	defer s.stateMu.Unlock()
	return s.state
}
