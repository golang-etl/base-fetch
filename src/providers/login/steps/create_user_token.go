package loginsteps

import (
	"fmt"

	packagegeneralutils "github.com/golang-etl/package-general/src/utils"
	packageusertokenmodels "github.com/golang-etl/package-user-token/src/models"
)

func CreateUserToken(userTokenModel packageusertokenmodels.UserTokenModel, extra map[string]string) packageusertokenmodels.UserToken {
	token := packagegeneralutils.GenerateRandToken(128)

	userToken := packageusertokenmodels.UserToken{
		Token: token,
		Extra: extra,
	}

	err := userTokenModel.Insert(userToken)

	if err != nil {
		panic(fmt.Errorf("error al insertar el token de usuario: %w", err))
	}

	return userToken
}
