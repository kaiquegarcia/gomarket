package productcli

import "gomarket/internal/repository"

type CLI interface {
	List()
	Get()
	Create()
	Update()
	Delete()
}

type cliUsecases struct {
	repository repository.ProductRepository
}

// NewCLI returns an implemention of CLI interface
func NewCLI(r repository.ProductRepository) CLI {
	return &cliUsecases{
		repository: r,
	}
}
