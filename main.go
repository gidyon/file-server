package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	port = flag.String("port", ":8080", "Port to serve file content")
	dir  = flag.String("dir", "", "Directory to serve files")
)

func main() {
	flag.Parse()
	log.Printf("Http file server started on port: %s\n", *port)
	log.Fatalln(http.ListenAndServe(*port, http.FileServer(http.Dir(*dir))))
}
