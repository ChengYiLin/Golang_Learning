package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"

	"github.com/labstack/echo-contrib/session"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<h1>Go Go Session</h1><a href='/hello?name=Martin'>Link</a>")
	})

	e.GET("/hello", func(c echo.Context) error {
		name := c.QueryParam("name")

		sess, _ := session.Get("session", c)
		sess.Options = &sessions.Options{
			Path:     "/talk",
			MaxAge:   86400 * 7,
			HttpOnly: true,
		}
		sess.Values["name"] = name
		sess.Save(c.Request(), c.Response())

		return c.HTML(http.StatusOK, "<h1>Save Session</h1><a href='/talk'>Go Talk</a>")
	})

	e.GET("/talk", func(c echo.Context) error {
		sess, _ := session.Get("session", c)

		name, ok := sess.Values["name"].(string)

		if ok {
			return c.String(http.StatusOK, fmt.Sprintf("歡迎回來，%s", name))
		} else {
			return c.String(http.StatusOK, "不好意思，不認識你哦")
		}

	})

	e.Logger.Fatal(e.Start(":8000"))
}
