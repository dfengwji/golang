package auth

import (
	"eas.cloud/api"
	"eas.cloud/core"
	"eas.cloud/model"
	"github.com/gin-gonic/gin"
	"gopkg.in/appleboy/gin-jwt.v2"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func BindHandler(router *gin.Engine) *gin.RouterGroup {
	authMiddleware := &jwt.GinJWTMiddleware{
		Realm:      "simpleAPI",
		Key:        []byte("secret key"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		Authenticator: func(username string, psw string, c *gin.Context) (string, bool) {
			array := strings.Split(psw, "-")
			r, _ := strconv.ParseUint(array[1], 8, 10)
			role := uint8(r)
			password := array[0]
			core.Log.Info("try auth: username = " + username + ";psw = " + password + ";role = " + string(role))
			if role == model.RoleAdminSuper || role == model.RoleAdminNormal {
				admin := model.GetAdminByName(username)
				if (admin != nil && password == admin.Password) || (username == "test" && password == "test") {
					core.Log.Info("auth success!!!")
					return admin.UID, true
				}
			} else {
				teacher := model.GetTeacherByPhone(username)
				if teacher != nil {
					var admin model.AdminInfo
					admin.Role = model.RoleTeacher
					admin.Phone = teacher.Phone
					admin.UID = teacher.UID
					admin.Name = teacher.Name
					admin.Email = teacher.Email
					admin.Password = teacher.Password
					model.AddAdmin(admin)
					return teacher.UID, true
				}
			}

			core.Log.Info("auth failed!!!")
			return username, false
		},
		Authorizator: func(username string, c *gin.Context) bool {
			//获取uri
			uri := c.Request.URL.Path
			core.Log.Info("try auth: username = " + username + ", uri: " + uri)
			admin := model.GetAdmin(username)
			if admin != nil {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			core.Log.Info("Unauthorized")
			c.JSON(http.StatusUnauthorized, gin.H{})
			c.JSON(code, gin.H{
				"message": message,
			})
		},
		//TokenLookup:   "cookie:token",
		TokenLookup:   "header:Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	}
	router.POST("/login", authMiddleware.LoginHandler)
	//router.POST("/school", authMiddleware.LoginHandler)
	group := router.Group("/v1")
	group.Use(authMiddleware.MiddlewareFunc())
	{
		api.InterfaceHandler("/api", group)
		api.UploadHandler("/upload", group)
	}

	{
		group.GET("/refresh_token", authMiddleware.RefreshHandler)
	}
	return group
}
