package producthttp

import "context"

type DeleteInput struct {
	GetInput
}

func (u *httpUsecases) Delete(ctx context.Context, input DeleteInput) error {
	return u.repository.Delete(input.Code)
}
