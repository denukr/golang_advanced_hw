package main

import (
	"golang-adv/4-order-api/configs"
	"golang-adv/4-order-api/internal/products"
	"golang-adv/4-order-api/pkg/db"
	"golang-adv/4-order-api/pkg/middleware"
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

	mw := middleware.NewLogMiddleware()

	server := http.Server{
		Addr:    ":8081",
		Handler: mw.Log(router),
	}

	log.Println("Server is starting on port 8081")
	server.ListenAndServe()
}
