package main

import (
	"fmt"

	"github.com/autoika/api-config/src/config"
	"github.com/autoika/api-config/src/controllers/web"
	"github.com/autoika/api-config/src/database"
	"github.com/autoika/api-config/src/providers/health"
	"github.com/autoika/api-config/src/providers/setConfigGame"
	"github.com/go-playground/validator/v10"
	packagegeneralutils "github.com/golang-etl/package-general/src/utils"
	packageusertokenmodels "github.com/golang-etl/package-user-token/src/models"
	"github.com/golang-etl/package-user-token/src/providers/usertoken"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		panic(fmt.Errorf("error loading config: %w", err))
	}

	mainDB := database.MainDB{}
	mainDB.Connect(cfg.MongoDBURI)
	mainDB.Ping(cfg.MongoDBDatabaseName)
	defer mainDB.Disconnect()

	e := echo.New()
	e.Use(middleware.Recover())

	mainValidator := packagegeneralutils.PrepareValidator(validator.New())

	userTokenModel := packageusertokenmodels.UserTokenModel{Client: mainDB.Client, Secret: cfg.SecretKeyUserTokenData, Database: cfg.MongoDBDatabaseName}

	userTokenProvider := usertoken.UserTokenProvider{UserTokenModel: userTokenModel}
	healthProvider := health.HealthProvider{CfgGoModuleName: cfg.GoModuleName, CfgDebug: cfg.Debug, MongoClient: mainDB.Client}
	setConfigGameProvider := setConfigGame.SetConfigGameProvider{CfgGoModuleName: cfg.GoModuleName, CfgDebug: cfg.Debug, Validator: mainValidator, UserTokenModel: userTokenModel, UserTokenProvider: userTokenProvider}

	e.GET("/health", web.GetHealth(healthProvider))
	e.POST("/config/game", web.SetConfigGame(setConfigGameProvider))

	e.Logger.Fatal(e.Start(cfg.EchoAddress))
}
