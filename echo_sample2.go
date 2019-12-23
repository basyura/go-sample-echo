package main

import (
	"fmt"
	"html/template"
	"io"

	"github.com/basyura/go-sample-echo/action"
	"github.com/basyura/go-sample-echo/dto"
	"github.com/labstack/echo"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

// サイトで共通情報
var serviceInfo = dto.ServiceInfo{
	Title: "サイトのタイトル",
}

func main() {

	e := echo.New()

	e.Renderer = &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}

	// url パス, ディレクトリパス
	e.Static("/assets/js", "assets/js")

	e.GET("/page1", func(c echo.Context) error {
		fmt.Println("get /page1")
		for k, v := range c.QueryParams() {
			fmt.Println(k, "=", v)
		}
		return (&action.Page1{}).Execute(c, serviceInfo)
	})

	e.POST("/page1", func(c echo.Context) error {
		fmt.Println("post /page1")
		id1 := c.FormValue("id")
		fmt.Println("id1:", id1)
		return (&action.Page1{}).Execute(c, serviceInfo)
	})

	e.GET("/page1/:user", func(c echo.Context) error {
		fmt.Println("get /page1/:user")
		fmt.Println("path :", c.Param("user"))
		return (&action.Page1{}).Execute(c, serviceInfo)
	})

	e.POST("/page1/:user", func(c echo.Context) error {
		fmt.Println("post /page1/:user")
		fmt.Println("path :", c.Param("user"))
		fmt.Println("  id :", c.FormValue("id"))
		return (&action.Page1{}).Execute(c, serviceInfo)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
