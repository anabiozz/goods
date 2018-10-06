package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/anabiozz/goods/backend/common/asyncq"
	"github.com/anabiozz/goods/backend/common/utility"
	"github.com/anabiozz/goods/backend/tasks"
)

// UploadImageForm ..
type UploadImageForm struct {
	FieldNames []string
	Fields     map[string]string
	Errors     map[string]string
}

// DisplayUploadImageForm ..
func DisplayUploadImageForm(w http.ResponseWriter, r *http.Request, uploadForm *UploadImageForm) {
	RenderTemplate(w, "./templates/uploadimageform.html", uploadForm)
}

// ProcessUploadImage ..
func ProcessUploadImage(w http.ResponseWriter, r *http.Request, uploadForm *UploadImageForm) {

	shouldProcessThumbnailAsynchronously := false

	r.ParseMultipartForm(32 << 20)
	fileHeaders := r.MultipartForm.File["imagefiles"]
	var images []map[string]string

	for _, fileHeader := range fileHeaders {

		file, err := fileHeader.Open()

		if err != nil {
			log.Println("Encountered error when attempting to read uploaded file: ", err)
		}

		randomFileName := utility.GenerateUUID()

		if fileHeader != nil {

			extension := filepath.Ext(fileHeader.Filename)

			defer file.Close()

			fullImageFilePathWithoutExtension := "./static/images/full/" + randomFileName
			previewImageFilePathWithoutExtension := "./static/images/preview/" + randomFileName

			f, err := os.OpenFile(fullImageFilePathWithoutExtension+extension, os.O_WRONLY|os.O_CREATE, 0666)

			if err != nil {
				log.Println(err)
				return
			}
			defer f.Close()
			io.Copy(f, file)

			thumbnailResizeTask := tasks.NewImageResizeTask(fullImageFilePathWithoutExtension, previewImageFilePathWithoutExtension, extension)

			if shouldProcessThumbnailAsynchronously == true {
				asyncq.TaskQueue <- thumbnailResizeTask

			} else {
				thumbnailResizeTask.Perform()
			}

			u := make(map[string]string)
			u["Name"] = randomFileName
			images = append(images, u)

		} else {
			w.Write([]byte("Failed to process uploaded file!"))
		}
	}

	RenderTemplate(w, "./templates/imagepreview.html", images)
}

// ValidateUploadImageForm ..
func ValidateUploadImageForm(w http.ResponseWriter, r *http.Request, uploadForm *UploadImageForm) {
	ProcessUploadImage(w, r, uploadForm)
}

// UploadImageHandler ...
func UploadImageHandler(w http.ResponseWriter, r *http.Request) {

	uploadForm := UploadImageForm{}
	uploadForm.Fields = make(map[string]string)
	uploadForm.Errors = make(map[string]string)

	switch r.Method {

	case "GET":
		DisplayUploadImageForm(w, r, &uploadForm)
	case "POST":
		ValidateUploadImageForm(w, r, &uploadForm)
	default:
		DisplayUploadImageForm(w, r, &uploadForm)
	}

}
