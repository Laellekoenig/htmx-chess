package main

import (
	"html/template"
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
  "github.com/Laellekoenig/htmx-chess/data"
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

  app.GET("/", func(c echo.Context) error {
    return c.Render(http.StatusOK, "index.html", *d)
  })

  app.GET("/refresh-board", func(c echo.Context) error {
    return c.Render(http.StatusOK, "board", *d)
  })

  app.POST("/select-square/:id", func(c echo.Context) error {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
      return c.String(400, "Invalid id")
    }
    data.SetActive(d, id)
    return c.Render(http.StatusOK, "square", (*d).Squares[id])
  })

  app.DELETE("/remove-active", func(c echo.Context) error {
    data.ClearAllActiveSquares(d)
    return c.Render(http.StatusOK, "board", *d)
  })

  app.Logger.Fatal(app.Start(":3000"))
}
