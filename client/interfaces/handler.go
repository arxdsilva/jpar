package interfaces

import (
	"net/http"

	"github.com/arxdsilva/jpar/client/infrastructure/config"
	"github.com/kpango/glg"
	"github.com/labstack/echo/v4"
)

func Run(c config.Config) {
	service := NewHTTP(c.PortService)
	e := echo.New()
	e.GET("/", service.HealthCheck)
	e.GET("/ports", service.PortsList)
	glg.Info("starting server at ", config.Port())
	e.Logger.Fatal(e.Start(config.Port()))
}

func (h *HTTP) HealthCheck(c echo.Context) (err error) {
	ok := struct {
		Service string `json:"service"`
	}{"ok"}
	return c.JSON(http.StatusOK, ok)
}

// PortsList is the http handler for getting all ports
//
// Reposnses:
// 200 OK
// 500 internal server error
func (h *HTTP) PortsList(c echo.Context) (err error) {
	cid := c.Response().Header().Get(echo.HeaderXRequestID)
	glg.Info("[PortsList] cid: ", cid)
	ports, err := h.service.GetPorts()
	if err != nil {
		glg.Error("[PortsList] (GetPorts) error: ", err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "internal server error")
	}
	glg.Info("[PortsList] finish cid: ", cid, "SUCCESS")
	return c.JSON(http.StatusOK, ports)
}
