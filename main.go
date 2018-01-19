package main

import (
	"flag"
	"log"
	"net/http"
	"strings"
)

const defaultPort = ":5000"

func main() {
	portFlag := flag.String("p", "", "port number")
	flag.Parse()
	var port string
	if *portFlag == "" {
		log.Println("no port specified, using default port ", defaultPort)
		port = defaultPort
	} else {
		port = *portFlag
		if !strings.HasPrefix(port, ":") {
			port = ":" + port
		}
	}
	log.Println("listening on port ", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
