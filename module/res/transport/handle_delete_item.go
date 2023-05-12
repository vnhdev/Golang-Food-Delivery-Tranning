package transport

import (
	"Food_Delivery3/component"
	"Food_Delivery3/module/res/business"
	"Food_Delivery3/module/res/model"
	restaurantstorage "Food_Delivery3/module/res/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func HandleDeleteItem(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		var result model.RestaurantUpdate

		if err := c.ShouldBind(&result); err != nil {
			c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		store := restaurantstorage.NewMySQLStorage(appCtx.GetMainDBConnection())
		biz := business.NewDeleteResItemBiz(store)

		if err := biz.DeleteRestaurant(c.Request.Context(), id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
