package main

import (
	"My-Exercise/global"
	"My-Exercise/model/entity"
	"My-Exercise/router"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	initDB()
	initRouter()
}

func initRouter() {
	r := gin.Default()
	router.RegisterRouter(r)
	err := r.Run(":9000")
	if err != nil {
		panic(err)
	}
}

func initDB() {
	dsn := "root:nike5510@tcp(124.221.123.87)/my_exercise?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}
	global.DB = db
	syncTable(db)
}

func syncTable(db *gorm.DB) {
	err := db.AutoMigrate(
		&entity.Problem{},
		&entity.Category{},
		&entity.Submit{},
		&entity.User{},
	)
	if err != nil {
		panic(err)
	}
}
