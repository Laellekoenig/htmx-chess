package main

import (
	//"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
  index := views.Index

  app := echo.New()
  app.Use(middleware.Logger())
  //app.GET("/", func(c echo.Context) error {
  //  return c.Render(200, templ.Handler(index), count)
  //})
  app.Logger.Fatal(app.Start(":3000"))
}
