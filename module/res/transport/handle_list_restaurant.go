package transport

import (
	"Food_Delivery3/common"
	"Food_Delivery3/component"
	todobiz "Food_Delivery3/module/res/business"
	"Food_Delivery3/module/res/model"
	todostorage "Food_Delivery3/module/res/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleListItem(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter model.Filter

		if err := c.ShouldBind(&filter); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Setup dependencies
		storage := todostorage.NewMySQLStorage(appCtx.GetMainDBConnection())
		biz := todobiz.NewListRestaurantBiz(storage)

		result, err := biz.ListRestaurant(c.Request.Context(), &filter, &paging)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
