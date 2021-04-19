package router

import (
	"log"
	"net/http"
	"time"

	control "aws-go-s3/controllers"

	"github.com/gorilla/mux"
)

func Router() {
	r := mux.NewRouter()
	r.HandleFunc("/create_bucket", control.CreateBucket).Methods("POST").Schemes("http", "https")

	srv := &http.Server{
		Addr:         "127.0.0.1:8000",
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
