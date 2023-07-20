package main

import (
	beego "github.com/beego/beego/v2/server/web"
	beegoContext "github.com/beego/beego/v2/server/web/context"
	"github.com/beego/beego/v2/server/web/filter/cors"
	"github.com/beego/i18n"
	"github.com/gocolly/colly"
	"github.com/radyatama/cmlabs-backend-crawler-freelance-test/internal"
	crawlHandler "github.com/radyatama/cmlabs-backend-crawler-freelance-test/internal/crawl/delivery/http/v1"
	crawlUsecase "github.com/radyatama/cmlabs-backend-crawler-freelance-test/internal/crawl/usecase"
	"github.com/radyatama/cmlabs-backend-crawler-freelance-test/internal/middlewares"
	"github.com/radyatama/cmlabs-backend-crawler-freelance-test/pkg/httprate"
	"github.com/radyatama/cmlabs-backend-crawler-freelance-test/pkg/response"
	"github.com/radyatama/cmlabs-backend-crawler-freelance-test/pkg/validator"
	"github.com/radyatama/cmlabs-backend-crawler-freelance-test/pkg/zaplogger"
	"net/http"
	"strings"
	"time"
)

// @title Api Gateway V1
// @version v1
// @contact.name radyatama
// @contact.email mohradyatama24@gmail.com
// @description api "API Gateway v1"
// @BasePath /api
// @query.collection.format multi
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	err := beego.LoadAppConfig("ini", "conf/app.ini")
	if err != nil {
		panic(err)
	}
	// global execution timeout
	serverTimeout := beego.AppConfig.DefaultInt64("serverTimeout", 60)
	// global execution timeout
	requestTimeout := beego.AppConfig.DefaultInt("executionTimeout", 5)
	// global execution timeout to second
	timeoutContext := time.Duration(requestTimeout) * time.Second
	// web hook to slack error log
	slackWebHookUrl := beego.AppConfig.DefaultString("slackWebhookUrlLog", "")
	// app version
	appVersion := beego.AppConfig.DefaultString("version", "1")
	// log path
	logPath := beego.AppConfig.DefaultString("logPath", "./logs/api.log")

	// language
	lang := beego.AppConfig.DefaultString("lang", "en|id")
	languages := strings.Split(lang, "|")
	for _, value := range languages {
		if err := i18n.SetMessage(value, "./conf/"+value+".ini"); err != nil {
			panic("Failed to set message file for l10n")
		}
	}

	// beego config
	beego.BConfig.Log.AccessLogs = false
	beego.BConfig.Log.EnableStaticLogs = false
	beego.BConfig.Listen.ServerTimeOut = serverTimeout

	// zap logger
	zapLog := zaplogger.NewZapLogger(logPath, slackWebHookUrl)

	if beego.BConfig.RunMode == "dev" {
		// static files swagger
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// config validator
	validator.Validate.SetDatabaseConnection(nil)

	if beego.BConfig.RunMode != "prod" {
		// static files swagger
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.BConfig.WebConfig.StaticDir["/external"] = "external"

	// middleware init
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowMethods:    []string{http.MethodGet, http.MethodPost},
		AllowAllOrigins: true,
	}))

	maxRequestAPIInMinutes := time.Duration(beego.AppConfig.DefaultInt("maxRequestAPIInMinutes", 1)) * time.Minute

	beego.InsertFilterChain("/api/*", httprate.LimitByIP(beego.AppConfig.DefaultInt("maxRequestAPI", 100),
		maxRequestAPIInMinutes))
	beego.InsertFilterChain("*", middlewares.RequestID())
	beego.InsertFilterChain("/api/*", middlewares.BodyDumpWithConfig(middlewares.NewAccessLogMiddleware(zapLog, appVersion).Logger()))
	// health check
	beego.Get("/health", func(ctx *beegoContext.Context) {
		ctx.Output.SetStatus(http.StatusOK)
		ctx.Output.JSON(beego.M{"status": "alive"}, beego.BConfig.RunMode != "prod", false)
	})

	// colly Init
	coll := colly.NewCollector()

	coll.UserAgent = "Go program"

	// default error handler
	beego.ErrorController(&response.ErrorController{})

	// init repository

	//init usecase
	crawlUseCase := crawlUsecase.NewCrawlUseCase(coll, timeoutContext, zapLog)

	//init handler
	crawlHandler.NewCrawlHandler(crawlUseCase, zapLog)

	//default error handler
	beego.ErrorController(&internal.BaseController{})

	beego.Run()
}
