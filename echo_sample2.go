package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

type ServiceInfo struct {
	Title string
}

type Page1ReqDto struct {
	ServiceInfo
	Content string
}

// サイトで共通情報
var serviceInfo = ServiceInfo{
	"サイトのタイトル",
}

func main() {

	e := echo.New()

	e.Renderer = &Template{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
	e.Static("/assets/js", "assets/js")

	e.GET("/page1", func(c echo.Context) error {
		dto := Page1ReqDto{
			ServiceInfo: serviceInfo,
			Content:     "ページのコンテンツ",
		}
		return c.Render(http.StatusOK, "page1", dto)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
