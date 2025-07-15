package setConfigGamesteps

import (
	"fmt"

	setConfigGameservices "github.com/autoika/api-config/src/providers/setConfigGame/services"
	packageglobalsutils "github.com/autoika/package-globals/src/utils"
)

func SetConfig(
	proxy *packageglobalsutils.ParsedProxy,
	worldBaseUrl string,
	cookiesStr string,
	userAgentHeader string,
	secChUaHeader string,
	toChange string,
	newValue string,
	actionRequest string,
) (string, error) {
	queryParams := setConfigGameservices.SetGameConfigOptionQueryParams{
		Action:         "Options",
		Function:       "changeAvatarOptions",
		Category:       toChange,
		Value:          newValue,
		BackgroundView: "worldmap_iso",
		TemplateView:   "options",
		ActionRequest:  actionRequest,
		Ajax:           "1",
	}

	jsonContent, err := setConfigGameservices.SetGameConfigOption(proxy, worldBaseUrl, cookiesStr, userAgentHeader, secChUaHeader, queryParams)

	if err != nil {
		return "", fmt.Errorf("no se pudo establecer la configuración del juego: %w", err)
	}

	return jsonContent, nil
}
