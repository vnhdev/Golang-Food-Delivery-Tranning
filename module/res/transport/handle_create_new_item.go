package transport

import (
	"Food_Delivery3/common"
	"Food_Delivery3/module/component"
	todobiz "Food_Delivery3/module/res/business"
	"Food_Delivery3/module/res/model"
	todostorage "Food_Delivery3/module/res/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func HandleCreateItem(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data model.RestaurantCreate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		data.Name = strings.TrimSpace(data.Name)

		// Setup dependencies
		storage := todostorage.NewMySQLStorage(appCtx.GetMainDBConnection())
		biz := todobiz.NewCreateResItemBiz(storage)

		if err := biz.CreateNewRestaurant(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
