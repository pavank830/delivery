package main

import (
	"flag"
	"fmt"

	"github.com/pavank830/delivery"
)

func main() {
	var serverPortFlag = flag.String("port", "8080", "HTTP server port,default port 8080")
	flag.Parse()
	fmt.Println("HTTP server port value:", *serverPortFlag)
	delivery.Start(*serverPortFlag)
}
