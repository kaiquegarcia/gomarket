package storage_test

import (
	"context"
	"gomarket/cmd"
	"gomarket/pkg/ctx"
	"gomarket/pkg/storage"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func writeTestFiles(
	t *testing.T,
	app cmd.Application,
	testFiles map[string]string,
) {
	for path, content := range testFiles {
		if content == "" {
			continue
		}

		fullpath := app.StorageDirectory() + path
		err := os.WriteFile(fullpath, []byte(content), os.ModePerm)
		if err != nil {
			t.Errorf("could not write test file %s", fullpath)
			t.Fail()
		}
	}
}

func deleteTestFiles(
	t *testing.T,
	app cmd.Application,
	testFiles map[string]string,
) {
	for path := range testFiles {
		fullpath := app.StorageDirectory() + path
		err := os.Remove(fullpath)
		if err != nil {
			t.Errorf("could not delete test file %s", fullpath)
			t.Fail()
		}
	}
}

func Test_JsonStorage_Read(t *testing.T) {
	app := cmd.NewApp()
	contxt := ctx.CtxWithApp(context.Background(), app)
	js := storage.NewJsonStorage(contxt)
	type Example struct {
		Field string `json:"name"`
	}
	testFiles := map[string]string{
		"test__read_and_decode.json":     "{\"name\":\"just an example\"}",
		"test__read_but_not_decode.json": "{\"name\":\"just an example\"}",
	}
	writeTestFiles(t, app, testFiles)

	t.Run("should read and decode json file", func(t *testing.T) {
		// Arrange
		path := "test__read_and_decode.json"

		// Act
		var output Example
		err := js.Read(path, &output)

		// Assert
		assert.Nil(t, err, "error response should be nil")
		assert.Equal(t, "just an example", output.Field, "the decoded JSON field should be 'just an example'")
	})

	t.Run("should read but fail on decode", func(t *testing.T) {
		// Arrange
		path := "test__read_but_not_decode.json"
		type IntExample struct {
			Field int `json:"name"`
		}

		// Act
		var output IntExample
		err := js.Read(path, &output)

		// Assert
		assert.Equal(t, "json: cannot unmarshal string into Go struct field IntExample.name of type int", err.Error(), "error response should be related to struct field")
		assert.Equal(t, 0, output.Field, "the undecoded JSON field should be 0")
	})

	t.Run("should fail on read", func(t *testing.T) {
		// Arrange
		path := "test__not_readable.json"

		// Act
		var output Example
		err := js.Read(path, &output)

		// Assert
		assert.True(t, os.IsNotExist(err), "error response should be related to file not found")
		assert.Equal(t, "", output.Field, "the undecoded JSON field should be ''")
	})

	deleteTestFiles(t, app, testFiles)
}

func Test_JsonStorage_Write(t *testing.T) {
	app := cmd.NewApp()
	contxt := ctx.CtxWithApp(context.Background(), app)
	js := storage.NewJsonStorage(contxt)
	type Example struct {
		Field string `json:"name"`
	}
	testFiles := map[string]string{
		"test__write_override.json":   "{\"name\":\"just an example\"}",
		"test__encode_and_write.json": "", // delete only
	}
	writeTestFiles(t, app, testFiles)

	t.Run("should encode and write json file", func(t *testing.T) {
		// Arrange
		path := "test__encode_and_write.json"
		input := Example{Field: "ok...?"}

		// Act
		err := js.Write(path, input)

		// Assert
		var output Example
		readErr := js.Read(path, &output)
		assert.Nil(t, err, "error response from Write should be nil")
		assert.Nil(t, readErr, "error response from Read should be nil")
		assert.Equal(t, "ok...?", output.Field, "written field should be 'ok...?'")
	})

	t.Run("should encode and override previously written json file", func(t *testing.T) {
		// Arrange
		path := "test__write_override.json"
		input := Example{Field: "ok...?"}

		// Act
		err := js.Write(path, input)

		// Assert
		var output Example
		readErr := js.Read(path, &output)
		assert.Nil(t, err, "error response from Write should be nil")
		assert.Nil(t, readErr, "error response from Read should be nil")
		assert.Equal(t, "ok...?", output.Field, "written field should be 'ok...?'")
	})

	deleteTestFiles(t, app, testFiles)
}

func Test_JsonStorage_Delete(t *testing.T) {
	app := cmd.NewApp()
	contxt := ctx.CtxWithApp(context.Background(), app)
	js := storage.NewJsonStorage(contxt)
	type Example struct {
		Field string `json:"name"`
	}
	testFiles := map[string]string{
		"test__delete.json": "{\"name\":\"just an example\"}",
	}
	writeTestFiles(t, app, testFiles)

	t.Run("should delete file", func(t *testing.T) {
		// Arrange
		path := "test__delete.json"

		// Act
		err := js.Delete(path)

		// Assert
		_, readErr := os.ReadFile(app.StorageDirectory() + path)
		assert.Nil(t, err, "error response from Delete should be nil")
		assert.True(t, os.IsNotExist(readErr), "error response from ReadFile should be not exists")
	})

	t.Run("should not delete unexistent file", func(t *testing.T) {
		// Arrange
		path := "test__not_delete_by_not_found.json"

		// Act
		err := js.Delete(path)

		// Assert
		assert.True(t, os.IsNotExist(err), "error response from Delete should be not exists")
	})
}
