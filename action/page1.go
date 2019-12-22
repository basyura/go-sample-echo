package action

import (
	"net/http"

	"github.com/basyura/go-sample-echo/dto"
	"github.com/labstack/echo"
)

type Page1 struct {
}

func (a *Page1) Execute(c echo.Context, s dto.ServiceInfo) error {
	dto := dto.Page1Res{
		ServiceInfo: s,
		Content:     "ページのコンテンツ",
	}

	return c.Render(http.StatusOK, "page1", dto)
}
