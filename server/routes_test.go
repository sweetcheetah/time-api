package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestRootPath(t *testing.T) {
	t.Run("Root Path", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/", nil)
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()
		handler := http.HandlerFunc(handlerRoot)
		handler.ServeHTTP(recorder, req)

		if status := recorder.Code; status != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
		}
	})
}

func TestLocalTimePath(t *testing.T) {
	t.Run("localtime Path", func(t *testing.T) {
		router := mux.NewRouter()
		srv := NewServer(router)

		req, err := http.NewRequest(http.MethodGet, "/localtime", nil)
		if err != nil {
			t.Fatal(err)
		}

		recorder := httptest.NewRecorder()
		srv.ServeHTTP(recorder, req)

		if status := recorder.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}
	})

}
