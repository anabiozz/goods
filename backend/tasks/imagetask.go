package tasks

import (
	"fmt"
	"image/jpeg"
	"log"
	"os"

	"github.com/nfnt/resize"
)

// ImageResizeTask ...
type ImageResizeTask struct {
	BaseImageName      string
	ImageFileExtension string
}

// NewImageResizeTask .
func NewImageResizeTask(baseImageName string, imageFileExtension string) *ImageResizeTask {
	return &ImageResizeTask{BaseImageName: baseImageName, ImageFileExtension: imageFileExtension}
}

// Perform .
func (t *ImageResizeTask) Perform() {

	thumbImageFilePath := t.BaseImageName + "_thumb.jpg"
	fmt.Println("Creating new thumbnail at ", thumbImageFilePath)

	originalimagefile, err := os.Open(t.BaseImageName + t.ImageFileExtension)

	if err != nil {
		log.Println(err)
		return
	}

	img, err := jpeg.Decode(originalimagefile)

	if err != nil {
		log.Println("Encountered Error while decoding image file: ", err)
	}

	thumbImage := resize.Resize(270, 0, img, resize.Lanczos3)
	thumbImageFile, err := os.Create(thumbImageFilePath)

	if err != nil {
		log.Println("Encountered error while resizing image:", err)
	}

	jpeg.Encode(thumbImageFile, thumbImage, nil)

}
