package main

import (
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
		return (&action.Page1{}).Execute(c, serviceInfo)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
