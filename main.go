package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/InsomniaCoder/go-redis-load/server"
	"github.com/InsomniaCoder/go-redis-load/config"
)

func main() {
	environment := flag.String("e", "dev", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	server.Init()
}
