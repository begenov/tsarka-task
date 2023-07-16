package main

import (
	"fmt"
	"log"

	"github.com/begenov/tsarka-task/internal/app"
	"github.com/begenov/tsarka-task/internal/config"
)

func main() {
	cfg := config.NewConfig()

	fmt.Println(cfg)
	if err := app.Run(cfg); err != nil {
		log.Fatalf("error: application run \t %v", err)
	}
}
