package repository

import (
	"gomarket/internal/entity"
	"gomarket/internal/usecases/dto"
	"gomarket/pkg/storage"
	"time"
)

type ProductRepository interface {
	// Get will try to retrieve a product from the storage
	Get(code int) (*entity.Product, error)
	// Count will retrieve the quantity of products registered on the storage
	Count() int
	// List will try to retrieve a list of products from the storage based on the offset and limit defined on arguments
	List(offset int, limit int) ([]*entity.Product, error)
	// Insert will try to add a product on the storage
	Insert(dto dto.ProductDTO) (*entity.Product, error)
	// Update will try to change the values of a product on the storage
	Update(code int, dto dto.ProductDTO) (*entity.Product, error)
	// Delete will try to remove a product from the storage
	Delete(code int) error
}

type productRepository struct {
	db storage.Collection
}

func NewProductRepository(db storage.Collection) ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) Get(code int) (*entity.Product, error) {
	var dest entity.Product
	err := r.db.Get(code, &dest)
	if err != nil {
		return nil, err
	}

	return &dest, nil
}

func (r *productRepository) Count() int {
	return r.db.Count()
}

func (r *productRepository) List(offset int, limit int) ([]*entity.Product, error) {
	raws, err := r.db.List(offset, limit)
	if err != nil {
		return nil, err
	}

	list := make([]*entity.Product, len(raws))
	for index, raw := range raws {
		var product entity.Product
		err := r.db.DecodeRaw(raw, &product)
		if err != nil {
			return nil, err
		}

		list[index] = &product
	}

	return list, nil
}

func (r *productRepository) Insert(dto dto.ProductDTO) (*entity.Product, error) {
	product := dto.ToEntity(
		r.db.GetNextCode(),
		time.Now(),
		nil,
	)

	err := r.db.Save(product.Code, product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *productRepository) Update(code int, dto dto.ProductDTO) (*entity.Product, error) {
	var product entity.Product
	err := r.db.Get(code, &product)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	product = dto.ToEntity(
		code,
		product.CreatedAt,
		&now,
	)

	err = r.db.Save(code, product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (r *productRepository) Delete(code int) error {
	return r.db.Delete(code)
}
