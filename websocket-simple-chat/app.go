package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func main() {

	initDB()

	m := martini.Classic()

	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))
	m.Use(martini.Static("assets"))

	m.Get("/", Index)
	m.Get("/about", About)
	m.Get("/ws", WebSocket)
	m.Delete("/ws/:id", DeleteEntry)

	m.Run()
}
