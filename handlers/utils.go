package handlers

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
	"triple-s/models"
	"triple-s/base"
	"triple-s/helpers"
)

func validateBucketName(name string) error {
	if len(name) < 3 || len(name) > 63 {
		return fmt.Errorf("bucket name must be between 3 and 63 characters")
	}

	if strings.ContainsAny(name, ":*?\"<>|") {
		return fmt.Errorf("bucket name contains invalid characters")
	}

	if strings.Contains(name, "--") {
		return fmt.Errorf("bucket name cannot contain consecutive hyphens")
	}

	if strings.Contains(name, "..") {
		return fmt.Errorf("bucket name cannot contain consecutive periods")
	}

	if strings.HasPrefix(name, "-") || strings.HasSuffix(name, "-") {
		return fmt.Errorf("bucket name cannot start or end with a hyphen")
	}

	if net.ParseIP(name) != nil {
		return fmt.Errorf("bucket name cannot be an IP address")
	}

	return nil
}

func SearchBucketIDX(buckets []models.Bucket, name string) int {
	for i, bucket := range buckets {
		if bucket.Name == name {
			return i
		}
	}
	return -1
}

func SearchObjectIDX(objects []models.Object, name string) int {
	for i, object := range objects {
		if object.ObjectKey == name {
			return i
		}
	}
	return -1
}

func CreateBucketDir(bucketName string) error {
	dirPath := filepath.Join(base.Dir, bucketName)
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create bucket directory: %w", err)
	}
	return nil
}

func RemoveBucketDir(bucketName string) error {
	dirPath := filepath.Join(base.Dir, bucketName)
	if err := os.RemoveAll(dirPath); err != nil {
		return fmt.Errorf("failed to remove bucket directory: %w", err)
	}
	return nil
}

func ValidatePath(path string) (bucketName, objectKey string) {
	trimmedPath := strings.TrimPrefix(path, "/")
	parts := strings.SplitN(trimmedPath, "/", 2)
	bucketName = parts[0]
	if len(parts) > 1 {
		objectKey = parts[1]
	}
	return
}

func GetBucketPath(bucketName string) string {
	return filepath.Join(base.Dir, bucketName)
}

func GetObjectPath(bucketName, objectKey string) string {
	return filepath.Join(GetBucketPath(bucketName), objectKey)
}

func GetMetadataType(bucketPath, objectPath string) (string, error) {
	metadataPath := filepath.Join(bucketPath, "objects.csv")
	notes, err := helpers.ReadCSV(metadataPath)
	if err != nil {
		return "", fmt.Errorf("failed to read metadata: %w", err)
	}

	for _, note := range notes {
		if note[0] == objectPath {
			return note[1], nil
		}
	}
	return "", fmt.Errorf("metadata type not found")
}

func IsBucketEmpty(bucketName string) bool {
	bucketPath := GetBucketPath(bucketName)
	objectsMetadataPath := filepath.Join(bucketPath, "objects.csv")

	if _, err := os.Stat(objectsMetadataPath); os.IsNotExist(err) {
		return true
	}

	objects, err := helpers.ReadCSV(objectsMetadataPath)
	if err != nil {
		fmt.Printf("Error reading objects metadata: %v\n", err)
		return false
	}

	return len(objects) == 0
}
