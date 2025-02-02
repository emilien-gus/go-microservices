package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	handler "example.com/go-demo-1/product-api/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	h := handler.NewProduct(l)
	sm := http.NewServeMux()
	sm.Handle("/", h)

	ser := http.Server{
		Addr:    ":8080",
		Handler: sm,
	}

	go func() {
		err := ser.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Kill)
	signal.Notify(sigChan, os.Interrupt)

	sig := <-sigChan
	l.Println("Recieved terminate, graceful shutdwon: ", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	ser.Shutdown(ctx)
}
