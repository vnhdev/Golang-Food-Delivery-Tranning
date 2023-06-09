package transport

import (
	"Food_Delivery3/common"
	"Food_Delivery3/component"
	"Food_Delivery3/component/hasher"
	"Food_Delivery3/module/user/userbiz"
	"Food_Delivery3/module/user/usermodel"
	"Food_Delivery3/module/user/userstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewRegisterBusiness(store, md5)

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
