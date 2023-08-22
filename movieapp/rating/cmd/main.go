package main

import (
	"log"
	"net/http"

	"github.com/ullasasindhur/microservices_with_go/movieapp/rating/internal/controller/rating"
	httphandler "github.com/ullasasindhur/microservices_with_go/movieapp/rating/internal/handler"
	"github.com/ullasasindhur/microservices_with_go/movieapp/rating/internal/repository/memory"
)

func main() {
	log.Println("Starting the  rating service")
	repo := memory.New()
	ctrl := rating.New(repo)
	h := httphandler.New(ctrl)
	http.Handle("/rating", http.HandlerFunc(h.Handle))
	if err := http.ListenAndServe(":8082", nil); err != nil {
		panic(err)
	}
}
