// Package handlers wires HTTP requests to the greeting logic.
package handlers

import (
	"encoding/json"
	"net/http"

	"greeter/internal/greeting"
)

// Hello handles GET /hello?name=X, returning a JSON Greeting.
func Hello(w http.ResponseWriter, r *http.Request) {
	g := greeting.For(r.URL.Query().Get("name"))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(g)
}
