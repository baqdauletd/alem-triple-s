package helpers

import (
	"log"
	"path/filepath"
	"triple-s/models"
	"triple-s/base"
)

func ReadObject(bucketName string) (models.Objects, error) {
	objectPath := filepath.Join(base.Dir, bucketName, "objects.csv")
	log.Printf("Reading objects metadata: %s\n", objectPath)

	notes, err := ReadCSV(objectPath)
	if err != nil {
		return models.Objects{}, err
	}
	return NtoO(notes), nil
}

func WriteObject(bucketName string, objectData models.Objects) error {
	objectPath := filepath.Join(base.Dir, bucketName, "objects.csv")
	notes := OtoN(objectData)

	return WriteCSV(objectPath, base.ObjectsHeader, notes)
}
