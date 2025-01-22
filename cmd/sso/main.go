package main

import (
	"fmt"

	"github.com/carbon77/sso/internal/config"
)

func main() {
	config := config.MustLoad()
	fmt.Println(config)

	// TODO: init logger

	// TODO: init app

	// TODO: start grpc-server
}
