package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ant0ine/go-json-rest/rest"

	"main/kokodayo"
)

func hello(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteJson(kokodayo.GenHello(id))
	fmt.Println("Hello called!")
}

const host = "localhost"
const port = 8080
const staticPath = "./static"

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/api/hello/:id", hello),
	)
	if err != nil {
		log.Fatal(err)
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%v:%v", host, port), api.MakeHandler()))
}
