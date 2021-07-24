package server

import (
	"testing"

	"github.com/gorilla/mux"
)

func TestNewServer(t *testing.T) {
	t.Run("Create a new server", func(t *testing.T) {
		s := NewServer(&mux.Router{})

		if s == nil {
			t.Errorf("Could not create a new server.")
		}
	})
}
