package storage_test

import (
	"context"
	"gomarket/cmd"
	"gomarket/internal/errs"
	"gomarket/pkg/ctx"
	"gomarket/pkg/storage"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Collection_Load(t *testing.T) {
	app := cmd.NewApp()
	contxt := ctx.CtxWithApp(context.Background(), app)
	js := storage.NewJsonStorage(contxt)
	testFiles := map[string]string{
		"collection_test_load.json": "{\"name\":\"test_load\",\"codes\":[1,2,3,4],\"next_code\":5}",
	}
	writeTestFiles(t, app, testFiles)

	t.Run("should instance collection for the first time", func(t *testing.T) {
		// Act
		collection, err := storage.NewCollection(js, "test_new")

		// Assert
		assert.Nil(t, err, "err should be nil")
		assert.Equal(t, 1, collection.GetNextCode(), "first NextCode should be 1")
	})

	t.Run("should load collection from storage", func(t *testing.T) {
		// Act
		collection, err := storage.NewCollection(js, "test_load")

		// Assert
		assert.Nil(t, err, "err should be nil")
		assert.Equal(t, 5, collection.GetNextCode(), "loaded NextCode should be 5")
	})

	deleteTestFiles(t, app, testFiles)
}

func Test_Collection_Get(t *testing.T) {
	app := cmd.NewApp()
	contxt := ctx.CtxWithApp(context.Background(), app)
	js := storage.NewJsonStorage(contxt)
	testFiles := map[string]string{
		"collection_test.json":            "{\"name\":\"test\",\"codes\":[1,2],\"next_code\":3}",
		"collection_test_registry_1.json": "{\"code\":1,\"name\":\"John\"}",
	}
	type Example struct {
		Code int    `json:"code"`
		Name string `json:"name"`
	}
	writeTestFiles(t, app, testFiles)
	collection, err := storage.NewCollection(js, "test")
	if err != nil {
		t.Errorf("could not load collection: %s", err)
		t.FailNow()
	}

	t.Run("should get existent registries", func(t *testing.T) {
		// Act
		var data Example
		err := collection.Get(1, &data)

		// Assert
		assert.Nil(t, err, "err should be nil")
		assert.Equal(t, 1, data.Code, "retrieved data field 'code' should be 1")
		assert.Equal(t, "John", data.Name, "retrieved data field 'name' should be 'John'")
	})

	t.Run("should not get unexistent registries", func(t *testing.T) {
		// Act
		var data Example
		err := collection.Get(3, &data)

		// Assert
		assert.Equal(t, errs.RegistryNotFoundErr, err, "err should be RegistryNotFoundErr")
	})

	t.Run("should not get unexistent registry files", func(t *testing.T) {
		// Act
		var data Example
		err := collection.Get(2, &data)

		// Assert
		assert.True(t, os.IsNotExist(err), "err should be os.NotExist error")
	})

	deleteTestFiles(t, app, testFiles)
}

func Test_Collection_List(t *testing.T) {
	// TODO
}

func Test_Collection_DecodeRaw(t *testing.T) {
	// TODO
}

func Test_Collection_Save(t *testing.T) {
	// TODO
}

func Test_Collection_Delete(t *testing.T) {
	// TODO
}
