package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"goji.io"
	"goji.io/pat"

	"main/kokodayo"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, kokodayo.GenHello(pat.Param(r, "name")))
}

const host = "localhost"
const port = 8080
const staticPath = "./static"

func main() {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/hello/:name"), hello)
	path, err := filepath.Abs(staticPath)
	if err != nil {
		log.Fatal(err)
	}
	mux.Handle(pat.Get("/*"), http.FileServer(http.Dir(path)))

	fmt.Printf("Listening %v:%v...\n", host, port)
	http.ListenAndServe(fmt.Sprintf("%v:%v", host, port), mux)
}
