package main

import (
	"context"
	"log"

	"github.com/sumukhj1219/go-profiler/internals/config"
	"github.com/sumukhj1219/go-profiler/internals/server"
)

func main() {
	ctx := context.Background()
	cfg, err := config.Load()

	if err != nil {
		log.Fatal("unable to load config variables")
		return
	}
	
	server.Start(*cfg, ctx)
}
