package helpers

import (
	"log"
	"path/filepath"
	"triple-s/models"
	"triple-s/base"
)

func ReadBucket() (models.Buckets, error) {
	bucketPath := filepath.Join(base.Dir, "buckets.csv")
	log.Printf("Reading buckets metadate: %s\n", bucketPath)

	notes, err := ReadCSV(bucketPath)
	if err != nil {
		return models.Buckets{}, err
	}
	return NtoB(notes), nil
}

func WriteBucket(bucketData models.Buckets) error {
	bucketPath := filepath.Join(base.Dir, "buckets.csv")
	notes := BtoN(bucketData)

	return WriteCSV(bucketPath, base.BucketsHeader, notes)
}
