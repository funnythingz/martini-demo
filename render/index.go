package main

import(
    "github.com/martini-contrib/render"
)

type IndexViewModel struct {
    Title string
    Description string
}

func IndexRender(r render.Render) {

    viewModel := IndexViewModel{
        "Martini Demo",
        "Description",
    }

    r.HTML(200, "index", viewModel)
}
