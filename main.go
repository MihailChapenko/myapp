package main

import (
	"github.com/MihailChapenko/celeritas"
	"github.com/MihailChapenko/myapp/handlers"
)

type application struct {
	App      *celeritas.Celeritas
	Handlers *handlers.Handlers
}

func main() {
	c := initApplication()
	c.App.ListenAndServe()
}
