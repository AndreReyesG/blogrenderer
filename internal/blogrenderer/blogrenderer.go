package blogrenderer

import (
	"html/template"
	"io"

	"github.com/AndreReyesG/blogrenderer/internal/blogposts"
	"github.com/AndreReyesG/blogrenderer/ui"
)

func Render(w io.Writer, p blogposts.Post) error {
	templ, err := template.ParseFS(ui.PostTemplates, "templates/*.html")
	if err != nil {
		return err
	}

	if err := templ.Execute(w, p); err != nil {
		return err
	}

	return nil
}
