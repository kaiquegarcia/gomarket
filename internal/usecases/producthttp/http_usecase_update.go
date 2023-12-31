package producthttp

import (
	"context"
	"gomarket/internal/entity"
)

type UpdateInput struct {
	CreateInput
	Code int `validate:"required,min=1"`
}

func (u *httpUsecases) Update(ctx context.Context, input UpdateInput) (*entity.Product, error) {
	dto, err := newProductDTO(input.CreateInput)
	if err != nil {
		return nil, err
	}

	return u.repository.Update(input.Code, *dto)
}
