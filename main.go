package main

import (
	"Food_Delivery3/component"
	"Food_Delivery3/middleware"
	todotrpt "Food_Delivery3/module/res/transport"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "root:secret@tcp(127.0.0.1:3307)/Restaurant?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	appCtx := component.NewAppContext(db)
	if err != nil {
		log.Fatal("Can't connect to DB Mysql:", err)
	}
	router := gin.Default()

	router.Use(middleware.Recover(appCtx))
	fmt.Println(db)
	v1 := router.Group("/v1")
	{
		res := v1.Group("/res")
		{
			res.POST("", todotrpt.HandleCreateItem(appCtx))
			res.GET("", todotrpt.HandleListItem(appCtx))
			res.GET("/:id", todotrpt.HandleGetRestaurant(appCtx))
			res.PUT("/:id", todotrpt.HandleUpdateRestaurant(appCtx))
			res.DELETE("/:id", todotrpt.HandleDeleteItem(appCtx))
		}
	}
	router.Run()
}
