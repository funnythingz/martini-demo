package layout

import(
    "github.com/martini-contrib/render"
)

type IndexViewModel struct {
    Title string
    Description string
}

func Index(r render.Render) {

    viewModel := IndexViewModel{
        "Index",
        "description",
    }

    r.HTML(200, "index", viewModel)
}
