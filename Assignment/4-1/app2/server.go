package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", index)

	e.GET("/error", error404)

	e.GET("/query", queryRedirect)

	e.Logger.Fatal(e.Start(":8000"))
}

func index(c echo.Context) error {
	return c.HTML(
		http.StatusOK,
		"<h3>歡迎光臨</h3><div>這是我們的首頁</div>",
	)
}

func error404(c echo.Context) error {
	return c.HTML(
		http.StatusBadRequest,
		"<h3>發生錯誤</h3><div>請返回首頁</div>",
	)
}

func queryRedirect(c echo.Context) error {
	word := c.QueryParam("word")

	if word != "" {
		return c.JSON(
			http.StatusOK,
			map[string]string{
				"word":      word,
				"translate": word,
				"status":    "ok",
			},
		)
	} else {
		return c.Redirect(http.StatusSeeOther, "/error")
	}
}
