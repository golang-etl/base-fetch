package main

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/golang-etl/base-fetch/src/config"
	"github.com/golang-etl/base-fetch/src/controllers/web"
	"github.com/golang-etl/base-fetch/src/database"
	"github.com/golang-etl/base-fetch/src/providers/health"
	"github.com/golang-etl/base-fetch/src/providers/login"
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
	loginProvider := login.LoginProvider{CfgGoModuleName: cfg.GoModuleName, CfgDebug: cfg.Debug, Validator: mainValidator, UserTokenModel: userTokenModel, UserTokenProvider: userTokenProvider}

	e.GET("/health", web.GetHealth(healthProvider))
	e.POST("/login", web.Login(loginProvider))

	e.Logger.Fatal(e.Start(cfg.EchoAddress))
}
