package main

import(
    "github.com/go-martini/martini"
    "github.com/martini-contrib/render"
)

func main() {

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

    m.Get("/", IndexRender)
    m.Get("/about", AboutRender)

    m.Run()

}
