package service

import (
	"EcommersAPIHP/model/web"
	"context"
)

type KeranjangService interface {
	Create(ctx context.Context, request web.KeranjangCreateRequest) web.KeranjangResponse
	FindAll(ctx context.Context) []web.KeranjangResponse
	FindById(ctx context.Context, keranjangId int) web.KeranjangResponse
	Update(ctx context.Context, request web.KeranjangUpdateRequest) web.KeranjangResponse
	Delete(ctx context.Context, keranjangId int)
}
