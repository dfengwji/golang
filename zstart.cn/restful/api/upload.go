package api

import (
	"eas.cloud/core"
	"eas.cloud/model"
	"eas.cloud/proxy"
	"fmt"
	"github.com/Luxurioust/excelize"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gopkg.in/appleboy/gin-jwt.v2"
	"net/http"
	"os"
	"strconv"
	"time"
)

const (
	OptSchoolAddAll   = 1
	OptTeacherAddAll  = 2
	OptStudentAddAll  = 3
	OptPracticeAddAll = 4
)

func UpdateHandle(router *gin.Engine)  {
	router.POST("/extra/upload", func(ctx *gin.Context) {
		var response ResponseInfo
		response.Module = API_Upload
		a1, _ := strconv.Atoi(ctx.PostForm("opt"))
		option := uint16(a1)
		a2, _ := strconv.Atoi(ctx.PostForm("mod"))
		module := uint16(a2)

		core.Log.Info("api upload::module = " + strconv.Itoa(int(module)) + ";opt=" + strconv.Itoa(int(option)))

		if option == OptSchoolAddAll {
			handleSchools(&response, ctx)
		} else if option == OptTeacherAddAll {
			handleTeachers(&response, ctx)
		} else if option == OptStudentAddAll {
			handleStudents(&response, ctx)
		} else if option == OptPracticeAddAll {
			handlePractices(&response, ctx)
		}
	})
}

func UploadHandler(uri string, group *gin.RouterGroup) {
	group.POST(uri, func(ctx *gin.Context) {
		var response ResponseInfo
		response.Module = API_Upload
		a1, _ := strconv.Atoi(ctx.PostForm("opt"))
		option := uint16(a1)
		a2, _ := strconv.Atoi(ctx.PostForm("mod"))
		module := uint16(a2)

		core.Log.Info("api upload::module = " + strconv.Itoa(int(module)) + ";opt=" + strconv.Itoa(int(option)))

		if option == OptSchoolAddAll {
			handleSchools(&response, ctx)
		} 
	})
}

func handleSchools(response *ResponseInfo, ctx *gin.Context) {
	response.Option = OptSchoolAddAll
	core.Log.Info("UploadHandler:" + ctx.Param("data"))
	//得到上传的文件
	file, header, err := ctx.Request.FormFile("file") //file这个是前端页面定义的名称
	if err != nil {
		ctx.String(http.StatusBadRequest, "Bad request")
		return
	}
	//文件的名称
	filename := header.Filename
	
	var school model.SchoolInfo
	added := make([]model.SchoolInfo, 0)
	excel, err := excelize.OpenReader(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		response.Status = StatusAddException
		response.Message = "school readException"
		ctx.JSON(http.StatusOK, response)
		core.Log.Warn("can not read the excel", zap.String("err", err.Error()))
		return
	}

	rows := excel.GetRows("Sheet1")
	for m, row := range rows {
		if m == 0 {
			continue
		}
		for index, colCell := range row {
			fmt.Print(colCell, "\t")
			switch index {
			case 0:
				x0, _ := strconv.Atoi(colCell)
				school.ID = uint32(x0)
			case 1:
				x1, _ := strconv.Atoi(colCell)
				school.Admin = uint32(x1)
			case 2:
				school.Name = colCell
			case 3:
				school.Master = colCell
			case 4:
				x2, _ := strconv.Atoi(colCell)
				school.Address.Province = uint16(x2)
			case 5:
				x3, _ := strconv.Atoi(colCell)
				school.Address.City = uint16(x3)
			case 6:
				x4, _ := strconv.Atoi(colCell)
				school.Address.Zone = uint16(x4)
			case 7:
				x5, _ := strconv.Atoi(colCell)
				school.Address.Town = uint16(x5)
			case 8:
				x6, _ := strconv.Atoi(colCell)
				school.Address.Street = uint16(x6)
			}

		}
		has := model.HasSchool(school.ID)
		if has {
			response.Status = StatusNotExisted
			response.Message = "school existed "
			ctx.JSON(http.StatusOK, response)
			return
		}
		_, err := model.CreateSchool(school)
		if err != nil {
			response.Status = StatusAddException
			response.Message = "school addException "
			ctx.JSON(http.StatusOK, response)
			core.Log.Warn("can not add the school", zap.String("err", err.Error()))
			return
		}
		added = append(added, school)
	}
	response.Status = StatusSuccess
	response.Message = ""
	response.Data = added
	ctx.JSON(http.StatusOK, response)
}
