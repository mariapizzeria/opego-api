package main

import (
	"log"
	"net/http"

	"github.com/mariapizzeria/opego-api/internal/order"
	"github.com/mariapizzeria/opego-api/pkg/configs"
	"github.com/mariapizzeria/opego-api/pkg/db"
	"github.com/mariapizzeria/opego-api/pkg/middleware"
)

func main() {
	config := configs.LoadConfig()
	newDB := db.NewDb(config)
	router := http.NewServeMux()
	//repositories
	orderRepository := order.NewRepository(newDB)
	//handlers
	order.NewHandler(router, order.HandlerDeps{
		Repository: orderRepository,
	})
	server := &http.Server{
		Addr:    ":8080",
		Handler: middleware.Logging(router),
	}
	log.Println("Server is listening")
	server.ListenAndServe()
}
