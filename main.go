package main

import (
	"banking/internal/channel/rest"
	"banking/internal/config"
)

func main() {
	config.StartConfig()
	rest.NewServer()

}
