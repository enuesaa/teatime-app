package teas

import (
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

func (ctl *Ctl) ListTeas(c echo.Context) error {
	teapod := c.Param("teapod")

	teaSrv := service.NewTeaSrv(ctl.repos, teapod)
	list, err := teaSrv.List()
	if err != nil {
		return err
	}
	return ctl.WithData(c, list)
}