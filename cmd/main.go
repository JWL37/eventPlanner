package main

import (
	"eventPlanner/internal/app"
	"eventPlanner/internal/config"
)

func main() {

	cfg := config.Config{}
	config.ConfigInit(&cfg)
	app.RunApp(cfg)
}
