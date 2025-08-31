package main

import (
	"distributed-file-system/internal/routes"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	port = flag.Int("port", 8080, "port for the server endpoint")
)

func main() {
	flag.Parse()
	fmt.Printf("http://localhost:%v\n", *port)

	server := http.Server{
		Addr:         fmt.Sprintf(":%v", *port),
		ReadTimeout:  time.Second * 30,
		WriteTimeout: time.Second * 30,
		Handler:      routes.Routes(),
	}

	log.Fatal(server.ListenAndServe())
}
