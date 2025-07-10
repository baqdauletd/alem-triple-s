package handlers

import (
	"log"
	"net/http"
	"strings"
	"time"
	"triple-s/models"
	"triple-s/helpers"
)

func CreateBucket(w http.ResponseWriter, r *http.Request) {
	bucketName := strings.TrimPrefix(r.URL.Path, "/")

	if err := validateBucketName(bucketName); err != nil {
		log.Printf("Error validating bucket name %s: %v\n", bucketName, err)
		ErrXMLResponse(w, http.StatusBadRequest, "Invalid bucket name")
		return
	}

	bucketsData, err := helpers.ReadBucket()
	if err != nil {
		log.Printf("Error reading buckets file: %v\n", err)
		ErrXMLResponse(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	if SearchBucketIDX(bucketsData.Buckets, bucketName) != -1 {
		log.Printf("Bucket %s already exists\n", bucketName)
		ErrXMLResponse(w, http.StatusConflict, "Bucket already exists")
		return
	}

	newBucket := models.Bucket{
		Name:             bucketName,
		CreationTime:     time.Now().Format(time.RFC3339Nano),
		LastModifiedTime: time.Now().Format(time.RFC3339Nano),
		Status:           "Available",
	}
	bucketsData.Buckets = append(bucketsData.Buckets, newBucket)

	if err := helpers.WriteBucket(bucketsData); err != nil {
		log.Printf("Error writing buckets file: %v\n", err)
		ErrXMLResponse(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	if err := CreateBucketDir(bucketName); err != nil {
		log.Printf("Error creating bucket directory: %v\n", err)
		ErrXMLResponse(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	if err := InitObjectFile(bucketName); err != nil {
		log.Printf("Error initializing object file: %v\n", err)
		ErrXMLResponse(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	log.Printf("Bucket %s created successfully", bucketName)
	WriteXMLResponse(w, http.StatusOK, "Bucket created successfully")
}

func ListBuckets(w http.ResponseWriter, r *http.Request) {
	bucketData, err := helpers.ReadBucket()
	if err != nil {
		log.Printf("Error reading buckets file: %s", err)
		ErrXMLResponse(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	log.Println("Buckets listed successfully")
	WriteXMLResponse(w, http.StatusOK, bucketData)
}

func DeleteBucket(w http.ResponseWriter, r *http.Request) {
	bucketName := strings.TrimPrefix(r.URL.Path, "/")

	bucketsData, err := helpers.ReadBucket()
	if err != nil {
		log.Printf("Error reading buckets file: %v\n", err)
		ErrXMLResponse(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	bucketIDX := SearchBucketIDX(bucketsData.Buckets, bucketName)
	if bucketIDX == -1 {
		log.Printf("Bucket %s does not exist\n", bucketName)
		ErrXMLResponse(w, http.StatusNotFound, "Bucket does not exist")
		return
	}

	if !IsBucketEmpty(bucketName) {
		log.Printf("Bucket %s is not empty\n", bucketName)
		ErrXMLResponse(w, http.StatusConflict, "Bucket is not empty")
		return
	}

	if bucketsData.Buckets[bucketIDX].Status == "Marked for delete" {
		if err := RemoveBucketDir(bucketName); err != nil {
			log.Printf("Error removing bucket directory: %v\n", err)
			ErrXMLResponse(w, http.StatusInternalServerError, "Internal server error")
			return
		}

		bucketsData.Buckets = append(bucketsData.Buckets[:bucketIDX], bucketsData.Buckets[bucketIDX+1:]...)

		if err := helpers.WriteBucket(bucketsData); err != nil {
			log.Printf("Error writing buckets file: %v\n", err)
			ErrXMLResponse(w, http.StatusInternalServerError, "Internal server error")
			return
		}

		log.Printf("Bucket %s permanently deleted", bucketName)
		ErrXMLResponse(w, http.StatusNoContent, "Bucket permanently deleted")
		return
	}

	bucketsData.Buckets[bucketIDX].Status = "Marked for delete"
	bucketsData.Buckets[bucketIDX].LastModifiedTime = time.Now().Format(time.RFC3339Nano)

	if err := helpers.WriteBucket(bucketsData); err != nil {
		log.Printf("Error writing buckets file: %v\n", err)
		ErrXMLResponse(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	log.Printf("Bucket %s marked for delete (soft delete)", bucketName)
	WriteXMLResponse(w, http.StatusNoContent, "Bucket marked for delete")
}
