package main

import (
	_ "SwaggerDemo/docs"
	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB = nil
var err error

// @title gin+gorm crud 测试swagger（必填）
// @version 1.0 （必填）
// @description gin+gorm crud 测试swagger
// @license.name Apache 2.0
// @contact.name go-swagger帮助文档
// @contact.url https://github.com/swaggo/swag/blob/master/README_zh-CN.md
// @host localhost:8000
// @BasePath /
func main() {
	dsn := "root:hanxu@tcp(127.0.0.1:3307)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err.Error())
	}
	db.AutoMigrate(&Post{})

	server := gin.Default()
	server.GET("/posts", Posts)
	server.GET("/posts/:id", Show)
	server.POST("/posts", Store)
	//server.PATCH("/posts/:id", Update)
	server.DELETE("/posts/:id", Delete)

	//开启swagger
	server.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	server.Run(":8000")
}
