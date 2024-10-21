package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/leinodev/go-microservice-template/config"
	allInOne "github.com/leinodev/go-microservice-template/internal/app/all-in-one"
)

func main() {
	ctx := context.Background()

	var cfgPath string
	flag.StringVar(&cfgPath, "c", "./config/config.yaml", "path to configuration file")
	flag.Parse()

	c, err := config.NewConfig(cfgPath)
	panicOnErr(err)

	app := allInOne.NewApp(c)

	if err := app.Configure(ctx); err != nil {
		panicOnErr(fmt.Errorf("cannot configure app: %w", err))
	}

	if err := app.Run(ctx); err != nil {
		panicOnErr(fmt.Errorf("cannot run app: %w", err))
	}
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
