package uploadprovider

import (
	"Food_Delivery3/common"
	"context"
)

type UploadProvider interface {
	SaveFileUploaded(context context.Context, data []byte, dst string) (*common.Image, error)
}
