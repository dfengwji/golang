package external

import (
	"github.com/tidwall/gjson"
	"github.com/gin-gonic/gin"
	"net/http"
	"eas.cloud/model"
)

type ResponseInfo struct {
	Module  uint8 `json:"mod" binding:"required"`
	Option  uint8 `json:"opt" binding:"required"`
	Status  uint8 `json:"sta" binding:"required"`
	Message  string `json:"msg" binding:"required"`
	Data    interface{} `json:"data"`
}
func ExternalHandler(uri string, group *gin.RouterGroup) {
	group.POST(uri, func(context *gin.Context) {

	})
}

func OrderHandle(result *gjson.Result,ctx *gin.Context)  {
	//验证接口不需要验证时通否
	var response ResponseInfo
	response.Status = 0
	response.Message = ""
	response.Data = model.Courses
	ctx.JSON(http.StatusOK, response)
	return
}