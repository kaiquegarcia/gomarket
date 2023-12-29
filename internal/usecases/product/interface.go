package product

import "gomarket/internal/repository"

type Usecases interface {
	List()
	Get()
	Create()
	Update()
	Delete()
}

type usecases struct {
	repository repository.ProductRepository
}

func New(r repository.ProductRepository) Usecases {
	return &usecases{
		repository: r,
	}
}
