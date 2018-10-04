package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/anabiozz/goods/backend/common/asyncq"
	"github.com/anabiozz/goods/backend/common/utility"
	"github.com/anabiozz/goods/backend/tasks"
)

// UploadImageForm ..
type UploadImageForm struct {
	PageTitle  string
	FieldNames []string
	Fields     map[string]string
	Errors     map[string]string
}

// DisplayUploadImageForm ..
func DisplayUploadImageForm(w http.ResponseWriter, r *http.Request, u *UploadImageForm) {
	RenderTemplate(w, "./templates/uploadimageform.html", u)
}

// ProcessUploadImage ..
func ProcessUploadImage(w http.ResponseWriter, r *http.Request, u *UploadImageForm) {

	shouldProcessThumbnailAsynchronously := false

	file, fileheader, err := r.FormFile("imagefile")

	if err != nil {
		log.Println("Encountered error when attempting to read uploaded file: ", err)
	}

	randomFileName := utility.GenerateUUID()

	if fileheader != nil {

		extension := filepath.Ext(fileheader.Filename)
		r.ParseMultipartForm(32 << 20)

		defer file.Close()

		fullImageFilePathWithoutExtension := "./static/images/graphics/full/" + randomFileName
		previewImageFilePathWithoutExtension := "./static/images/graphics/preview/" + randomFileName

		f, err := os.OpenFile(fullImageFilePathWithoutExtension+extension, os.O_WRONLY|os.O_CREATE, 0666)

		if err != nil {
			log.Println(err)
			return
		}

		defer f.Close()
		io.Copy(f, file)

		// Note: Moved the thumbnail generation logic (commented out code block below) to the
		// ImageResizeTask object in the tasks package.
		thumbnailResizeTask := tasks.NewImageResizeTask(fullImageFilePathWithoutExtension, previewImageFilePathWithoutExtension, extension)

		if shouldProcessThumbnailAsynchronously == true {

			asyncq.TaskQueue <- thumbnailResizeTask

		} else {

			thumbnailResizeTask.Perform()
		}

		m := make(map[string]string)
		m["thumbnailPath"] = strings.TrimPrefix(previewImageFilePathWithoutExtension, ".") + "_thumb.jpg"
		m["imagePath"] = strings.TrimPrefix(fullImageFilePathWithoutExtension, ".") + ".jpg"
		m["PageTitle"] = "Image Preview"

		RenderTemplate(w, "./templates/imagepreview.html", m)

	} else {
		w.Write([]byte("Failed to process uploaded file!"))
	}
}

// ValidateUploadImageForm ..
func ValidateUploadImageForm(w http.ResponseWriter, r *http.Request, u *UploadImageForm) {
	ProcessUploadImage(w, r, u)
}

// UploadImageHandler ...
func UploadImageHandler(w http.ResponseWriter, r *http.Request) {

	u := UploadImageForm{}
	u.Fields = make(map[string]string)
	u.Errors = make(map[string]string)
	u.PageTitle = "Upload Image"

	switch r.Method {

	case "GET":
		DisplayUploadImageForm(w, r, &u)
	case "POST":
		ValidateUploadImageForm(w, r, &u)
	default:
		DisplayUploadImageForm(w, r, &u)
	}

}
