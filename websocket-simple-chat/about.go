package main

import(
    "github.com/martini-contrib/render"
)

func About(r render.Render) {
    r.HTML(200, "about", nil)
}
