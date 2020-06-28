package main

import (
	"github.com/go-chi/chi"
	//v3 "github.com/go-chi/chi/_examples/versions/presenter/v3"
	"net/http"
)

func main(){

	r := chi.NewRouter()
	r.Get("/", performTest)

	http.ListenAndServe(":3333", r)

}

func performTest(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("welcome"))
}