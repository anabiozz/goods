package tasks

import (
	"fmt"
	"image/jpeg"
	"os"

	"github.com/anabiozz/logger"
	"github.com/nfnt/resize"
)

// ImageResizeTask ...
type ImageResizeTask struct {
	FullImageName      string
	PreviewImageName   string
	ImageFileExtension string
}

// NewImageResizeTask .
func NewImageResizeTask(fullImageName, previewImageName, imageFileExtension string) *ImageResizeTask {
	return &ImageResizeTask{
		FullImageName:      fullImageName,
		PreviewImageName:   previewImageName,
		ImageFileExtension: imageFileExtension,
	}
}

// Perform .
func (t *ImageResizeTask) Perform() {

	thumbImageFilePath := t.PreviewImageName + "_thumb.jpg"
	fmt.Println("Creating new thumbnail at ", thumbImageFilePath)

	originalimagefile, err := os.Open(t.FullImageName + t.ImageFileExtension)

	if err != nil {
		logger.Error(err)
		return
	}

	img, err := jpeg.Decode(originalimagefile)

	if err != nil {
		logger.Error("Encountered Error while decoding image file: ", err)
	}

	thumbImage := resize.Resize(270, 0, img, resize.Lanczos3)
	thumbImageFile, err := os.Create(thumbImageFilePath)

	if err != nil {
		logger.Error("Encountered error while resizing image:", err)
	}

	jpeg.Encode(thumbImageFile, thumbImage, nil)
}
