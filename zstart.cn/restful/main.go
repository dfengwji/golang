package main

import (
	"eas.cloud/api"
	"eas.cloud/auth"
	"eas.cloud/core"
	"eas.cloud/model"
	"eas.cloud/proxy"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	core.InitSettings()
	core.InitLogger()

	err := proxy.InitMongoDB()
	if nil != err {
		panic(err)
	}

	proxy.InitDB()
	model.InitTables()

	api.InitCourse()

	gin.DisableConsoleColor()

	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	gin.SetMode(gin.DebugMode)
	// 跨域调用
	conf := cors.Config{
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"content-type", "access-control-allow-headers", "origin", "authorization", "access-control-allow-origin"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	conf.AllowAllOrigins = true
	router.Use(cors.New(conf))

	// 需要认证
	auth.BindHandler(router)
	//不需要认证
	auth.ExternalHandler(router)

	path, _ := os.Getwd()
	router.StaticFS("/index", http.Dir(path+"/src/zstart.cn/web"))
	router.Run(":3000")
}
