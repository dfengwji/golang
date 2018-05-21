package api

import (
	"bytes"
	"eas.cloud/core"
	"eas.cloud/model"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	API_Login     = 0
	API_Admin     = 1
	API_Upload    = 2
)

const (
	StatusSuccess         = 0
	StatusNoAdmin         = 1
	StatusBindError       = 2
	StatusNotExisted      = 3
	StatusAddException    = 4
	StatusRemoveException = 5
	StatusUpdateException = 6
	StatusApplyError      = 7
	StatusApplyNotMatch   = 8
)

type ResponseInfo struct {
	Module  uint8       `json:"mod" binding:"required"`
	Option  uint8       `json:"opt" binding:"required"`
	Status  uint8       `json:"sta" binding:"required"`
	Message string      `json:"msg" binding:"required"`
	Data    interface{} `json:"data"`
}

func GetJson(req *http.Request) string {
	result, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return ""
	} else {
		return bytes.NewBuffer(result).String()
	}
}

func InterfaceHandler(uri string, group *gin.RouterGroup) {
	group.POST(uri, func(context *gin.Context) {
		raw := GetJson(context.Request)
		core.Log.Info("API AdminHandler..." + raw)
		var response ResponseInfo
		response.Module = API_Admin

		if !gjson.Valid(raw) {
			response.Status = StatusBindError
			response.Message = "json format error!!!"
			context.JSON(http.StatusBadRequest, response)
			core.Log.Info("api admin::parser error = " + raw)
			return
		}
		result := gjson.Parse(raw)
		module := uint8(result.Get("mod").Uint())
		option := uint8(result.Get("opt").Uint())

		if module == API_Admin {
			dealAdmin(option, result, context)
		} 
	})
}


