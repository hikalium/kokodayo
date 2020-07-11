package main

import (
	"fmt"
	"net/http"

	"goji.io"
	"goji.io/pat"

	"main/kokodayo"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, kokodayo.GenHello(pat.Param(r, "name")))
}

func main() {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/hello/:name"), hello)

	http.ListenAndServe("localhost:8000", mux)
}
