package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	var dir string

	flag.StringVar(&dir, "dir", "./www", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()
	r := mux.NewRouter()

	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(dir))))

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalln("There's an error with the server", err)
	}
}
