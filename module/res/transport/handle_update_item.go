package transport

import (
	"Food_Delivery3/component"
	"Food_Delivery3/module/res/business"
	"Food_Delivery3/module/res/model"
	"Food_Delivery3/module/res/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func HandleUpdateRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var data model.RestaurantUpdate

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := storage.NewMySQLStorage(appCtx.GetMainDBConnection())
		biz := business.NewUpdateRestaurantBiz(store)

		if err := biz.UpdateRestaurant(c.Request.Context(), id, &data); err != nil {
			c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{"ok": 1})
	}
}
