package main

import (
	"Food_Delivery3/component"
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
	fmt.Println(db)
	v1 := router.Group("/v1")
	{
		res := v1.Group("/res")
		{
			res.PUT("/:id", todotrpt.HandleUpdateRestaurant(appCtx))
			res.POST("", todotrpt.HandleCreateItem(appCtx))
			res.GET("", todotrpt.HandleListItem(appCtx))
			res.GET("/:id", todotrpt.HandleListItem(appCtx))
			res.DELETE("/:id", todotrpt.HandleDeleteItem(appCtx))
		}
	}
	router.Run()
}

//
//// EditRes restaurants by ID
//func editResbyId(db *gorm.DB) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		id, err := strconv.Atoi(c.Param("id"))
//		var data Restaurant
//
//		if err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//			return
//		}
//
//		if err := c.ShouldBind(&data); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//			return
//		}
//
//		if err := db.Where("id = ?", id).Updates(&data).Error; err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//			return
//		}
//		c.JSON(http.StatusOK, gin.H{"data": true})
//
//	}
//}
//
//// ListRes restaurants by ID
//func listResById(db *gorm.DB) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		id, err := strconv.Atoi(c.Param("id"))
//		var data Restaurant
//
//		if err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": "Id isn't exist"})
//			return
//		}
//
//		if err := c.ShouldBind(&data); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//			return
//		}
//
//		if err := db.Where("id = ?", id).First(&data).Error; err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//			return
//		}
//		c.JSON(http.StatusOK, gin.H{"data": data})
//	}
//
//}
//
//// ListRes restaurants
//func listRes(db *gorm.DB) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		var data []Restaurant
//
//		if err := c.ShouldBind(&data); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//			return
//		}
//
//		if err := db.Find(&data).Error; err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error})
//			return
//		}
//		c.JSON(http.StatusOK, gin.H{"data": data})
//
//	}
//}
//
//// DeleteRes restaurant
//func deleteRes(db *gorm.DB) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		id, err := strconv.Atoi(c.Param("id"))
//
//		if err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//			return
//		}
//
//		if err := db.Table(Restaurant{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//			return
//		}
//		c.JSON(http.StatusOK, gin.H{"data": true})
//	}
//}
//
//// CreateRes restaurant
//func createRes(db *gorm.DB) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		var data Restaurant
//
//		if err := c.ShouldBind(&data); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//			return
//		}
//
//		if err := db.Create(&data).Error; err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//			return
//		}
//		c.JSON(http.StatusOK, gin.H{"data": data})
//	}
//}
//
//type Restaurant struct {
//	Id   int    `json:"id" gorm:"column:id"`
//	Name string `json:"name" gorm:"column:name"`
//	Addr string `json:"addr" gorm:"column:addr"`
//}
//
//func (Restaurant) TableName() string {
//	return "restaurants"
//}
