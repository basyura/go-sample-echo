package action

import (
	"fmt"
	"net/http"

	"github.com/basyura/go-sample-echo/dto"
	"github.com/basyura/go-sample-echo/html"
	"github.com/labstack/echo"
)

type Page1 struct {
}

func (a *Page1) Execute(c echo.Context, s dto.ServiceInfo) error {

	cols := 300

	table := &html.Table{}
	table.AddHeaderWithClass("cmd", "head_cmd")
	table.AddHeaderWithClass("time", "head_time")
	table.FillHeader(cols)

	i := 0

	for i < 30 {
		row := table.NewRow()
		row.Add(html.Element{Tag: "th", Text: "HogeCommand", Class: "head_cmd head_cmd_value"})
		row.Add(html.Element{Tag: "th", Text: "4321", Class: "head_time head_time_value"})

		if i == 0 {
			k := 0
			for k < 10 {
				row.AddCell("pre")
				row.AddCell("pre")
				row.AddCell("pre")
				row.AddCell("pre")
				row.AddCell("web")
				row.AddCell("web")
				row.AddCell("web")
				row.AddCell("web")
				row.AddCell("web")
				row.AddCell("web")
				row.AddCell("web")
				row.AddCell("web")
				row.AddCell("comp")
				row.AddCell("comp")
				row.AddCell("comp")
				row.AddCell("comp")
				row.AddCell("diff")
				row.AddCell("diff")
				row.AddCell("diff")
				row.AddCell("diff")
				k++
			}
			//row.FillCell(cols)
		} else {
			//row.FillCell(cols)
		}
		i++
	}

	html := table.ToHtml()
	fmt.Println(html)

	dto := dto.Page1Res{
		ServiceInfo: s,
		Content:     "ページのコンテンツ",
		Table:       html,
	}

	return c.Render(http.StatusOK, "page1", dto)
}
