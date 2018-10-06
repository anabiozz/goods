package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/anabiozz/goods/backend/common"
	"github.com/anabiozz/logger"
)

// GetImagesHandler ..
func GetImagesHandler(env *common.Env) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		imageType := r.URL.Query()["imageType"]

		imageTypeInt, err := strconv.Atoi(imageType[0])
		if err != nil {
			logger.Info(err)
			return
		}

		images, err := env.DB.GetImagesByType(imageTypeInt)
		if err != nil {
			logger.Info(err)
			return
		}

		json.NewEncoder(w).Encode(images)
	})
}
