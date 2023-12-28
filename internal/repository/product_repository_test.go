package repository_test

import (
	"gomarket/cmd"
	"gomarket/internal/entity"
	"gomarket/internal/errs"
	"gomarket/internal/repository"
	"gomarket/internal/usecases/dto"
	"gomarket/pkg/storage"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func Test_ProductRepository_Get(t *testing.T) {
	app := cmd.NewApp()
	js := storage.NewJsonStorage(app.StorageDirectory())
	db, err := storage.NewCollection(js, "test_product_repository")
	if err != nil {
		t.Errorf("could not initialize test dependencies: %s", err)
		t.FailNow()
	}

	testFiles := map[string]string{
		"collection_test_product_repository.json":            "",
		"collection_test_product_repository_registry_1.json": "",
		"collection_test_product_repository_registry_2.json": "",
		"collection_test_product_repository_registry_3.json": "",
	}

	r := repository.NewProductRepository(db)
	r.Insert(dto.ProductDTO{
		Name:              "Example 1",
		Materials:         make([]entity.Material, 0),
		SellingPriceCents: 100,
	})
	r.Insert(dto.ProductDTO{
		Name:              "Example 2",
		Materials:         make([]entity.Material, 0),
		SellingPriceCents: 100,
	})
	r.Insert(dto.ProductDTO{
		Name:              "Example 3",
		Materials:         make([]entity.Material, 0),
		SellingPriceCents: 100,
	})

	t.Run("should get example 2", func(t *testing.T) {
		// Act
		product, err := r.Get(2)

		// Assert
		assert.Nil(t, err, "err should be nil")
		assert.Equal(t, "Example 2", product.Name, "name should be 'Example 2'")
	})

	t.Run("should not get unexistent code", func(t *testing.T) {
		// Act
		product, err := r.Get(4)

		// Assert
		assert.Equal(t, errs.RegistryNotFoundErr, err, "err should be RegistryNotFoundErr")
		assert.Nil(t, product, "product should be nil")
	})

	deleteTestFiles(t, app, testFiles)
}

func Test_ProductRepository_List(t *testing.T) {
	app := cmd.NewApp()
	js := storage.NewJsonStorage(app.StorageDirectory())
	db, err := storage.NewCollection(js, "test_product_repository2")
	if err != nil {
		t.Errorf("could not initialize test dependencies: %s", err)
		t.FailNow()
	}

	testFiles := map[string]string{
		"collection_test_product_repository2.json":            "",
		"collection_test_product_repository2_registry_1.json": "",
		"collection_test_product_repository2_registry_2.json": "",
		"collection_test_product_repository2_registry_3.json": "",
	}

	r := repository.NewProductRepository(db)
	r.Insert(dto.ProductDTO{
		Name:              "Example 1",
		Materials:         make([]entity.Material, 0),
		SellingPriceCents: 100,
	})
	r.Insert(dto.ProductDTO{
		Name:              "Example 2",
		Materials:         make([]entity.Material, 0),
		SellingPriceCents: 100,
	})
	r.Insert(dto.ProductDTO{
		Name:              "Example 3",
		Materials:         make([]entity.Material, 0),
		SellingPriceCents: 100,
	})

	t.Run("should list 3 examples", func(t *testing.T) {
		// Act
		products, err := r.List(0, 100)

		// Assert
		assert.Nil(t, err, "err should be nil")
		assert.Len(t, products, 3, "products should have 3 items")
		assert.Equal(t, "Example 1", products[0].Name, "first product name should be 'Example 1'")
		assert.Equal(t, "Example 2", products[1].Name, "second product name should be 'Example 2'")
		assert.Equal(t, "Example 3", products[2].Name, "third product name should be 'Example 3'")
	})

	t.Run("should list blank on unexistent range", func(t *testing.T) {
		// Act
		products, err := r.List(5, 100)

		// Assert
		assert.Nil(t, err, "err should be nil")
		assert.Len(t, products, 0, "products should be blank")
	})

	deleteTestFiles(t, app, testFiles)
}

func Test_ProductRepository_Insert(t *testing.T) {
	app := cmd.NewApp()
	js := storage.NewJsonStorage(app.StorageDirectory())
	db, err := storage.NewCollection(js, "test_product_repository3")
	if err != nil {
		t.Errorf("could not initialize test dependencies: %s", err)
		t.FailNow()
	}

	testFiles := map[string]string{
		"collection_test_product_repository3.json":            "",
		"collection_test_product_repository3_registry_1.json": "",
	}

	r := repository.NewProductRepository(db)

	t.Run("should insert", func(t *testing.T) {
		// Arrange
		productsBeforeInsert, _ := r.List(0, 100)

		// Act
		_, err := r.Insert(dto.ProductDTO{
			Name:              "My Product",
			Materials:         make([]entity.Material, 0),
			SellingPriceCents: 0,
		})

		// Assert
		productsAfterInsert, _ := r.List(0, 100)
		assert.Nil(t, err, "err should be nil")
		assert.Len(t, productsBeforeInsert, 0, "products before insert should be blank")
		assert.Len(t, productsAfterInsert, 1, "products after insert should be filled with 1 item")
		assert.Equal(t, "My Product", productsAfterInsert[0].Name, "first product name should be 'My Product'")
		assert.Equal(t, 1, productsAfterInsert[0].Code, "first code should be 1")
		assert.Nil(t, productsAfterInsert[0].UpdatedAt, "updated_at should be nil")
	})

	deleteTestFiles(t, app, testFiles)
}

func Test_ProductRepository_Update(t *testing.T) {
	app := cmd.NewApp()
	js := storage.NewJsonStorage(app.StorageDirectory())
	db, err := storage.NewCollection(js, "test_product_repository4")
	if err != nil {
		t.Errorf("could not initialize test dependencies: %s", err)
		t.FailNow()
	}

	testFiles := map[string]string{
		"collection_test_product_repository4.json":            "",
		"collection_test_product_repository4_registry_1.json": "",
	}

	r := repository.NewProductRepository(db)
	r.Insert(dto.ProductDTO{
		Name:              "My Product",
		Materials:         make([]entity.Material, 0),
		SellingPriceCents: 0,
	})

	t.Run("should update", func(t *testing.T) {
		// Arrange
		productsBeforeUpdate, _ := r.List(0, 100)

		// Act
		_, err := r.Update(1, dto.ProductDTO{
			Name:              "My Product Changed",
			Materials:         make([]entity.Material, 0),
			SellingPriceCents: 0,
		})

		// Assert
		productsAfterUpdate, _ := r.List(0, 100)
		assert.Nil(t, err, "err should be nil")
		assert.Len(t, productsBeforeUpdate, len(productsAfterUpdate), "productsBeforeUpdate length should be equal to productsAfterUpdate length")
		assert.Equal(t, "My Product Changed", productsAfterUpdate[0].Name, "first product name should be 'My Product Changed'")
		assert.Equal(t, productsBeforeUpdate[0].Code, productsAfterUpdate[0].Code, "first code should not change")
		assert.NotNil(t, productsAfterUpdate[0].UpdatedAt, "updated_at should be filled")
	})

	t.Run("should not update unexistent product", func(t *testing.T) {
		// Act
		_, err := r.Update(6, dto.ProductDTO{
			Name:              "My Product Changed",
			Materials:         make([]entity.Material, 0),
			SellingPriceCents: 0,
		})

		// Assert
		assert.Equal(t, errs.RegistryNotFoundErr, err, "err should be RegistryNotFoundErr")
	})

	deleteTestFiles(t, app, testFiles)
}

func Test_ProductRepository_Delete(t *testing.T) {
	app := cmd.NewApp()
	js := storage.NewJsonStorage(app.StorageDirectory())
	db, err := storage.NewCollection(js, "test_product_repository5")
	if err != nil {
		t.Errorf("could not initialize test dependencies: %s", err)
		t.FailNow()
	}

	testFiles := map[string]string{
		"collection_test_product_repository5.json": "",
	}

	r := repository.NewProductRepository(db)
	r.Insert(dto.ProductDTO{
		Name:              "My Product",
		Materials:         make([]entity.Material, 0),
		SellingPriceCents: 0,
	})

	t.Run("should delete", func(t *testing.T) {
		// Arrange
		productsBeforeDelete, _ := r.List(0, 100)

		// Act
		err := r.Delete(1)

		// Assert
		productsAfterDelete, _ := r.List(0, 100)
		assert.Nil(t, err, "err should be nil")
		assert.Len(t, productsBeforeDelete, 1, "productsBeforeUpdate length should be 1")
		assert.Len(t, productsAfterDelete, 0, "productsAfterDelete length should be 0")
	})

	t.Run("should not delete unexistent product", func(t *testing.T) {
		// Act
		err := r.Delete(2)

		// Assert
		assert.Equal(t, errs.RegistryNotFoundErr, err, "err should be RegistryNotFoundErr")
	})

	deleteTestFiles(t, app, testFiles)
}
