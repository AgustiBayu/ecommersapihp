package service

import (
	"EcommersAPIHP/helper"
	"EcommersAPIHP/model/domain"
	"EcommersAPIHP/model/web"
	"EcommersAPIHP/repository"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type DetailPesananServiceImpl struct {
	DetailPesananRespository repository.DetailPesananRepository
	PesananRespository       repository.PesananRepository
	UserRepository           repository.UserRepository
	ProdukRepository         repository.ProdukRepository
	DB                       *sql.DB
	Validate                 *validator.Validate
}

func NewDetailPesananService(detailPesananRepository repository.DetailPesananRepository, pesananRepository repository.PesananRepository,
	userRepository repository.UserRepository, produkRepository repository.ProdukRepository, DB *sql.DB, validate *validator.Validate) DetailPesananService {
	return &DetailPesananServiceImpl{
		DetailPesananRespository: detailPesananRepository,
		PesananRespository:       pesananRepository,
		UserRepository:           userRepository,
		ProdukRepository:         produkRepository,
		DB:                       DB,
		Validate:                 validate,
	}
}

func (service *DetailPesananServiceImpl) Create(ctx context.Context, request web.DetailPesananCreateRequest) web.DetailPesananResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	detail := domain.DetailPesanan{
		PesananId:            request.PesananId,
		ProdukId:             request.ProdukId,
		JumlahProduk:         request.JumlahProduk,
		HargaProdukPembelian: request.HargaProdukPembelian,
	}
	detail = service.DetailPesananRespository.Save(ctx, tx, detail)
	pesanan, _, err := service.PesananRespository.FindById(ctx, tx, detail.PesananId)
	helper.PanicIfError(err)
	user, err := service.UserRepository.FindById(ctx, tx, pesanan.UserId)
	helper.PanicIfError(err)
	produk, err := service.ProdukRepository.FindById(ctx, tx, detail.ProdukId)
	helper.PanicIfError(err)
	return helper.ToDetailPesananResponse(detail, pesanan, user, produk)
}
func (service *DetailPesananServiceImpl) FindAll(ctx context.Context) []web.DetailPesananResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	detail, pesanan, user, produk := service.DetailPesananRespository.FindAll(ctx, tx)
	return helper.ToDetailPesananResponses(detail, pesanan, user, produk)
}
func (service *DetailPesananServiceImpl) FindById(ctx context.Context, detailPesananId int) web.DetailPesananResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	detail, pesanan, user, produk, err := service.DetailPesananRespository.FindById(ctx, tx, detailPesananId)
	helper.PanicIfError(err)
	return helper.ToDetailPesananResponse(detail, pesanan, user, produk)
}
func (service *DetailPesananServiceImpl) Update(ctx context.Context, request web.DetailPesananUpdateRequest) web.DetailPesananResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)
	detail, _, _, _, err := service.DetailPesananRespository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)
	pesanan, _, err := service.PesananRespository.FindById(ctx, tx, detail.PesananId)
	helper.PanicIfError(err)
	user, err := service.UserRepository.FindById(ctx, tx, pesanan.UserId)
	helper.PanicIfError(err)
	produk, err := service.ProdukRepository.FindById(ctx, tx, detail.ProdukId)
	helper.PanicIfError(err)

	detail.Id = request.Id
	detail.PesananId = request.PesananId
	detail.ProdukId = request.ProdukId
	detail.JumlahProduk = request.JumlahProduk
	detail.HargaProdukPembelian = request.HargaProdukPembelian
	detail = service.DetailPesananRespository.Update(ctx, tx, detail)

	return helper.ToDetailPesananResponse(detail, pesanan, user, produk)
}
func (service *DetailPesananServiceImpl) Delete(ctx context.Context, detailPesananId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	detail, _, _, _, err := service.DetailPesananRespository.FindById(ctx, tx, detailPesananId)
	helper.PanicIfError(err)
	service.DetailPesananRespository.Delete(ctx, tx, detail)
}
