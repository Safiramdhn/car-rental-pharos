package main

import (
	"car-rental/infra"
	"car-rental/routes"
	"flag"
	"log"
)

func main() {
	ctx, err := infra.NewServiceContext()
	if err != nil {
		log.Fatal("can't init service context %w", err)
	}

	if shouldNotLaunchServer() {
		return
	}

	routes.NewRoutes(*ctx)
}

func shouldNotLaunchServer() bool {
	shouldNotLaunch := false

	flag.Parse()
	flag.Visit(func(f *flag.Flag) {
		if f.Name == "m" {
			shouldNotLaunch = true
		}

		if f.Name == "s" {
			shouldNotLaunch = true
		}
	})

	return shouldNotLaunch
}
