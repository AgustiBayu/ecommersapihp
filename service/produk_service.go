package service

import (
	"EcommersAPIHP/model/web"
	"context"
)

type ProdukService interface {
	Create(ctx context.Context, request web.ProdukCreateRequest) web.ProdukResponse
	FindAll(ctx context.Context) []web.ProdukResponse
	FindById(ctx context.Context, produkId int) web.ProdukResponse
	Update(ctx context.Context, request web.ProdukUpdateRequest) web.ProdukResponse
	Delete(ctx context.Context, produkId int)
}
