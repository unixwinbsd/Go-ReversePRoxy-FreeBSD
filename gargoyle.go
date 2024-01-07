package main

import (
	"os"

	"github.com/unixwinbsd/Go-ReversePRoxy-FreeBSD/internal/server"
)

func main() {
	var configFile string
	if len(os.Args) > 1 {
		configFile = os.Args[1]
	}
	if configFile == "" {
		configFile = "./config.json"
	}
	server.Start(configFile)
}
