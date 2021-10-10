package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", helloEcho)

	e.GET("/query", queryResult)

	e.GET("/word/:text", wordResult)

	e.Logger.Fatal(e.Start(":8000"))
}

func helloEcho(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, Echo")
}

func queryResult(c echo.Context) error {
	word := c.QueryParam("word")

	if word != "" {
		return c.String(http.StatusOK, fmt.Sprintf("Query Result : %s", word))
	} else {
		return c.String(http.StatusOK, "Query Result")
	}
}

func wordResult(c echo.Context) error {
	text := c.Param("text")

	return c.String(http.StatusOK, fmt.Sprint(text))
}
