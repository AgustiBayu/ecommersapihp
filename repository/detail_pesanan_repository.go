package repository

import (
	"EcommersAPIHP/model/domain"
	"context"
	"database/sql"
)

type DetailPesananRepository interface {
	Save(ctx context.Context, tx *sql.Tx, detailPesanan domain.DetailPesanan) domain.DetailPesanan
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.DetailPesanan, map[int]domain.Pesanan, map[int]domain.User, map[int]domain.Produk)
	FindById(ctx context.Context, tx *sql.Tx, detailPesananId int) (domain.DetailPesanan, domain.Pesanan, domain.User, domain.Produk, error)
	Update(ctx context.Context, tx *sql.Tx, detailPesanan domain.DetailPesanan) domain.DetailPesanan
	Delete(ctx context.Context, tx *sql.Tx, detailPesanan domain.DetailPesanan)
}
