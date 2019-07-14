package main

import (
	"net/http"
	"os"

	"myfunc"

	"github.com/gorilla/mux"
)

func Main() {
	r := mux.NewRouter()
	r.Handle("/", http.FileServer(http.Dir("./src/web")))
	myfunc.Println("Start server at " + os.Getenv("HTTP_PORT"))
	http.ListenAndServe(":"+os.Getenv("HTTP_PORT"), r)
}
