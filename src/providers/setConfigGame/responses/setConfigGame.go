package setConfigGameresponses

import (
	"net/http"

	"github.com/golang-etl/package-http/src/consts"
	packagehttpinterfaces "github.com/golang-etl/package-http/src/interfaces"
)

type SetConfigGameSuccessResponseBodySuccess struct {
	Animations bool `json:"animations"`
	Tutorial   bool `json:"tutorial"`
}

type SetConfigGameSuccessResponseBodyErrors struct {
	Animations *string `json:"animations,omitempty"`
	Tutorial   *string `json:"tutorial,omitempty"`
}

type SetConfigGameSuccessResponseBody struct {
	Success SetConfigGameSuccessResponseBodySuccess `json:"success"`
	Errors  *SetConfigGameSuccessResponseBodyErrors `json:"errors,omitempty"`
}

func SetConfigGameSuccessResponse(body SetConfigGameSuccessResponseBody) packagehttpinterfaces.Response {
	return packagehttpinterfaces.Response{
		StatusCode: http.StatusOK,
		Headers:    consts.HeaderContentType.JSON,
		Body:       body,
	}
}
