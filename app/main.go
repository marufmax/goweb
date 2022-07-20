package main

import "github.com/marufmax/larago"

type application struct {
	App *larago.Larago
}

func main() {
	l := initApplication()
	l.App.ListenAndServe()
}
