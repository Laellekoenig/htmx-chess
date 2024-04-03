package routes

import (
	"net/http"
	"strconv"

	"github.com/Laellekoenig/htmx-chess/data"
	"github.com/labstack/echo/v4"
)

func AddRoutes(app *echo.Echo, d *data.Data) {
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
    return c.Render(http.StatusOK, "board", *d)
  })

  app.DELETE("/remove-active", func(c echo.Context) error {
    data.ClearAllActiveSquares(d)
    return c.Render(http.StatusOK, "board", *d)
  })
}
