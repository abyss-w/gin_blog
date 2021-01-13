package model

import (
	"github.com/abyss-w/gin_blog/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	DB  *gorm.DB
	err error
)

func InitDB() {
	DB, err = gorm.Open(mysql.Open(utils.DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{}, &Article{}, &Category{})
	if err != nil {
		panic(err)
	}

	rawDB, err := DB.DB()
	if err != nil {
		panic(err)
	}

	rawDB.SetMaxIdleConns(10)
	rawDB.SetMaxOpenConns(100)
	rawDB.SetConnMaxLifetime(10 * time.Second)
}
