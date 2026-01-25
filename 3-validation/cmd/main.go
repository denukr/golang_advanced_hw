package main

import (
	"go/email-verify/config"
	"go/email-verify/internal/verify"
	"log"
	"net/http"
	"strings"
)

func main() {
	cfg := config.Load()
	router := http.NewServeMux()

	verify.NewVerifyHandler(router, &verify.VerifyHandlerDeps{
		Config: cfg.Verify,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	log.Println("Server is running on port", strings.Split(server.Addr, ":")[1])
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Error occured! ", err)
	}
}
