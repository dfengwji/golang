package proxy

import (
	"eas.cloud/core"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.uber.org/zap"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

const (
	TableAdmin    = "admins"
)

type BaseDBInfo struct {
	CreatedTime time.Time `bson:"created_at"`
	UpdatedTime time.Time `bson:"updated_at"`
	Name        string    `bson:"name"`
}

var dbSql *gorm.DB
var noSql *mgo.Database

func InitMongoDB() error {
	//mongodb://myuser:mypass@localhost:40001
	session, err := mgo.Dial(core.DB_URL)
	if err != nil {
		return err
	}

	noSql = session.DB(core.DB_NAME)
	tables, _ := noSql.CollectionNames()
	for i := 0; i < len(tables); i++ {
		core.Log.Info("no sql table name = " + tables[i])
	}

	return nil
}

func InitDB() {
	
}

func InitMysql() error {
	uri := core.DB_USER + ":" + core.DB_PASSWORD + "@tcp(" + core.DB_URL + ")/" + core.DB_NAME
	db, err := gorm.Open(core.DB_TYPE, uri)
	if err != nil {
		panic("failed to connect database!!!" + uri)
		return err
	}
	dbSql = db
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	dbSql.LogMode(true)

	core.Log.Info("connect database success!!!")
	initTeacherTable()
	return nil
}
