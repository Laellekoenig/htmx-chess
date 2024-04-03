package main

import (
	"html/template"
	"io"

	"github.com/Laellekoenig/htmx-chess/data"
	"github.com/Laellekoenig/htmx-chess/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
  tmpl *template.Template
}

func newTemplate() *Template {
  return &Template{
    tmpl: template.Must(template.ParseGlob("views/*.html")),
  }
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
  return t.tmpl.ExecuteTemplate(w, name, data)
}

func main() {
  d := data.CreateData()

  app := echo.New()
  app.Renderer = newTemplate()
  app.Use(middleware.Logger())
  routes.AddRoutes(app, d)
  app.Logger.Fatal(app.Start(":3000"))
}
