package main

import (
	"html/template"
	"io"

	"github.com/Laellekoenig/htmx-chess/game"
	"github.com/Laellekoenig/htmx-chess/routes"
	"github.com/labstack/echo/v4"
	//"github.com/labstack/echo/v4/middleware"
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
	g := game.CreateGame()

	app := echo.New()
	app.Renderer = newTemplate()
	//app.Use(middleware.Logger())
	app.Static("/static", "static")
	routes.AddRoutes(app, g)
	//app.Logger.Fatal(app.Start(":3000"))
  app.Start(":3000")
}
