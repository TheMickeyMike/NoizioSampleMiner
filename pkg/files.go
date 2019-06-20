package core

import (
	"io/ioutil"
	"os"
	"path"
)

// FileManager provides methods for files management
type FileManager struct {
	saveDirectory string
}

// NewFileManager provides FileManager
func NewFileManager(saveDirectory string) *FileManager {
	return &FileManager{
		saveDirectory: saveDirectory,
	}
}

// SaveToFile creates file with given filename & data
func (fm *FileManager) SaveToFile(filename string, data []byte) error {
	dstPath := path.Join(fm.saveDirectory, filename)
	return ioutil.WriteFile(dstPath, data, 0644)
}

// EnsureDirectoryExist creates directory if not exist
func EnsureDirectoryExist(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.Mkdir(dir, 0755)
	}
	return nil
}
