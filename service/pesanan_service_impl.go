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

type PesananServiceImpl struct {
	PesananRepository repository.PesananRepository
	UserRepository    repository.UserRepository
	DB                *sql.DB
	Validate          *validator.Validate
}

func NewPesananService(pesananRepository repository.PesananRepository,
	userRepository repository.UserRepository, DB *sql.DB,
	validate *validator.Validate) PesananService {
	return &PesananServiceImpl{
		PesananRepository: pesananRepository,
		UserRepository:    userRepository,
		DB:                DB,
		Validate:          validate,
	}
}

func (service *PesananServiceImpl) Create(ctx context.Context, request web.PesananCreateRequest) web.PesananResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	TanggalPesanan, err := time.Parse("02-01-2006", request.TanggalPesanan)
	helper.PanicIfError(err)
	pesanan := domain.Pesanan{
		UserId:         request.UserId,
		TotalHarga:     request.TotalHarga,
		Status:         domain.StatusPesanan(request.Status),
		TanggalPesanan: TanggalPesanan,
	}
	pesanan = service.PesananRepository.Save(ctx, tx, pesanan)
	user, err := service.UserRepository.FindById(ctx, tx, pesanan.UserId)
	helper.PanicIfError(err)

	return helper.ToPesananResponse(pesanan, user)
}
func (service *PesananServiceImpl) FindAll(ctx context.Context) []web.PesananResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)
	pesanan, user := service.PesananRepository.FindAll(ctx, tx)
	return helper.ToPesananResponses(pesanan, user)
}
func (service *PesananServiceImpl) FindById(ctx context.Context, pesananId int) web.PesananResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	pesanan, user, err := service.PesananRepository.FindById(ctx, tx, pesananId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToPesananResponse(pesanan, user)
}
func (service *PesananServiceImpl) Update(ctx context.Context, request web.PesananUpdateRequest) web.PesananResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	TanggalPesanan, err := time.Parse("02-01-2006", request.TanggalPesanan)
	helper.PanicIfError(err)
	pesanan, _, err := service.PesananRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	user, err := service.UserRepository.FindById(ctx, tx, pesanan.UserId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	err = helper.ValidateTanggalBaru(pesanan.TanggalPesanan, TanggalPesanan)
	helper.PanicIfError(err)

	pesanan.UserId = request.UserId
	pesanan.TotalHarga = request.TotalHarga
	pesanan.Status = domain.StatusPesanan(request.Status)
	pesanan.TanggalPesanan = TanggalPesanan

	pesanan = service.PesananRepository.Update(ctx, tx, pesanan)
	return helper.ToPesananResponse(pesanan, user)
}
func (service *PesananServiceImpl) Delete(ctx context.Context, pesananId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	pesanan, _, err := service.PesananRepository.FindById(ctx, tx, pesananId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.PesananRepository.Delete(ctx, tx, pesanan)
}
