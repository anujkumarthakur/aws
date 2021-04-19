package controllers

import (
	"aws-go-s3/functions"
	"aws-go-s3/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateBucket(w http.ResponseWriter, r *http.Request) {
	w = functions.SetDefaultHeaders(w)
	decoder := json.NewDecoder(r.Body)
	var newBucket models.Bucket
	err := decoder.Decode(&newBucket)
	if err != nil {
		fmt.Println("No Bucket Created!")
	}
	response := models.CreateNewBucket(newBucket.BucketName)
	json.NewEncoder(w).Encode(response)
}
