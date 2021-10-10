package main

import (
	"io"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	e.Static("/", "assets")

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("*.html")),
	}
	e.Renderer = renderer

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<h1>Go Go Go</h1><a href='/word/Hello Gopher'>Link</a>")
	})

	e.GET("/word/:text", func(c echo.Context) error {
		text := c.Param("text")

		return c.Render(
			http.StatusOK,
			"template.html",
			map[string]interface{}{
				"word": text,
			},
		)
	})

	e.Logger.Fatal(e.Start(":8000"))
}
