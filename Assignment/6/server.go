package main

import (
	"database/sql"
	"html/template"
	"io"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	// Static File
	e.Static("/static", "public/assets")

	// html template
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("public/view/*.html")),
	}
	e.Renderer = renderer

	e.GET("/", index)

	e.GET("/member", member)

	e.GET("/error", errorPage)

	e.POST("/signup", signup)

	e.POST("/signin", signin)

	e.Logger.Fatal(e.Start(":1323"))
}

func index(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{})
}

func member(c echo.Context) error {
	user := c.QueryParam("name")

	return c.Render(http.StatusOK, "member.html", map[string]interface{}{
		"user": user,
	})
}

func errorPage(c echo.Context) error {
	msg := c.QueryParam("errorMsg")

	return c.Render(http.StatusOK, "error.html", map[string]interface{}{
		"errorMsg": msg,
	})
}

// === API ===
type Signup struct {
	Name     string `json:"name"`
	Account  string `json:"account"`
	PassWord string `json:"passWord"`
}

type Signin struct {
	Account  string `json:"account"`
	PassWord string `json:"passWord"`
}

type DatabaseResult struct {
	Id       int
	Name     string
	Username string
	Password string
	Time     time.Time
}

type Response struct {
	Status       string
	ErrorMessage string
	Name         string
}

// Sign Up
func signup(c echo.Context) error {
	req := new(Signup)
	if err := c.Bind(req); err != nil {
		return err
	}

	rep := new(Response)
	rep.Name = req.Name

	// Connect DB
	db, err := sql.Open("mysql", "root:a1b2c3448@/website")
	checkErr(err)

	// Check Account
	rows, err := db.Query("SELECT COUNT(username) FROM website.user WHERE username = ?", req.Name)
	checkErr(err)

	for rows.Next() {
		var result int
		err = rows.Scan(&result)
		checkErr(err)

		if result == 0 {
			rep.Status = "success"

			_, err := db.Exec("INSERT INTO user (name, username, password) VALUES (?, ?, ?)", req.Name, req.Account, req.PassWord)
			checkErr(err)
			println("Create Account")
		} else {
			rep.Status = "error"
			rep.ErrorMessage = "帳號已經被註冊"
			println("User Alreadty exist")
		}
	}

	db.Close()

	return c.JSON(http.StatusOK, rep)
}

// Sign Up
func signin(c echo.Context) error {
	req := new(Signin)
	if err := c.Bind(req); err != nil {
		return err
	}

	rep := new(Response)

	// Connect DB
	db, err := sql.Open("mysql", "root:a1b2c3448@/website")
	checkErr(err)

	// Check Account
	rows, err := db.Query("SELECT COUNT(username) FROM website.user WHERE username = ?", req.Account)
	checkErr(err)

	for rows.Next() {
		var result int
		err = rows.Scan(&result)
		checkErr(err)

		if result == 0 {
			rep.Status = "error"
			rep.ErrorMessage = "帳號或密碼輸入錯誤"
			println("Account or Password Error")
		} else {
			var name string
			row := db.QueryRow("SELECT name FROM website.user WHERE username = ?", req.Account)
			row.Scan(&name)

			rep.Status = "success"
			rep.Name = name
		}
	}

	db.Close()

	return c.JSON(http.StatusOK, rep)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
