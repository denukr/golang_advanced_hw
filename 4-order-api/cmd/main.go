package main

import (
	"golang-adv/4-order-api/configs"
	"golang-adv/4-order-api/internal/products"
	"golang-adv/4-order-api/pkg/db"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	cfg := configs.LoadConfig()

	db := db.NewDb(cfg)

	productRepo := products.NewProductRepository(db)

	products.NewProductHandler(router, &products.ProductHandlerDeps{
		Configs:     cfg,
		ProductRepo: productRepo,
	})

	server := http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	log.Println("Server is starting on port 8081")
	server.ListenAndServe()
}
