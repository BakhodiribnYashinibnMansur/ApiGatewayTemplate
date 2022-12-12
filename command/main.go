package main

import (
	"api_getaway_web/config"
	"api_getaway_web/package/handler"
	"api_getaway_web/util/logrus_log"
	"fmt"
)

// @title Rizo Mulk
// @version 1.0
// @description API Server for Rizo Mulk Application
// @termsOfService gitlab.com/rizoMulk
// @host gitlab.com/rizoMulk
// @BasePath /
// @contact.name   Bakhodir Yashin Mansur
// @contact.email  phapp0224mb@gmail.com
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	logrus := logrus_log.GetLogger()
	cfg := config.Config()
	handlers := handler.NewHandler(logrus, cfg)
	app := handlers.InitRoutes()
	port := fmt.Sprintf(":%d", cfg.ServerPort)
	app.Run(port)
}
