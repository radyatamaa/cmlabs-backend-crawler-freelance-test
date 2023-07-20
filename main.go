package main

import (
	//"context"
	//"fmt"
	//"log"
	//"net/http"
	//"strings"
	//"time"
	//
	//"github.com/hibiken/asynq"
	//"github.com/radyatama/cmlabs-backend-crawler-freelance-test/internal/domain"
	//"github.com/radyatama/cmlabs-backend-crawler-freelance-test/pkg/helper"
	//"github.com/radyatama/cmlabs-backend-crawler-freelance-test/pkg/httpclient"
	//"github.com/radyatama/cmlabs-backend-crawler-freelance-test/pkg/httprate"
	//"github.com/radyatama/cmlabs-backend-crawler-freelance-test/pkg/jwt"
	//"github.com/radyatama/cmlabs-backend-crawler-freelance-test/pkg/validator"
	//
	//"github.com/beego/beego/v2/client/cache"
	//_ "github.com/beego/beego/v2/client/cache/redis"
	//
	//"github.com/radyatama/cmlabs-backend-crawler-freelance-test/internal"
	//"github.com/radyatama/cmlabs-backend-crawler-freelance-test/pkg/database"
	//
	//beego "github.com/beego/beego/v2/server/web"
	//beegoContext "github.com/beego/beego/v2/server/web/context"
	//"github.com/beego/beego/v2/server/web/filter/cors"
	//"github.com/beego/i18n"
	//"github.com/radyatama/cmlabs-backend-crawler-freelance-test/internal/middlewares"
	//"github.com/radyatama/cmlabs-backend-crawler-freelance-test/pkg/response"
	//"github.com/radyatama/cmlabs-backend-crawler-freelance-test/pkg/zaplogger"

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
	//err := beego.LoadAppConfig("ini", "conf/app.ini")
	//if err != nil {
	//	panic(err)
	//}
	//// token expired
	//tokenExpired := beego.AppConfig.DefaultInt64("tokenExpired", 86400)
	//// global execution timeout
	//serverTimeout := beego.AppConfig.DefaultInt64("serverTimeout", 60)
	//// global execution timeout
	//requestTimeout := beego.AppConfig.DefaultInt("executionTimeout", 5)
	//// global execution timeout to second
	//timeoutContext := time.Duration(requestTimeout) * time.Second
	//// web hook to slack error log
	//slackWebHookUrl := beego.AppConfig.DefaultString("slackWebhookUrlLog", "")
	//// app version
	//appVersion := beego.AppConfig.DefaultString("version", "1")
	//// log path
	//logPath := beego.AppConfig.DefaultString("logPath", "./logs/api.log")
	//// redis connection config
	//redisConnectionConfig := beego.AppConfig.DefaultString("redisBeegoConConfig", `{"conn":"127.0.0.1:6379"}`)
	//// jwt secret key
	//jwtSecretKey := beego.AppConfig.DefaultString("jwtSecretKey", "secret")
	//// init data
	//initData := beego.AppConfig.DefaultString("initData", "true")
	//
	//// config helper
	//appConfigHelper := helper.ConfigHelper{
	//	AppUrl: beego.AppConfig.DefaultString("appUrl", "http://localhost:8082"),
	//}
	//// database initialization
	//db := database.DB()
	//
	//// language
	//lang := beego.AppConfig.DefaultString("lang", "en|id")
	//languages := strings.Split(lang, "|")
	//for _, value := range languages {
	//	if err := i18n.SetMessage(value, "./conf/"+value+".ini"); err != nil {
	//		panic("Failed to set message file for l10n")
	//	}
	//}
	//
	//// beego config
	//beego.BConfig.Log.AccessLogs = false
	//beego.BConfig.Log.EnableStaticLogs = false
	//beego.BConfig.Listen.ServerTimeOut = serverTimeout
	//
	//// zap logger
	//zapLog := zaplogger.NewZapLogger(logPath, slackWebHookUrl)
	//
	//if beego.BConfig.RunMode == "dev" {
	//	// db auto migrate dev environment
	//	if err := db.AutoMigrate(
	//		&domain.Admin{},
	//		&domain.Customer{},
	//		&domain.Order{},
	//		&domain.Product{},
	//		&domain.OrderDetail{},
	//	); err != nil {
	//		panic(err)
	//	}
	//
	//	// static files swagger
	//	beego.BConfig.WebConfig.DirectoryIndex = true
	//	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	//}
	//
	//// init redis
	//redisCache, err := cache.NewCache("redis", redisConnectionConfig)
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//// config validator
	//validator.Validate.SetDatabaseConnection(db)
	//
	//// jwt middleware
	//auth, err := jwt.NewJwt(&jwt.Options{
	//	SignMethod:  jwt.HS256,
	//	SecretKey:   jwtSecretKey,
	//	Locations:   "header:Authorization",
	//	IdentityKey: "uid",
	//})
	//if err != nil {
	//	panic(err)
	//}
	//
	//// set adapter redis for jwt middleware
	//auth.SetAdapter(redisCache)
	//if initData == "true" {
	//	domain.SeederDataAdmin(db)
	//	domain.SeederDataCustomer(db)
	//	domain.SeederDataProduct(db)
	//}
	//if beego.BConfig.RunMode != "prod" {
	//	// static files swagger
	//	beego.BConfig.WebConfig.DirectoryIndex = true
	//	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	//}
	//
	//beego.BConfig.WebConfig.StaticDir["/external"] = "external"
	//
	//// middleware init
	//beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
	//	AllowMethods:    []string{http.MethodGet, http.MethodPost},
	//	AllowAllOrigins: true,
	//}))
	//
	//maxRequestAPIInMinutes := time.Duration(beego.AppConfig.DefaultInt("maxRequestAPIInMinutes", 1)) * time.Minute
	//
	//beego.InsertFilterChain("/api/*", httprate.LimitByIP(beego.AppConfig.DefaultInt("maxRequestAPI", 100),
	//	maxRequestAPIInMinutes))
	//beego.InsertFilterChain("*", middlewares.RequestID())
	//beego.InsertFilterChain("/api/*", middlewares.BodyDumpWithConfig(middlewares.NewAccessLogMiddleware(zapLog, appVersion).Logger()))
	//beego.InsertFilterChain("/api/v1/*", middlewares.NewJwtMiddleware().JwtMiddleware(auth))
	//// health check
	//beego.Get("/health", func(ctx *beegoContext.Context) {
	//	ctx.Output.SetStatus(http.StatusOK)
	//	ctx.Output.JSON(beego.M{"status": "alive"}, beego.BConfig.RunMode != "prod", false)
	//})
	//
	//// resty http client
	//var configDebug bool
	//if beego.BConfig.RunMode == "local" {
	//	configDebug = true
	//}
	//
	//restyClient := httpclient.NewRestyHttpClient(
	//	httpclient.ConfigDebug(configDebug),
	//	httpclient.ConfigTimeout(30*time.Second),
	//	httpclient.ConfigRetryCount(3),
	//	httpclient.ConfigRetryWaitTime(100*time.Millisecond),
	//	httpclient.ConfigRetryMaxWaitTime(2*time.Second),
	//	httpclient.ConfigLogger(zapLog),
	//)
	//
	//// default error handler
	//beego.ErrorController(&response.ErrorController{})
	//
	//// init repository
	////adminRepo := adminRepository.NewMysqlAdminRepository(db, zapLog)
	////customerRepo := customerRepository.NewMysqlCustomerRepository(db, zapLog)
	////orderRepo := orderRepository.NewMysqlOrderRepository(db, zapLog)
	////productRepo := productRepository.NewMysqlProductRepository(db, zapLog)
	////orderDetailRepo := orderDetailRepository.NewMysqlOrderDetailRepository(db, zapLog)
	////orderRestRepo := orderRestRepository.NewRestOrderRepository(restyClient, zapLog)

	// init usecase
	//authUseCase := authUsecase.NewAuthUseCase(timeoutContext, customerRepo, adminRepo, auth, int(tokenExpired), zapLog)
	//orderUseCase := orderUsecase.NewOrderUseCase(timeoutContext, orderRepo, orderDetailRepo, productRepo, appConfigHelper, authUseCase, orderRestRepo, zapLog)
	//productUseCase := productUsecase.NewProductUseCase(timeoutContext, productRepo, zapLog)

	// init handler
	//authtHandler.NewAuthHandler(authUseCase, zapLog)
	//orderHandler.NewOrderHandler(orderUseCase, zapLog)
	//productHandler.NewProductHandler(productUseCase, zapLog)

	// default error handler
	//beego.ErrorController(&internal.BaseController{})
	//
	//beego.Run()
}
