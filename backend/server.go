package main

import (
	"net/http"
	"os"
	"time"

	"github.com/anabiozz/goods/backend/api"
	"github.com/anabiozz/goods/backend/common"
	"github.com/anabiozz/goods/backend/common/datastore"
	"github.com/anabiozz/goods/backend/handlers"
	"github.com/anabiozz/logger"
	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

func main() {

	// asyncq.StartTaskDispatcher(1)

	logger.Init(os.Stdout, os.Stdout, os.Stderr, os.Stderr)

	db, err := datastore.NewDatastore(datastore.POSTGRES)
	if err != nil {
		logger.Error(err)
	}
	defer db.CloseDB()

	env := common.Env{DB: db}

	router := mux.NewRouter()
	// Handlers
	router.Handle("/", handlers.IndexHandler(&env))
	router.HandleFunc("/upload-images", handlers.UploadImageHandler)

	// API
	router.Handle("/api/get-graphics", api.GetGraphicsHandler(&env))

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	srv := &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:8080",
		IdleTimeout:  5 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	logger.Info("server run at localhost:8080")

	logger.Fatal(srv.ListenAndServe())
}
