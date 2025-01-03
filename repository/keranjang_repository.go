package repository

import (
	"EcommersAPIHP/model/domain"
	"context"
	"database/sql"
)

type KeranjangRepository interface {
	Save(ctx context.Context, tx *sql.Tx, keranjang domain.Keranjang) domain.Keranjang
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Keranjang, map[int]domain.User, map[int]domain.Produk)
	FindById(ctx context.Context, tx *sql.Tx, keranjangId int) (domain.Keranjang, domain.User, domain.Produk, error)
	Update(ctx context.Context, tx *sql.Tx, keranjang domain.Keranjang) domain.Keranjang
	Delete(ctx context.Context, tx *sql.Tx, keranjang domain.Keranjang)
}
