package main

import (
	"log"
	"net/http"
	"os"

	handler "example.com/go-demo-1/product-api/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	sm := handler.NewProduct(l)

	ser := http.Server{
		Addr:    "9090",
		Handler: sm,
	}
}
