package service

import (
	"EcommersAPIHP/model/web"
	"context"
)

type DetailPesananService interface {
	Create(ctx context.Context, request web.DetailPesananCreateRequest) web.DetailPesananResponse
	FindAll(ctx context.Context) []web.DetailPesananResponse
	FindById(ctx context.Context, detailPesananId int) web.DetailPesananResponse
	Update(ctx context.Context, request web.DetailPesananUpdateRequest) web.DetailPesananResponse
	Delete(ctx context.Context, detailPesananId int)
}
