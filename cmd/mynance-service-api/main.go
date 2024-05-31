package main

import "github.com/iacopoghilardi/mynance-service-api/internal/app"

func main() {
	err := app.InitApp()
	if err != nil {
		panic(err)
	}
}
