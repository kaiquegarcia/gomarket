package storage

import (
	"encoding/json"
	"os"
)

type JsonStorage interface {
	// Read will try to retrieve the data stored on "{PROGRAM_DIR}storage/{path}" and decode into {dest} using json.Unmarshal.
	Read(path string, dest interface{}) error
	// Write will try to encode {data} using json.Marshal and store it on "{PROGRAM_DIR}storage/{path}"
	Write(path string, data interface{}) error
	// Delete will try to delete the file stored on "{PROGRAM_DIR}storage/{path}"
	Delete(path string) error
}

type jsonStorage struct {
	basePath string
}

// NewJsonStorage initializes an implementation of JsonStorage interface
func NewJsonStorage(
	basePath string,
) JsonStorage {
	return jsonStorage{
		basePath: basePath,
	}
}

func (js jsonStorage) Read(path string, dest interface{}) error {
	bytes, err := os.ReadFile(js.basePath + path)
	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, dest)
}

func (js jsonStorage) Write(path string, data interface{}) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return os.WriteFile(js.basePath+path, bytes, os.ModePerm)
}

func (js jsonStorage) Delete(path string) error {
	return os.Remove(js.basePath + path)
}
