package transport

import (
	todobiz "Food_Delivery3/module/res/business"
	"Food_Delivery3/module/res/model"
	todostorage "Food_Delivery3/module/res/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

func HandleCreateItem(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data model.Restaurant

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		data.Name = strings.TrimSpace(data.Name)

		// Setup dependencies
		storage := todostorage.NewMySQLStorage(db)
		biz := todobiz.NewCreateResItemBiz(storage)

		if err := biz.CreateNewRes(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": data.Id})
	}
}
