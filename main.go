package main

import (
	. "github.com/Alynt/alynt/app"
	. "github.com/Alynt/alynt/config"
)

func main() {
	config := GetConfig()

	app := &App{}
	app.Initialize(config)
	app.Run()
}
