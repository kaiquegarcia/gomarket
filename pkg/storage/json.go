package storage

import (
	"context"
	"gomarket/pkg/ctx"
)

type JsonStorage interface {
	// Read will try to retrieve the data stored on "{PROGRAM_DIR}storage/{path}" and decode into {dest} using json.Unmarshal.
	Read(path string, dest interface{}) error
	// Write will try to encode {data} using json.Marshal and store it on "{PROGRAM_DIR}storage/{path}"
	Write(path string, data interface{}) error
}

type jsonStorage struct {
	basePath string
}

// NewJsonStorage initializes an implementation of JsonStorage interface
func NewJsonStorage(
	contxt context.Context,
) JsonStorage {
	app := ctx.AppFromCtx(contxt)

	return jsonStorage{
		basePath: app.RootDirectory() + "storage" + app.Separator(),
	}
}

func (js jsonStorage) Read(path string, dest interface{}) error {
	// TODO
}

func (js jsonStorage) Write(path string, data interface{}) error {
	// TODO
}
