package teas

import (
	"github.com/enuesaa/teatime/pkg/router/ctx"
	"github.com/enuesaa/teatime/pkg/service"
	"github.com/labstack/echo/v4"
)

type Tea struct {
	Teaid  string   `json:"teaid"`
	Vals   []TeaVal `json:"vals"`
}
type TeaVal struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
	Type  string      `json:"type"`
}

func List(c echo.Context) error {
	cc := ctx.Use(c)
	teapod := cc.Param("teapod")

	teaSrv := service.NewTeaSrv(cc.Repos, teapod)
	list, err := teaSrv.List()
	if err != nil {
		return err
	}
	return cc.WithData(list)
}
