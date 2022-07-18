package main

import (
	"fmt"
	"github.com/marufmax/goweb/pkg/config"
	"github.com/marufmax/goweb/pkg/handlers"
	"github.com/marufmax/goweb/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":9095"

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Can not create template cache", err)
	}

	app.TemplateCache = tc

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port: http://localhost%s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
