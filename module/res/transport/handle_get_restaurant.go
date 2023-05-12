package transport

import (
	"Food_Delivery3/component"
	"Food_Delivery3/module/res/business"
	storageres "Food_Delivery3/module/res/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func HandleGetRestaurant(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := storageres.NewMySQLStorage(appCtx.GetMainDBConnection())
		biz := business.FindSingleRestaurant(store)

		result, err := biz.FindRestaurantById(c.Request.Context(), map[string]interface{}{"id": id})

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": result})
	}
}
