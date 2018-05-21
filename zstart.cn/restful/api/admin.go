package api

import (
	"eas.cloud/core"
	"eas.cloud/model"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"go.uber.org/zap"
	"gopkg.in/appleboy/gin-jwt.v2"
	"net/http"
	"strconv"
)

const (
	OptAdminGetUser     = 1
	OptAdminSetPassword = 2
	OptAdminGetAll      = 3
	OptAdminAdd         = 4
	OptLogout           = 5
)

func AdminHandler(uri string, group *gin.RouterGroup) {
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

		claims := jwt.ExtractClaims(context)
		user := claims["id"]
		core.Log.Info("api admin::module = " + strconv.Itoa(int(module)) + ";opt=" + strconv.Itoa(int(option)))
		core.Log.Info("api admin:: ", zap.Any("body", result.Get("body")))

		admin := model.GetAdmin(user.(string))
		if admin == nil {
			response.Status = StatusNoAdmin
			response.Message = "not found the admin"
			context.JSON(http.StatusNotFound, response)
			core.Log.Warn("can not found the user!!!", zap.String("id", user.(string)))
			return
		}

		if option == OptAdminGetUser {
			handleUser(admin, &response, context)
		} else if option == OptAdminSetPassword {
			handleAdminPassword(admin.ID, &result, &response, context)
		} else if option == OptAdminGetAll {
			handleAdminGetAll(&response, context)
		} else if option == OptAdminAdd {
			handleAdminAdd(&result, &response, context)
		} else if option == OptLogout {
			handleLogout(&response, context)
		}
	})
}

func dealAdmin(option uint8, result gjson.Result, context *gin.Context) {
	var response ResponseInfo
	response.Module = API_Admin

	claims := jwt.ExtractClaims(context)
	user := claims["id"]
	core.Log.Info("api admin::opt = " + strconv.Itoa(int(option)))
	core.Log.Info("api admin:: ", zap.Any("body", result.Get("body")))

	admin := model.GetAdmin(user.(string))
	if admin == nil {
		response.Status = StatusNoAdmin
		response.Message = "not found the admin"
		context.JSON(http.StatusNotFound, response)
		core.Log.Warn("can not found the user!!!", zap.String("id", user.(string)))
		return
	}

	if option == OptAdminGetUser {
		handleUser(admin, &response, context)
	} else if option == OptAdminSetPassword {
		handleAdminPassword(admin.ID, &result, &response, context)
	} else if option == OptAdminGetAll {
		handleAdminGetAll(&response, context)
	} else if option == OptAdminAdd {
		handleAdminAdd(&result, &response, context)
	}
}

func handleUser(admin *model.AdminInfo, response *ResponseInfo, ctx *gin.Context) {
	response.Option = OptAdminGetUser
	response.Data = admin

	response.Status = StatusSuccess
	response.Message = ""
	ctx.JSON(http.StatusOK, response)
}

func handleAdminPassword(id uint32, result *gjson.Result, response *ResponseInfo, ctx *gin.Context) {
	response.Option = OptAdminSetPassword
	var psw struct {
		Old string `json:"old"`
		New string `json:"new"`
	}
	psw.New = result.Get("body.new").String()
	psw.Old = result.Get("body.old").String()

	model.UpdateAdmin(id, psw.New)

	response.Status = StatusSuccess
	response.Message = ""
	ctx.JSON(http.StatusOK, response)
}

func handleAdminGetAll(response *ResponseInfo, ctx *gin.Context) {
	response.Option = OptAdminGetAll
	response.Status = StatusSuccess
	response.Message = ""
	response.Data = model.Administrators
	ctx.JSON(http.StatusOK, response)
}

func handleAdminAdd(result *gjson.Result, response *ResponseInfo, ctx *gin.Context) {
	response.Option = OptAdminAdd
	var admin model.AdminInfo
	admin.ID = uint32(result.Get("body.id").Uint())
	admin.Email = result.Get("body.email").String()
	admin.Phone = result.Get("body.phone").String()
	admin.Role = uint8(result.Get("body.role").Uint())
	if admin.Role > model.RoleAdminNormal {
		admin.Role = model.RoleAdminNormal
	}
	admin.Password = result.Get("body.psw").String()
	admin.Name = result.Get("body.name").String()

	adminInfo, err := model.CreateAdmin(admin)
	if err != nil {
		response.Status = StatusAddException
		response.Message = "create admin failed"
		ctx.JSON(http.StatusOK, response)
		core.Log.Warn("create admin failed", zap.String("raw", result.String()))
		return
	}
	response.Data = adminInfo
	response.Status = StatusSuccess
	response.Message = ""
	ctx.JSON(http.StatusOK, response)
}

func handleLogout(response *ResponseInfo, ctx *gin.Context) {

}
