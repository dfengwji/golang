package proxy

import (
	"eas.cloud/core"
	"go.uber.org/zap"
	"gopkg.in/mgo.v2/bson"
)

type Admin struct {
	UID bson.ObjectId `bson:"_id"`
	BaseDBInfo
	ID       uint32 `bson:"id"`
	Role     uint8  `bson:"role"`
	Password string `bson:"psw"`
	Phone    string `bson:"phone"`
	Email    string `bson:"email"`
}

func CreateAdmin(info Admin) (error, string) {
	err := noSql.C(TableAdmin).Insert(&info)
	if err != nil {
		core.Log.Info("create admin error", zap.String("name", info.Name),
			zap.Error(err))
		return err, ""
	}
	core.Log.Info("create admin...", zap.String("name", info.Name))
	return nil, info.UID.Hex()
}

func HasAdmin(id uint32) bool {
	var admin Admin
	err := noSql.C(TableAdmin).Find(bson.M{"id": id}).One(&admin)
	if err == nil {
		return true
	}
	core.Log.Error(err.Error())
	return false
}

func GetAdmin(uid string) *Admin {
	admin := new(Admin)
	objID := bson.ObjectIdHex(uid)
	err := noSql.C(TableAdmin).FindId(objID).One(&admin)
	if err == nil {
		return admin
	}
	core.Log.Error(err.Error())
	return nil
}

func GetAllAdmins() []Admin {
	var admins []Admin
	err := noSql.C(TableAdmin).Find(nil).All(&admins)
	if err == nil {
		return admins
	}
	core.Log.Error(err.Error())
	return nil
}

func RemoveAdmin(uid string) bool {
	objID := bson.ObjectIdHex(uid)
	err := noSql.C(TableAdmin).RemoveId(objID)
	if err == nil {
		return true
	}
	core.Log.Error(err.Error())
	return false
}
