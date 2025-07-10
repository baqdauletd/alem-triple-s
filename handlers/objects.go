package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"triple-s/models"
	"triple-s/helpers"
)

func CreateObject(w http.ResponseWriter, r *http.Request) {
	bucketName, objectKey := ValidatePath(r.URL.Path)
	if bucketName == "" || objectKey == "" {
		log.Println("Invalid bucket or object key")
		ErrXMLResponse(w, http.StatusBadRequest, "Invalid bucket or object key")
		return
	}

	bucketPath := GetBucketPath(bucketName)
	if _, err := os.Stat(bucketPath); os.IsNotExist(err) {
		log.Printf("Bucket not found: %s\n", bucketName)
		ErrXMLResponse(w, http.StatusNotFound, "Bucket not found")
		return
	}

	objects, err := helpers.ReadObject(bucketName)
	fmt.Println(objects)
	if err != nil {
		log.Printf("Failed to read objects file for bucket %s: %v\n", bucketName, err)
		ErrXMLResponse(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	objectIndex := SearchObjectIDX(objects.Objects, objectKey)
	if objectIndex != -1 {
		objects.Objects = append(objects.Objects[:objectIndex], objects.Objects[objectIndex+1:]...)
	}

	objectPath := filepath.Join(bucketPath, objectKey)
	file, err := os.Create(objectPath)
	if err != nil {
		log.Printf("Failed to create object %s in bucket %s: %v\n", objectKey, bucketName, err)
		ErrXMLResponse(w, http.StatusInternalServerError, "Failed to create object")
		return
	}

	defer file.Close()

	_, err = io.Copy(file, r.Body)
	if err != nil {
		log.Printf("Failed to write object data for %s in bucket %s: %v\n", objectKey, bucketName, err)
		ErrXMLResponse(w, http.StatusInternalServerError, "Failed to write object data")
		return
	}

	newObj := models.Object{
		ObjectKey:    objectKey,
		ContentType:  r.Header.Get("Content-Type"),
		Size:         strconv.FormatInt(r.ContentLength, 10),
		LastModified: time.Now().Format(time.RFC3339),
	}

	if newObj.ContentType == "" {
		newObj.ContentType = "application/octet-stream"
	}

	objects.Objects = append(objects.Objects, newObj)
	fmt.Println(objects)
	err = helpers.WriteObject(bucketName, objects)
	if err != nil {
		log.Printf("Failed to update objects file for bucket %s: %v\n", bucketName, err)
		ErrXMLResponse(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	log.Printf("Object %s created successfully in bucket %s", objectKey, bucketName)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Object created successfully"))
}

func ListObjects(w http.ResponseWriter, r *http.Request) {
	bucketName := strings.TrimPrefix(r.URL.Path, "/")

	objData, err := helpers.ReadObject(bucketName)
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("Bucket not found: %s\n", bucketName)
			ErrXMLResponse(w, http.StatusNotFound, "Bucket not found")
		} else {
			log.Printf("Error reading objects file for bucket %s: %v\n", bucketName, err)
			ErrXMLResponse(w, http.StatusInternalServerError, "Internal server error")
		}
		return
	}

	log.Printf("Objects listed successfully for bucket %s", bucketName)
	WriteXMLResponse(w, http.StatusOK, objData)
}

func GetObject(w http.ResponseWriter, r *http.Request) {
	bucketName, objectKey := ValidatePath(r.URL.Path)
	if bucketName == "" || objectKey == "" {
		log.Println("Invalid bucket or object key")
		ErrXMLResponse(w, http.StatusBadRequest, "Invalid bucket or object key")
		return
	}

	bucketPath := GetBucketPath(bucketName)
	if _, err := os.Stat(bucketPath); os.IsNotExist(err) {
		log.Printf("Bucket not found: %s\n", bucketName)
		ErrXMLResponse(w, http.StatusNotFound, "Bucket not found")
		return
	}

	objectPath := GetObjectPath(bucketName, objectKey)
	file, err := os.Open(objectPath)
	if err != nil {
		log.Printf("Failed to open object %s in bucket %s: %v\n", objectKey, bucketName, err)
		ErrXMLResponse(w, http.StatusInternalServerError, "Failed to open object")
		return
	}
	defer file.Close()

	metadataType, err := GetMetadataType(bucketPath, objectKey)
	if err != nil {
		log.Printf("Failed to get metadata type for object %s in bucket %s: %v\n", objectKey, bucketName, err)
		ErrXMLResponse(w, http.StatusInternalServerError, "Failed to get metadata type")
		return
	}

	w.Header().Set("Content-Type", metadataType)
	_, err = io.Copy(w, file)
	if err != nil {
		log.Printf("Failed to read object %s in bucket %s: %v\n", objectKey, bucketName, err)
		ErrXMLResponse(w, http.StatusInternalServerError, "Failed to read object")
		return
	}

	log.Printf("Object %s read successfully in bucket %s", objectKey, bucketName)
}

func DeleteObject(w http.ResponseWriter, r *http.Request) {
	bucketName, objectKey := ValidatePath(r.URL.Path)
	if bucketName == "" || objectKey == "" {
		log.Println("Invalid bucket or object key")
		ErrXMLResponse(w, http.StatusBadRequest, "Invalid bucket or object key")
		return
	}

	bucketPath := GetBucketPath(bucketName)
	if _, err := os.Stat(bucketPath); os.IsNotExist(err) {
		log.Printf("Bucket not found: %s\n", bucketName)
		ErrXMLResponse(w, http.StatusNotFound, "Bucket not found")
		return
	}

	objectPath := GetObjectPath(bucketName, objectKey)
	if _, err := os.Stat(objectPath); os.IsNotExist(err) {
		log.Printf("Object not found: %s in bucket %s\n", objectKey, bucketName)
		ErrXMLResponse(w, http.StatusNotFound, "Object not found")
		return
	}

	objects, err := helpers.ReadObject(bucketName)
	if err != nil {
		log.Printf("Failed to read objects file for bucket %s: %v\n", bucketName, err)
		ErrXMLResponse(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	objectIDX := SearchObjectIDX(objects.Objects, objectKey)
	if objectIDX == -1 {
		log.Printf("Object %s not found in bucket %s\n", objectKey, bucketName)
		ErrXMLResponse(w, http.StatusNotFound, "Object not found")
		return
	}

	objects.Objects = append(objects.Objects[:objectIDX], objects.Objects[objectIDX+1:]...)
	err = helpers.WriteObject(bucketName, objects)
	if err != nil {
		log.Printf("Failed to update objects file for bucket %s: %v\n", bucketName, err)
		ErrXMLResponse(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	err = os.Remove(objectPath)
	if err != nil {
		log.Printf("Failed to delete object %s in bucket %s: %v\n", objectKey, bucketName, err)
		ErrXMLResponse(w, http.StatusInternalServerError, "Failed to delete object")
		return
	}

	log.Printf("Object %s deleted successfully from bucket %s\n", objectKey, bucketName)
	w.WriteHeader(http.StatusNoContent)
}
