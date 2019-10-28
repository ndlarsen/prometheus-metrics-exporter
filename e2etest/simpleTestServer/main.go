package main

import (
	"flag"
	"os"
	. "simpleTestServer/server"
)

func main() {

	port := flag.String("port", "", "the port the server will listen on")

	flag.Parse()

	if *port == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	Server(port)
}