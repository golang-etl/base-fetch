package setConfigGamesteps

import (
	setConfigGameresponses "github.com/autoika/api-config/src/providers/setConfigGame/responses"
	packageglobalsutils "github.com/autoika/package-globals/src/utils"
)

func SetConfigTutorial(
	proxy *packageglobalsutils.ParsedProxy,
	worldBaseUrl string,
	cookiesStr string,
	userAgentHeader string,
	secChUaHeader string,
	newValue bool,
	responseBody *setConfigGameresponses.SetConfigGameSuccessResponseBody,
	actionRequest string,
	cfgDebug bool,
) string {
	responseBody.Success.Tutorial = true
	var value = "off"

	if newValue {
		value = "on"
	}

	jsonContent, err := SetConfig(proxy, worldBaseUrl, cookiesStr, userAgentHeader, secChUaHeader, "tutorialOptions", value, actionRequest)

	if err != nil {
		responseBody.Success.Tutorial = false

		if cfgDebug {
			if responseBody.Errors == nil {
				responseBody.Errors = &setConfigGameresponses.SetConfigGameSuccessResponseBodyErrors{}
			}

			msg := err.Error()
			responseBody.Errors.Tutorial = &msg
		}
	}

	return jsonContent
}
