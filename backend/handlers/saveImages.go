package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/anabiozz/goods/backend/common"
	"github.com/anabiozz/goods/backend/models"
	"github.com/anabiozz/logger"
)

// SaveImagesHandler ...
func SaveImagesHandler(env *common.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		respBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			logger.Error(err)
			return
		}

		var images []*models.Image
		json.Unmarshal(respBody, &images)

		err = env.DB.SaveImages(images)
		if err != nil {
			logger.Error(err)
			return
		}
	})
}
