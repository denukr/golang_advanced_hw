package main

import (
	"fmt"
	"math/rand/v2"
	"net/http"
)

type RandomVarHandler struct{}

func (h *RandomVarHandler) getRandomVar() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		randVar := byte(rand.IntN(6) + 1)
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "%d", randVar)
	}
}

func NewRandomVarHandler(r *http.ServeMux) {
	h := RandomVarHandler{}
	r.HandleFunc("GET /rand", h.getRandomVar())
}
