package main

import (
	"net/http"
	"os"
	"time"

	"github.com/anabiozz/goods/backend/common"
	"github.com/anabiozz/goods/backend/common/datastore"
	"github.com/anabiozz/goods/backend/handlers"
	"github.com/anabiozz/logger"
	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

func main() {

	logger.Init(os.Stdout, os.Stdout, os.Stderr, os.Stderr)

	db, err := datastore.NewDatastore(datastore.POSTGRES)
	if err != nil {
		logger.Error(err)
	}
	defer db.CloseDB()

	env := common.Env{DB: db}

	router := mux.NewRouter()
	router.Handle("/", handlers.IndexHandler(&env))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	srv := &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:8080",
		IdleTimeout:  5 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	logger.Fatal(srv.ListenAndServe())
}
