package main

import (
	"html/template"
	"io"
	"net/http"
	"strconv"

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

type Square struct {
  Num int
  IsBlack bool
  IsActive bool
}

type Data struct {
  Squares []Square
  ActiveSquares []*Square
}

func createData() *Data {
  var squares []Square

  for i := 0; i < 8; i++ {
    for j := 0; j < 8; j++ {
      isBlack := true

      if i % 2 == 0 && j % 2 == 0  {
        isBlack = false
      } else if i % 2 != 0 && j % 2 != 0 {
        isBlack = false
      }

      squares = append(squares, Square{Num: i * 8 + j, IsBlack: isBlack, IsActive: false})
    }
  }

  return &Data{Squares: squares, ActiveSquares: []*Square{}}
}

func setActive(data *Data, num int) {
  square := &(*data).Squares[num]
  (*square).IsActive = true
  data.ActiveSquares = append(data.ActiveSquares, square)
}

func clearAllActiveSquares(data *Data) {
  for _, square := range (*data).ActiveSquares {
    (*square).IsActive = false
  }
  data.ActiveSquares = []*Square{}
}

func main() {
  data := createData()

  app := echo.New()
  app.Renderer = newTemplate()
  app.Use(middleware.Logger())

  app.GET("/", func(c echo.Context) error {
    return c.Render(http.StatusOK, "index.html", *data)
  })

  app.POST("/select-square/:id", func(c echo.Context) error {
    idStr := c.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
      return c.String(400, "Invalid id")
    }
    setActive(data, id)
    return c.Render(http.StatusOK, "board", *data)
  })

  app.DELETE("/remove-active", func(c echo.Context) error {
    clearAllActiveSquares(data)
    return c.Render(http.StatusOK, "board", *data)
  })

  app.Logger.Fatal(app.Start(":3000"))
}
