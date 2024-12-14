package service

import (
	"context"
)

type ProdukService interface {
	Create(ctx context.Context)
	FindAll(ctx context.Context)
	FindById(ctx context.Context)
	Update(ctx context.Context)
	Delete(ctx context.Context)
}
