package main

import (
	"My-Exercise/global"
	"My-Exercise/model/entity"
	"My-Exercise/router"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	initDB()
	InitRedis()
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
		&entity.ProblemCategory{},
		&entity.TestCase{},
	)
	if err != nil {
		panic(err)
	}
}

func InitRedis() {
	global.RDB = redis.NewClient(&redis.Options{
		Addr:     "124.221.123.87:6379",
		Password: "nike5510",
		DB:       0, // use default DB
	})
}
