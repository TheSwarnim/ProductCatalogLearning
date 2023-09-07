package main

import (
	"context"
	"go_microservice_learning_1/handlers"
	"go_microservice_learning_1/models"
	"go_microservice_learning_1/respository"
	"go_microservice_learning_1/services"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	productRepository := respository.NewProductRepository(productList)
	productService := services.NewProductService(productRepository)
	productHandler := handlers.NewProductHandler(productService, l)

	serveMux := http.NewServeMux()

	serveMux.Handle("/products", productHandler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      serveMux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			l.Println(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := server.Shutdown(tc)
	if err != nil {
		return
	}
}

var productList = []*models.Product{
	{
		Id:          1,
		Name:        "Toy Car",
		Description: "Classic red toy car",
		Quantity:    50,
	},
	{
		Id:          2,
		Name:        "Book",
		Description: "Hardcover novel",
		Quantity:    120,
	},
	{
		Id:          3,
		Name:        "Smartphone",
		Description: "Latest model smartphone",
		Quantity:    20,
	},
	{
		Id:          4,
		Name:        "Laptop",
		Description: "Powerful work laptop",
		Quantity:    15,
	},
	{
		Id:          5,
		Name:        "Board Game",
		Description: "Fun family board game",
		Quantity:    30,
	},
}
