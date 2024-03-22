package model

import (
	"github.com/ppolariss/paper-submission-review-system-go/user/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

var gormConfig = &gorm.Config{
	NamingStrategy: schema.NamingStrategy{
		SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
	},
}

func Init() {
	var err error
	switch config.Config.DbType {
	case "mysql":
		DB, err = gorm.Open(mysql.Open(config.Config.DbURL), gormConfig)
	case "sqlite":
		if config.Config.DbURL == "" {
			config.Config.DbURL = "data/sqlite.db"
		}
		DB, err = gorm.Open(sqlite.Open(config.Config.DbURL), gormConfig)
	default:
		panic("unknown db type")
	}
	if err != nil {
		panic(err)
	}
	err = DB.AutoMigrate(&User{}, &UserConferenceRole{})
	if err != nil {
		panic(err)
	}
}
