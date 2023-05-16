package model

import "Food_Delivery3/common"

func ErrCannotSaveFile(err error) *common.AppError {
	return common.NewErrorResponse(err, "can not save file", err.Error())
}

func ErrFileIsNotImage(err error) *common.AppError {
	return common.NewErrorResponse(err, "file is not image", err.Error())

}
