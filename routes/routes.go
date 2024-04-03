package routes

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

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
		game.SetActive(g, id)
		return c.Render(http.StatusOK, "board", *g)
	})

	app.POST("/set-fen", func(c echo.Context) error {
		fen := c.FormValue("fen")
		game.ClearBoard(g)

    if strings.Contains(fen, " ") {
      game.ParseFEN(fen, g)
    } else {
		  game.ParseFENPosition(fen, g)
    }

		return c.Render(http.StatusOK, "page", *g)
	})

	app.DELETE("/reset-board", func(c echo.Context) error {
    err := game.ParseFEN(game.FEN_START, g)
    fmt.Println(err)
		return c.Render(http.StatusOK, "board", *g)
	})

	app.DELETE("/clear-board", func(c echo.Context) error {
		game.ClearBoard(g)
		return c.Render(http.StatusOK, "board", *g)
	})

	app.DELETE("/remove-active", func(c echo.Context) error {
		game.ClearAllActiveSquares(g)
		return c.Render(http.StatusOK, "board", *g)
	})
}
