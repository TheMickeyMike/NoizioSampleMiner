package pkg

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	log "github.com/sirupsen/logrus"
)

// FileManager provides methods for files management
type FileManager struct {
	saveDirectory string
}

// NewFileManager provides FileManager
func NewFileManager(saveDirectory string) *FileManager {
	if err := CreateDirIfNotExist(saveDirectory); err != nil {
		log.Fatal("Can't create destination directory.", err)
	}
	return &FileManager{
		saveDirectory: saveDirectory,
	}
}

// SaveToFile creates file with given filename & data
func (fm *FileManager) SaveToFile(filename string, data []byte) error {
	dstPath := path.Join(fm.saveDirectory, fmt.Sprintf("%s.caf", filename))
	return ioutil.WriteFile(dstPath, data, 0644)
}

// CreateDirIfNotExist creates directory if not exist
func CreateDirIfNotExist(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.Mkdir(dir, 0755)
	}
	return nil
}
