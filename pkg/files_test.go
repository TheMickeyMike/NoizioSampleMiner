package core

import (
	"io/ioutil"
	"log"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveToFile(t *testing.T) {
	dir, err := ioutil.TempDir("", "TestEnsureDirectoryExist")
	if err != nil {
		log.Fatal(err)
	}

	defer os.RemoveAll(dir)

	fm := NewFileManager(dir)
	fileName := "example.file"
	fileContent := []byte("some content")
	err = fm.SaveToFile(fileName, fileContent)
	assert.NoError(t, err)

	data, err := ioutil.ReadFile(path.Join(dir, fileName))
	assert.NoError(t, err)
	assert.Equal(t, fileContent, data)

}

func TestEnsureDirectoryExistWhenNoDirectory(t *testing.T) {
	dir, err := ioutil.TempDir("", "TestEnsureDirectoryExist")
	if err != nil {
		log.Fatal(err)
	}

	defer os.RemoveAll(dir)

	expectedDirectory := path.Join(dir, "sounds")

	err = EnsureDirectoryExist(expectedDirectory)
	assert.NoError(t, err)
	assert.DirExists(t, expectedDirectory)

}

func TestEnsureDirectoryExistWhenDirectoryAlreadyExist(t *testing.T) {
	dir, err := ioutil.TempDir("", "TestEnsureDirectoryExist")
	if err != nil {
		log.Fatal(err)
	}

	defer os.RemoveAll(dir)

	expectedDirectory := path.Join(dir, "sounds")
	err = os.Mkdir(expectedDirectory, 0755)
	assert.NoError(t, err)

	err = EnsureDirectoryExist(expectedDirectory)
	assert.NoError(t, err)
	assert.DirExists(t, expectedDirectory)

}
