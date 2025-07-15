package setConfigGame

import (
	"fmt"

	setConfigGameinterfaces "github.com/autoika/api-config/src/providers/setConfigGame/interfaces"
	setConfigGamepreparers "github.com/autoika/api-config/src/providers/setConfigGame/preparers"
	setConfigGameresponses "github.com/autoika/api-config/src/providers/setConfigGame/responses"
	setConfigGamesteps "github.com/autoika/api-config/src/providers/setConfigGame/steps"
	packageglobalsutils "github.com/autoika/package-globals/src/utils"
	packageikariamservices "github.com/autoika/package-ikariam/src/services"
	packageikariamutils "github.com/autoika/package-ikariam/src/utils"
	"github.com/go-playground/validator/v10"
	packagegeneralinterfaces "github.com/golang-etl/package-general/src/interfaces"
	packagehttpinterfaces "github.com/golang-etl/package-http/src/interfaces"
	packagehttputils "github.com/golang-etl/package-http/src/utils"
	packageplaywrightutils "github.com/golang-etl/package-playwright/src/utils"
	packageusertokenmodels "github.com/golang-etl/package-user-token/src/models"
	"github.com/golang-etl/package-user-token/src/providers/usertoken"
	packageusertokenutils "github.com/golang-etl/package-user-token/src/utils"
)

type SetConfigGameProvider struct {
	CfgGoModuleName   string
	CfgDebug          bool
	CfgUserAgent      string
	CfgSecChUaHeader  string
	Validator         *validator.Validate
	UserTokenModel    packageusertokenmodels.UserTokenModel
	UserTokenProvider usertoken.UserTokenProvider
}

func (provider SetConfigGameProvider) SetConfigGame(shared *packagegeneralinterfaces.Shared, originalInputData setConfigGameinterfaces.InputData) packagehttpinterfaces.Response {
	inputData := setConfigGamepreparers.DefaultInputData(originalInputData)

	err := provider.Validator.Struct(inputData)

	if err != nil {
		return packagehttputils.ValidationErrorHandlerToUnprocessableEntityResponse(err, inputData, map[string]string{
			"xUserWorldToken.required": "El token de usuario de mundo es obligatorio.",
		})
	}

	proxy := packageglobalsutils.ParseProxyString(inputData.XProxyAuth)
	userWorldToken, response := provider.UserTokenProvider.GetInstance(inputData.XUserWorldToken)

	if response != nil {
		return *response
	}

	storage := packageplaywrightutils.GenerateOptionalStorageState(userWorldToken.Context)
	cookies := packageplaywrightutils.OptionalCookiesToHeader(storage.Cookies)

	actionRequest, err, response := packageikariamservices.GetRenewActionRequestService(proxy, userWorldToken.Extra["worldUrl"], cookies, provider.CfgUserAgent, provider.CfgSecChUaHeader)

	if err != nil {
		panic(fmt.Errorf("error al obtener el 'actionRequest': %w", err))
	}

	if response != nil {
		return *response
	}

	responseBody := setConfigGameresponses.SetConfigGameSuccessResponseBody{}

	if inputData.Animations != nil {
		jsonContent := setConfigGamesteps.SetConfigAnimations(proxy, userWorldToken.Extra["worldUrl"], cookies, provider.CfgUserAgent, provider.CfgSecChUaHeader, *inputData.Animations, &responseBody, actionRequest, provider.CfgDebug)

		if packageikariamutils.IsExpiredSessionFromAjax(jsonContent) {
			return packageusertokenutils.ExpiredUserTokenResponse()
		}

		actionRequest = packageikariamutils.GetActionRequestFromAjax(jsonContent, actionRequest)
	}

	if inputData.Tutorial != nil {
		jsonContent := setConfigGamesteps.SetConfigTutorial(proxy, userWorldToken.Extra["worldUrl"], cookies, provider.CfgUserAgent, provider.CfgSecChUaHeader, *inputData.Tutorial, &responseBody, actionRequest, provider.CfgDebug)

		if packageikariamutils.IsExpiredSessionFromAjax(jsonContent) {
			return packageusertokenutils.ExpiredUserTokenResponse()
		}

		actionRequest = packageikariamutils.GetActionRequestFromAjax(jsonContent, actionRequest)
	}

	return setConfigGameresponses.SetConfigGameSuccessResponse(responseBody)
}
