package storage_helper

import (
	"github.com/google/uuid"
	"github.com/my-epoch/api-gateway/pkg/logger"
	"io/ioutil"
	"os"
)

func init() {
	if !IsFileOrDirectoryExist("storage") {
		if err := os.Mkdir("storage", os.ModePerm); err != nil {
			logger.Fatalf("cannot create storage directory in the root of app: %e", err)
		}
	}
}

func IsFileOrDirectoryExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}

func PutFileInStorage(fileId uuid.UUID, content []byte) error {
	f, err := os.Create("storage/" + fileId.String())
	if err != nil {
		return err
	}

	defer f.Close()

	if _, err = f.Write(content); err != nil {
		return err
	}
	return nil
}

func GetFileFromStorage(fileId uuid.UUID) ([]byte, error) {
	return ioutil.ReadFile("storage/" + fileId.String())
}
