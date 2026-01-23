package main

import "net/http"

func main() {

	router := http.NewServeMux()

	NewRandomVarHandler(router)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe()
}
