package handlers

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"triple-s/base"
)

func setupFileWithDefaultHeader(filePath string, header []string) error {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0o666)
	if err != nil {
		log.Println(err)
		return err
	}
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		log.Println(err)
		return err
	}

	if fileStat.Size() == 0 {
		_, err := file.WriteString(strings.Join(header, ",") + "\n")
		if err != nil {
			log.Println(err)
			return err
		}
	}

	return nil
}

func DirInit() error {
	err := os.MkdirAll(base.Dir, 0o666)
	if err != nil {
		return err
	}
	return setupFileWithDefaultHeader(filepath.Join(base.Dir, "buckets.csv"), base.BucketsHeader)
}

func InitObjectFile(bucketName string) error {
	err := os.MkdirAll(filepath.Join(base.Dir, bucketName), 0o666)
	if err != nil {
		return err
	}

	return setupFileWithDefaultHeader(filepath.Join(base.Dir, bucketName, "objects.csv"), base.ObjectsHeader)
}
