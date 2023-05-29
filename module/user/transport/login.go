package transport

import (
	"Food_Delivery3/common"
	"Food_Delivery3/component"
	"Food_Delivery3/component/hasher"
	"Food_Delivery3/component/tokenprovider/jwt"
	"Food_Delivery3/module/user/userbiz"
	"Food_Delivery3/module/user/usermodel"
	"Food_Delivery3/module/user/userstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData usermodel.UserLogin

		if err := c.ShouldBind(&loginUserData); err != nil {
			c.JSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMainDBConnection()
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()

		biz := userbiz.NewLoginBusiness(store, tokenProvider, md5, appCtx.NewTokenConfig())
		account, err := biz.Login(c.Request.Context(), &loginUserData)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
