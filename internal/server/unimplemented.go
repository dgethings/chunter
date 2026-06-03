package server

import (
	"context"
	"fmt"
)

func (s *Server) Unimplemented(ctx context.Context) error {
	return fmt.Errorf("method not implemented")
}
