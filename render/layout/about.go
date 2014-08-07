package layout

import(
    "github.com/martini-contrib/render"
)

type Profile struct {
    Name string
    Skill []string
}

type AboutViewModel struct {
    Title string
    Profile Profile
}

func AboutRender(r render.Render) {

    skill := []string{"TypeScript", "Sass/Compass", "Go"}

    profile := Profile{
        "hogeyan",
        skill,
    }

    viewModel := AboutViewModel{
        "About me",
        profile,
    }

    r.HTML(200, "about", viewModel)
}
