package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	host = flag.String("host", "127.0.0.1", "http server host")
	port = flag.Uint("port", 8080, "http server port")
	dir  = flag.String("dir", ".", "directory to serve")
)

func main() {
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*dir)))

	hostport := fmt.Sprintf("%s:%d", *host, *port)
	log.Printf("serving directory %q on %q ...\n", *dir, hostport)

	if err := http.ListenAndServe(hostport, nil); err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}
}
