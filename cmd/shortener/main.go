package main

import "github.com/Makovey/shortener/internal/app"

func main() {
	a := app.NewApp()
	if err := a.Run(); err != nil {
		panic(err)
	}
}
