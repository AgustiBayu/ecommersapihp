package repository

import (
	"EcommersAPIHP/model/domain"
	"context"
	"database/sql"
)

type DetailPesananRepository interface {
	Save(ctx context.Context, tx *sql.Tx, detailPesanan domain.DetailPesanan) domain.DetailPesanan
	FindAll(ctx context.Context, tx *sql.Tx) []domain.DetailPesanan
	FindById(ctx context.Context, tx *sql.Tx, detailPesananId int) (domain.DetailPesanan, error)
	Update(ctx context.Context, tx *sql.Tx, detailPesanan domain.DetailPesanan) domain.DetailPesanan
	Delete(ctx context.Context, tx *sql.Tx, detailPesanan domain.DetailPesanan)
}
