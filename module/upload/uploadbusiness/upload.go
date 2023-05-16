package uploadbusiness

import (
	"Food_Delivery3/common"
	"Food_Delivery3/component/uploadprovider"
	"Food_Delivery3/module/upload/uploadmodel"
	"bytes"
	"context"
	"fmt"
	"image"
	"io"
	"log"
	"path/filepath"
	"strings"
	"time"
)

type CreateImageStorage interface {
	CreateImage(ctx context.Context, data *common.Image) error
}

type uploadBiz struct {
	provider  uploadprovider.UploadProvider
	itemStore CreateImageStorage
}

func NewUploadBiz(provider uploadprovider.UploadProvider, imageStore CreateImageStorage) *uploadBiz {
	return &uploadBiz{
		provider:  provider,
		itemStore: imageStore,
	}
}
func (biz *uploadBiz) Upload(ctx context.Context, data []byte, folder, fileName string) (*common.Image, error) {
	fileBytes := bytes.NewBuffer(data)

	w, h, err := getImageDimension(fileBytes)

	if err != nil {
		return nil, uploadmodel.ErrFileIsNotImage(err)
	}

	if strings.TrimSpace(folder) == "" {
		folder = "images"
	}

	fileExt := filepath.Ext(fileName)
	fileName = fmt.Sprintf("%s-%v%s", fileNameWithoutExtSliceNotation(fileName), time.Now().Nanosecond(), fileExt)

	img, err := biz.provider.SaveFileUploaded(ctx, data, fmt.Sprintf("%s/%s", folder, fileName))

	if err != nil {
		return nil, uploadmodel.ErrCannotSaveFile(err)
	}

	img.Width = w
	img.Height = h
	img.CloudName = "s3"
	img.Extension = fileExt

	return img, nil
}

func fileNameWithoutExtSliceNotation(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}

// other way
func fileNameWithoutExtTrimSuffix(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func getImageDimension(reader io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(reader)

	if err != nil {
		log.Println("Err====>", err)
		return 0, 0, err
	}

	return img.Width, img.Height, err
}
