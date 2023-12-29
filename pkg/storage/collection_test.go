package storage_test

import (
	"encoding/json"
	"gomarket/cmd"
	"gomarket/internal/errs"
	"gomarket/pkg/storage"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Collection_Load(t *testing.T) {
	app := cmd.NewApp()
	js := storage.NewJsonStorage(app.StorageDirectory())
	testFiles := map[string]string{
		"collection_test_load_index.json": "{\"name\":\"test_load\",\"codes\":[1,2,3,4],\"next_code\":5}",
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
	js := storage.NewJsonStorage(app.StorageDirectory())
	testFiles := map[string]string{
		"collection_test_index.json":      "{\"name\":\"test\",\"codes\":[1,2],\"next_code\":3}",
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
	app := cmd.NewApp()
	js := storage.NewJsonStorage(app.StorageDirectory())
	testFiles := map[string]string{
		"collection_test2_index.json":      "{\"name\":\"test2\",\"codes\":[1,2],\"next_code\":3}",
		"collection_test2_registry_1.json": "{\"code\":1,\"name\":\"John\"}",
		"collection_test2_registry_2.json": "{\"code\":2,\"name\":\"Jane\"}",
	}
	type Example struct {
		Code int    `json:"code"`
		Name string `json:"name"`
	}
	writeTestFiles(t, app, testFiles)
	collection, err := storage.NewCollection(js, "test2")
	if err != nil {
		t.Errorf("could not load collection: %s", err)
		t.FailNow()
	}

	t.Run("should list all registries", func(t *testing.T) {
		// Act
		raws, err := collection.List(0, 2)

		// Assert
		assert.Nil(t, err, "err should be nil")
		assert.Len(t, raws, 2, "raws should have 2 Raw entities")
		var decoded Example
		unmarshalErr1 := json.Unmarshal(raws[0], &decoded)
		assert.Nil(t, unmarshalErr1, "unmarshalErr1 should be nil")
		assert.Equal(t, "John", decoded.Name, "decoded.Name should be 'John'")
		unmarshalErr2 := json.Unmarshal(raws[1], &decoded)
		assert.Nil(t, unmarshalErr2, "unmarshalErr2 should be nil")
		assert.Equal(t, "Jane", decoded.Name, "decoded.Name should be 'Jane'")
	})

	t.Run("should list a single registry", func(t *testing.T) {
		// Act
		raws, err := collection.List(1, 100)

		// Assert
		assert.Nil(t, err, "err should be nil")
		assert.Len(t, raws, 1, "raws should have 1 Raw entity")
		var decoded Example
		unmarshalErr1 := json.Unmarshal(raws[0], &decoded)
		assert.Nil(t, unmarshalErr1, "unmarshalErr1 should be nil")
		assert.Equal(t, "Jane", decoded.Name, "decoded.Name should be 'Jane'")
	})

	t.Run("should return a blank list", func(t *testing.T) {
		// Act
		raws, err := collection.List(2, 100)

		// Assert
		assert.Nil(t, err, "err should be nil")
		assert.Len(t, raws, 0, "raws should be empty")
	})

	deleteTestFiles(t, app, testFiles)
}

func Test_Collection_DecodeRaw(t *testing.T) {
	app := cmd.NewApp()
	js := storage.NewJsonStorage(app.StorageDirectory())
	testFiles := map[string]string{
		"collection_test3_index.json":      "{\"name\":\"test3\",\"codes\":[1,2],\"next_code\":3}",
		"collection_test3_registry_1.json": "{\"code\":1,\"name\":\"John\"}",
		"collection_test3_registry_2.json": "{\"code\":2,\"name\":\"Jane\"}",
	}
	type Example struct {
		Code int    `json:"code"`
		Name string `json:"name"`
	}
	writeTestFiles(t, app, testFiles)
	collection, err := storage.NewCollection(js, "test3")
	if err != nil {
		t.Errorf("could not load collection: %s", err)
		t.FailNow()
	}

	t.Run("should decode entity", func(t *testing.T) {
		// Act
		raws, err := collection.List(1, 100)
		var decoded Example
		decodeErr := collection.DecodeRaw(raws[0], &decoded)

		// Assert
		assert.Nil(t, err, "err should be nil")
		assert.Len(t, raws, 1, "raws should have 1 Raw entity")
		assert.Nil(t, decodeErr, "decodeErr should be nil")
		assert.Equal(t, "Jane", decoded.Name, "decoded.Name should be 'Jane'")
	})

	deleteTestFiles(t, app, testFiles)
}

func Test_Collection_Save(t *testing.T) {
	app := cmd.NewApp()
	js := storage.NewJsonStorage(app.StorageDirectory())
	testFiles := map[string]string{
		"collection_test4_index.json":      "{\"name\":\"test4\",\"codes\":[1],\"next_code\":2}",
		"collection_test4_registry_1.json": "{\"code\":1,\"name\":\"John\"}",
		"collection_test4_registry_2.json": "",
	}
	type Example struct {
		Code int    `json:"code"`
		Name string `json:"name"`
	}
	writeTestFiles(t, app, testFiles)
	collection, err := storage.NewCollection(js, "test4")
	if err != nil {
		t.Errorf("could not load collection: %s", err)
		t.FailNow()
	}

	t.Run("should insert a new entity", func(t *testing.T) {
		// Arrange
		entity := Example{
			Code: collection.GetNextCode(),
			Name: "Jane",
		}

		// Act
		err := collection.Save(entity.Code, entity)

		// Assert
		assert.Nil(t, err, "err should be nil")
		var stored Example
		getErr := collection.Get(entity.Code, &stored)
		assert.Nil(t, getErr, "getErr should be nil")
		assert.Equal(t, entity.Code, stored.Code, "stored code should be equal to inserted code")
		assert.Equal(t, entity.Name, stored.Name, "stored name should be equal to inserted name")
	})

	t.Run("should update an entity", func(t *testing.T) {
		// Arrange
		var entity Example
		err := collection.Get(1, &entity)
		if err != nil {
			t.Errorf("could not get entity: %s\n", err)
			t.FailNow()
		}

		// Act
		entity.Name = "John Do"
		err = collection.Save(entity.Code, entity)

		// Assert
		assert.Nil(t, err, "err should be nil")
		var stored Example
		getErr := collection.Get(entity.Code, &stored)
		assert.Nil(t, getErr, "getErr should be nil")
		assert.Equal(t, entity.Code, stored.Code, "stored code should be equal to updated code")
		assert.Equal(t, "John Do", stored.Name, "stored name should be equal to updated name")
	})

	deleteTestFiles(t, app, testFiles)
}

func Test_Collection_Delete(t *testing.T) {
	app := cmd.NewApp()
	js := storage.NewJsonStorage(app.StorageDirectory())
	testFiles := map[string]string{
		"collection_test5_index.json":      "{\"name\":\"test5\",\"codes\":[1],\"next_code\":2}",
		"collection_test5_registry_1.json": "{\"code\":1,\"name\":\"John\"}",
	}
	type Example struct {
		Code int    `json:"code"`
		Name string `json:"name"`
	}
	writeTestFiles(t, app, testFiles)
	collection, err := storage.NewCollection(js, "test5")
	if err != nil {
		t.Errorf("could not load collection: %s", err)
		t.FailNow()
	}

	t.Run("should delete registry 1", func(t *testing.T) {
		// Act
		err := collection.Delete(1)

		// Assert
		assert.Nil(t, err, "err should be nil")
		raws, listErr := collection.List(0, 10)
		assert.Nil(t, listErr, "listErr should be nil")
		assert.Len(t, raws, 0, "raws should be empty")
	})

	t.Run("should delete registry 2 should result error", func(t *testing.T) {
		// Act
		err := collection.Delete(2)

		// Assert
		assert.Equal(t, errs.RegistryNotFoundErr, err, "err should be RegistryNotFoundErr")
	})

	delete(testFiles, "collection_test5_registry_1.json")
	deleteTestFiles(t, app, testFiles)
}
