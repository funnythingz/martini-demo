package layout

import(
    "github.com/martini-contrib/render"
)

type AboutViewModel struct {
    Title string
    Description string
}

func AboutRender(r render.Render) {

    viewModel := AboutViewModel{
        "About",
        "description",
    }

    r.HTML(200, "about", viewModel)
}
