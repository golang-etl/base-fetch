package web

import (
	"fmt"

	"github.com/autoika/api-config/src/providers/setConfigGame"
	setConfigGameinterfaces "github.com/autoika/api-config/src/providers/setConfigGame/interfaces"
	packagegeneralinterfaces "github.com/golang-etl/package-general/src/interfaces"
	packagehttputils "github.com/golang-etl/package-http/src/utils"
	"github.com/labstack/echo/v4"
)

func SetConfigGame(setConfigGameProvider setConfigGame.SetConfigGameProvider) func(c echo.Context) error {
	return func(c echo.Context) error {
		var shared *packagegeneralinterfaces.Shared = &packagegeneralinterfaces.Shared{}
		var bindData struct {
			Animations *bool `json:"animations"`
			Tutorial   *bool `json:"tutorial"`
		}

		defer packagehttputils.InternalServerErrorResponse(c, shared, setConfigGameProvider.CfgGoModuleName, setConfigGameProvider.CfgDebug, setConfigGameProvider.CfgDebug)

		if err := c.Bind(&bindData); err != nil {
			panic(fmt.Errorf("error binding request body: %w", err))
		}

		xUserWorldToken := c.Request().Header.Get("x-user-world-token")
		xProxyAuth := c.Request().Header.Get("x-proxy-auth")

		return packagehttputils.AdaptEchoResponse(c, shared, setConfigGameProvider.SetConfigGame(shared, setConfigGameinterfaces.InputData{
			XUserWorldToken: xUserWorldToken,
			XProxyAuth:      xProxyAuth,
			Animations:      bindData.Animations,
			Tutorial:        bindData.Tutorial,
		}))
	}
}
