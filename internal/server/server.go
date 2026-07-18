package server

import (
	"fmt"
	"sync"

	"github.com/dgethings/chunter/internal/document"
	"github.com/dgethings/chunter/internal/features"
	"github.com/dgethings/chunter/internal/features/cisco_ios_jinja2"
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

	stateMu sync.Mutex
	state   serverState
	version string
}

func New(version string) *Server {
	s := &Server{
		documents: document.NewStore(),
		features:  features.NewRouter(),
		version:   version,
	}
	s.RegisterFeature(cisco_ios_jinja2.New())
	return s
}

func (s *Server) RegisterFeature(f features.Feature) {
	s.features.Register(f)
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
