package main

import (
	"log"
	"net/http"

	metadatagateway "github.com/ullasasindhur/microservices_with_go/movieapp/movie/gateway/metadata/http"
	ratinggateway "github.com/ullasasindhur/microservices_with_go/movieapp/movie/gateway/rating/http"
	"github.com/ullasasindhur/microservices_with_go/movieapp/movie/internal/controller/movie"
	httphandler "github.com/ullasasindhur/microservices_with_go/movieapp/movie/internal/handler"
)

func main() {
	log.Println("Starting  the movie service")
	metadataGateway := metadatagateway.New("localhost:8081")
	ratingGateway := ratinggateway.New("localhost:8082")
	ctrl := movie.New(ratingGateway, metadataGateway)
	h := httphandler.New(ctrl)
	http.Handle("/movie", http.HandlerFunc(h.GetMovieDetails))
	if err := http.ListenAndServe(":8083", nil); err != nil {
		panic(err)
	}

}
