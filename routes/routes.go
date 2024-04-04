package routes

import (
	"net/http"
	"strconv"

	"github.com/Laellekoenig/htmx-chess/game"
	"github.com/labstack/echo/v4"
)

func AddRoutes(app *echo.Echo, g *game.Game) {
	app.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", *g)
	})

	app.GET("/refresh-board", func(c echo.Context) error {
		return c.Render(http.StatusOK, "board", *g)
	})

	app.POST("/select-square/:id", func(c echo.Context) error {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.String(400, "Invalid id")
		}

		g.ActivateSquare(id)
    g.MoveActiveSquares()

		return c.Render(http.StatusOK, "page", *g)
	})

	app.POST("/set-fen", func(c echo.Context) error {
		fen := c.FormValue("fen")
		g.ClearBoard()
		g.FillFen(fen)
		return c.Render(http.StatusOK, "page", *g)
	})

	app.DELETE("/reset-board", func(c echo.Context) error {
		g.SetStartingPos()
		return c.Render(http.StatusOK, "page", *g)
	})

	app.DELETE("/remove-active", func(c echo.Context) error {
		g.ClearActiveSquares()
		return c.Render(http.StatusOK, "board", *g)
	})
}
