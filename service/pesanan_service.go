package service

import (
	"EcommersAPIHP/model/web"
	"context"
)

type PesananService interface {
	Create(ctx context.Context, request web.PesananCreateRequest) web.PesananResponse
	FindAll(ctx context.Context) []web.PesananResponse
	FindById(ctx context.Context, pesananId int) web.PesananResponse
	Update(ctx context.Context, request web.PesananUpdateRequest) web.PesananResponse
	Delete(ctx context.Context, pesananId int)
}
