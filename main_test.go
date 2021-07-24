package main

import (
	"testing"

	"github.com/gorilla/mux"
)

func TestMakeRouter(t *testing.T) {
	t.Run("Make Router", func(t *testing.T) {
		router := mux.NewRouter()
		if router == nil {
			t.Errorf("Could not create new mux.Router")
		}
	})
}
