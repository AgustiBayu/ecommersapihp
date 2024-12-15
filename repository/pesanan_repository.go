package repository

import (
	"EcommersAPIHP/model/domain"
	"context"
	"database/sql"
)

type PesananRepository interface {
	Save(ctx context.Context, tx *sql.Tx, pesanan domain.Pesanan) domain.Pesanan
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Pesanan
	FindById(ctx context.Context, tx *sql.Tx, pesananId int) (domain.Pesanan, error)
	Update(ctx context.Context, tx *sql.Tx, pesanan domain.Pesanan) domain.Pesanan
	Delete(ctx context.Context, tx *sql.Tx, pesanan domain.Pesanan)
}
