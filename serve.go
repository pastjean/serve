package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	port = flag.Int("port", 1337, "The http listening port")
)

func main() {
	flag.Parse()

	dir, err := os.Getwd()
	if err != nil {
		log.Println("Error", err)
	}

	fs := http.FileServer(http.Dir(dir))
	http.Handle("/", fs)

	log.Println(fmt.Sprintf("Listening on http://localhost:%d", *port))
	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}
