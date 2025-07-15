package setConfigGamesteps

import (
	setConfigGameresponses "github.com/autoika/api-config/src/providers/setConfigGame/responses"
	packageglobalsutils "github.com/autoika/package-globals/src/utils"
)

func SetConfigAnimations(
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
	responseBody.Success.Animations = true
	var value = "0"

	if newValue {
		value = "1"
	}

	jsonContent, err := SetConfig(proxy, worldBaseUrl, cookiesStr, userAgentHeader, secChUaHeader, "animations", value, actionRequest)

	if err != nil {
		responseBody.Success.Animations = false

		if cfgDebug {
			if responseBody.Errors == nil {
				responseBody.Errors = &setConfigGameresponses.SetConfigGameSuccessResponseBodyErrors{}
			}

			msg := err.Error()
			responseBody.Errors.Animations = &msg
		}
	}

	return jsonContent
}
