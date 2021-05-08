package main

import (
	"workerpool-application/internal/app"
)

func main() {
	app := app.NewApp()
	app.Run()
}
