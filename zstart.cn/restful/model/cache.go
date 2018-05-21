package model

import (
	"eas.cloud/core"
	"eas.cloud/proxy"
	"errors"
	"go.uber.org/zap"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type BaseInfo struct {
	UID  string `json:"uid"`
	Name string `json:"name"`
}

var Administrators 		[]AdminInfo

func InitTables() {
	admins := proxy.GetAllAdmins()
	for i := 0; i < len(admins); i++ {
		var info AdminInfo
		info.CopyAdmin(&admins[i])
		AddAdmin(info)
	}
}

//region Admin Function
func HasAdmin(id uint32) bool {
	length := len(Administrators)
	for i := 0; i < length; i++ {
		if Administrators[i].ID == id {
			return true
		}
	}
	return false
}

func GetAdminByName(identify string) *AdminInfo {
	length := len(Administrators)
	for i := 0; i < length; i++ {
		info := Administrators[i]
		//core.Log.Info(info.Name + "---"+identify)
		if (info.Name == identify) || (info.Email == identify) || (info.Phone == identify) {
			return &info
		}
	}
	return nil
}

func GetAdmin(uid string) *AdminInfo {
	length := len(Administrators)
	for i := 0; i < length; i++ {
		if Administrators[i].UID == uid {
			return &Administrators[i]
		}
	}
	return nil
}

func AddAdmin(info AdminInfo) error {
	Administrators = append(Administrators, info)
	//core.Log.Info("add admin....",zap.Uint32("id", info.ID),zap.String("name", info.Name),zap.Int("length",len(Administrators)))
	return nil
}

func CreateAdmin(info AdminInfo) (*AdminInfo, error) {
	id := proxy.GetAdminCount() + 1
	if id < 0 {
		return nil, errors.New("create failed that id error")
	}
	if id < 3000 {
		id = 3000
	}
	var admin1 proxy.Admin
	admin1.UID = bson.NewObjectId()
	admin1.ID = info.ID
	admin1.CreatedTime = time.Now()
	admin1.Name = info.Name
	admin1.Phone = info.Phone
	admin1.Role = info.Role
	admin1.Password = info.Password
	admin1.Email = info.Email
	err, uid := proxy.CreateAdmin(admin1)
	if err != nil {
		return nil, err
	}
	info.UID = uid
	Administrators = append(Administrators, info)
	//core.Log.Info("add admin....",zap.Uint32("id", info.ID),zap.String("name", info.Name),zap.Int("length",len(Administrators)))
	return &info, nil
}

func RemoveAdmin(id uint32) bool {
	length := len(Administrators)
	for i := 0; i < length; i++ {
		if Administrators[i].ID == id {
			Administrators = append(Administrators[:i], Administrators[i+1:]...)
			return true
		}
	}
	return false
}

func UpdateAdmin(id uint32, psw string) {

}

//endregion

