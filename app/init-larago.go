package main

import (
	"github.com/marufmax/larago"
	"log"
	"os"
)

func initApplication() *application {
	path, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	lar := &larago.Larago{
		AppName: "My AppLiCaTiOn",
		Debug:   true,
		Version: "1.0.0",
	}
	err = lar.New(path)

	if err != nil {
		log.Fatal(err)
	}

	app := &application{
		App: lar,
	}

	return app
}
