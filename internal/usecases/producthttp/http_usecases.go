package producthttp

import (
	"context"
	"gomarket/internal/entity"
	"gomarket/internal/repository"
)

type HTTP interface {
	// Create will try to register a Product on the storage, returning its entity if successful or an error if not
	Create(ctx context.Context, input CreateInput) (*entity.Product, error)
	// Update will try to update a Product from the storage, returning its entity if successful or an error if not
	Update(ctx context.Context, input UpdateInput) (*entity.Product, error)
	// List will try to return a list of Products from the storage, returning the complete list or an error if it's not possible
	List(ctx context.Context, input ListInput) ([]ListOutput, error)
	// Get will try to retrieve a more detailed information about a Product from the storage based on its code, returning the detailed response if succesful or an error if not
	Get(ctx context.Context, input GetInput) (*GetOutput, error)
	// Delete will try to remove a Product from the storage, returning nil if successful or an error if not
	Delete(ctx context.Context, input DeleteInput) error
}

type httpUsecases struct {
	repository repository.ProductRepository
}

// NewHTTP returns an implemention of HTTP interface
func NewHTTP(r repository.ProductRepository) HTTP {
	return &httpUsecases{
		repository: r,
	}
}
