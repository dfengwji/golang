package auth

import (
	"github.com/gin-gonic/gin"
	"eas.cloud/external"
	"eas.cloud/core"
)

func ExternalHandler(router *gin.Engine) {
	core.Log.Info("ExternalHandler...")
	group := router.Group("/extra")
	{
		external.ExternalHandler("/orders",group)
		external.ExternalHandler("/save",group)
	}
}


