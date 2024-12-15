package service

import (
	"EcommersAPIHP/helper"
	"EcommersAPIHP/model/domain"
	"EcommersAPIHP/model/web"
	"EcommersAPIHP/repository"
	"context"
	"database/sql"
	"time"

	"github.com/go-playground/validator/v10"
)

type ProdukServiceImpl struct {
	ProdukRepository repository.ProdukRepository
	DB               *sql.DB
	Validate         *validator.Validate
}

func NewProdukService(produkRepository repository.ProdukRepository,
	DB *sql.DB, validate *validator.Validate) ProdukService {
	return &ProdukServiceImpl{
		ProdukRepository: produkRepository,
		DB:               DB,
		Validate:         validate,
	}
}

func (service *ProdukServiceImpl) Create(ctx context.Context, request web.ProdukCreateRequest) web.ProdukResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	tanggalMasuk, err := time.Parse("02-01-2006", request.TanggalMasuk)
	helper.PanicIfError(err)
	produk := domain.Produk{
		Name:         request.Name,
		Deskripsi:    request.Deskripsi,
		Harga:        request.Harga,
		JumlahStok:   request.JumlahStok,
		TanggalMasuk: tanggalMasuk,
	}
	produk = service.ProdukRepository.Save(ctx, tx, produk)
	return helper.ToProdukResponse(produk)
}
func (service *ProdukServiceImpl) FindAll(ctx context.Context) []web.ProdukResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	produk := service.ProdukRepository.FindAll(ctx, tx)
	return helper.ToProdukResponses(produk)
}
func (service *ProdukServiceImpl) FindById(ctx context.Context, produkId int) web.ProdukResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	produk, err := service.ProdukRepository.FindById(ctx, tx, produkId)
	helper.PanicIfError(err)
	return helper.ToProdukResponse(produk)
}
func (service *ProdukServiceImpl) Update(ctx context.Context, request web.ProdukUpdateRequest) web.ProdukResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	tanggalMasuk, err := time.Parse("02-01-2006", request.TanggalMasuk)
	helper.PanicIfError(err)
	produk, err := service.ProdukRepository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	produk.Name = request.Name
	produk.Deskripsi = request.Deskripsi
	produk.Harga = request.Harga
	produk.JumlahStok = request.JumlahStok
	produk.TanggalMasuk = tanggalMasuk

	produk = service.ProdukRepository.Update(ctx, tx, produk)
	return helper.ToProdukResponse(produk)
}
func (service *ProdukServiceImpl) Delete(ctx context.Context, produkId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	produk, err := service.ProdukRepository.FindById(ctx, tx, produkId)
	helper.PanicIfError(err)
	service.ProdukRepository.Delete(ctx, tx, produk)
}
