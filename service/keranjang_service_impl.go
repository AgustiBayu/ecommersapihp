package service

import (
	"EcommersAPIHP/exception"
	"EcommersAPIHP/helper"
	"EcommersAPIHP/model/domain"
	"EcommersAPIHP/model/web"
	"EcommersAPIHP/repository"
	"context"
	"database/sql"
	"time"

	"github.com/go-playground/validator/v10"
)

type KeranjangServiceImpl struct {
	KeranjangRepository repository.KeranjangRepository
	UserRepository      repository.UserRepository
	ProdukRepository    repository.ProdukRepository
	DB                  *sql.DB
	Validate            *validator.Validate
}

func NewKeranjangService(keranjangRepository repository.KeranjangRepository,
	userRepository repository.UserRepository, produkRepository repository.ProdukRepository,
	DB *sql.DB, validate *validator.Validate) KeranjangService {
	return &KeranjangServiceImpl{
		KeranjangRepository: keranjangRepository,
		UserRepository:      userRepository,
		ProdukRepository:    produkRepository,
		DB:                  DB,
		Validate:            validate,
	}
}

func (service *KeranjangServiceImpl) Create(ctx context.Context, request web.KeranjangCreateRequest) web.KeranjangResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	TanggalPenambahan, err := time.Parse("02-01-2006", request.TanggalPenambahan)
	helper.PanicIfError(err)
	keranjang := domain.Keranjang{
		UserId:            request.UserId,
		ProdukId:          request.ProdukId,
		JumlahProduk:      request.JumlahProduk,
		TanggalPenambahan: TanggalPenambahan,
	}
	keranjang = service.KeranjangRepository.Save(ctx, tx, keranjang)
	user, err := service.UserRepository.FindById(ctx, tx, keranjang.UserId)
	helper.PanicIfError(err)
	produk, err := service.ProdukRepository.FindById(ctx, tx, keranjang.ProdukId)
	helper.PanicIfError(err)

	return helper.ToKeranjangResponse(keranjang, user, produk)
}
func (service *KeranjangServiceImpl) FindAll(ctx context.Context) []web.KeranjangResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	keranjang, user, produk := service.KeranjangRepository.FindAll(ctx, tx)
	return helper.ToKeranjangResponses(keranjang, user, produk)
}
func (service *KeranjangServiceImpl) FindById(ctx context.Context, keranjangId int) web.KeranjangResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)
	keranjang, user, produk, err := service.KeranjangRepository.FindById(ctx, tx, keranjangId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToKeranjangResponse(keranjang, user, produk)
}
func (service *KeranjangServiceImpl) Update(ctx context.Context, request web.KeranjangUpdateRequest) web.KeranjangResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	TanggalPenambahan, err := time.Parse("02-01-2006", request.TanggalPenambahan)
	helper.PanicIfError(err)
	keranjang, _, _, err := service.KeranjangRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	user, err := service.UserRepository.FindById(ctx, tx, keranjang.UserId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	produk, err := service.ProdukRepository.FindById(ctx, tx, keranjang.ProdukId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	err = helper.ValidateTanggalBaru(keranjang.TanggalPenambahan, TanggalPenambahan)
	helper.PanicIfError(err)

	keranjang.Id = request.Id
	keranjang.UserId = request.UserId
	keranjang.ProdukId = request.ProdukId
	keranjang.JumlahProduk = request.JumlahProduk
	keranjang.TanggalPenambahan = TanggalPenambahan

	keranjang = service.KeranjangRepository.Update(ctx, tx, keranjang)
	return helper.ToKeranjangResponse(keranjang, user, produk)
}
func (service *KeranjangServiceImpl) Delete(ctx context.Context, keranjangId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	keranjang, _, _, err := service.KeranjangRepository.FindById(ctx, tx, keranjangId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.KeranjangRepository.Delete(ctx, tx, keranjang)
}
