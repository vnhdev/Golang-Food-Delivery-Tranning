package userbiz

import (
	"Food_Delivery3/common"
	"Food_Delivery3/module/user/usermodel"
	"context"
)

type RegisterStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}

type registerBusiness struct {
	RegisterStorage RegisterStorage
	Hasher          Hasher
}

func NewRegisterBusiness(RegisterStorage RegisterStorage, hasher Hasher) *registerBusiness {
	return &registerBusiness{
		RegisterStorage: RegisterStorage,
		Hasher:          hasher,
	}
}
func (business *registerBusiness) Register(ctx context.Context, data *usermodel.UserCreate) error {
	user, err := business.RegisterStorage.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if user != nil {
		return common.ErrEntityExisted(usermodel.EntityName, error(err))
	}

	salt := common.GenSalt(50)

	data.Password = business.Hasher.Hash(data.Password + salt)
	data.Salt = salt
	data.Role = "user"
	data.Status = 1

	if err := business.RegisterStorage.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}
	return nil
}
