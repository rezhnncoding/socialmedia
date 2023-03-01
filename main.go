package main

import (
	"fmt"
	"github.com/config"
	"github.com/gorilla/mux"
	"github.com/router"
	"log"
	"net/http"
)

func main() {
	config.Load()
	startServer(router.Build(), config.Port)
}

func startServer(router *mux.Router, port int) {
	fmt.Printf("Starting server at port [%d].\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
