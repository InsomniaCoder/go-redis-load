package server

import (
	"fmt"
	"github.com/InsomniaCoder/go-redis-load/config"
)

func Init() {
	config := config.GetConfig()
	r := NewRouter()
	portNumber := fmt.Sprintf(":%d", config.Server.Port)
	r.Run(portNumber)
}
