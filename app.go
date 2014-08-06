package main

import(
    "github.com/go-martini/martini"
    "github.com/martini-contrib/render"
    "./Layout"
)

func app() {

    m := martini.Classic()

    m.Use(render.Renderer(render.Options{
        Directory: "templates",
        Layout: "layout",
        Extensions: []string{".tmpl"},
        Charset: "utf-8",
    }))

    m.NotFound(func (r render.Render){
        r.Redirect("/")
    })

    m.Get("/", layout.Index)

    m.Run()
}

func main() {
    app()
}
