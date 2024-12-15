package repository

import (
	"EcommersAPIHP/helper"
	"EcommersAPIHP/model/domain"
	"context"
	"database/sql"
	"errors"
)

type ProduRepositoryImpl struct {
}

func NewProdukRepository() ProdukRepository {
	return &ProduRepositoryImpl{}
}

func (p *ProduRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, produk domain.Produk) domain.Produk {
	SQL := "INSERT INTO produks(name, deskripsi, harga, jumlah_stok, tanggal_masuk) VALUES(?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, produk.Name, produk.Deskripsi, produk.Harga, produk.JumlahStok, produk.TanggalMasuk)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	produk.Id = int(id)
	return produk
}
func (p *ProduRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Produk {
	SQL := "SELECT id, name, deskripsi, harga, jumlah_stok, tanggal_masuk FROM produks"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var produks []domain.Produk
	for rows.Next() {
		produk := domain.Produk{}
		err := rows.Scan(&produk.Id, &produk.Name, &produk.Deskripsi, &produk.Harga, &produk.JumlahStok, &produk.TanggalMasuk)
		helper.PanicIfError(err)
		produks = append(produks, produk)
	}
	return produks
}
func (p *ProduRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, produkId int) (domain.Produk, error) {
	SQL := "SELECT id, name, deskripsi, harga, jumlah_stok, tanggal_masuk FROM produks WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, produkId)
	helper.PanicIfError(err)
	defer rows.Close()

	produk := domain.Produk{}
	if rows.Next() {
		err := rows.Scan(&produk.Id, &produk.Name, &produk.Deskripsi, &produk.Harga, &produk.JumlahStok, &produk.TanggalMasuk)
		helper.PanicIfError(err)
		return produk, nil
	} else {
		return produk, errors.New("id produk is not found")
	}
}
func (p *ProduRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, produk domain.Produk) domain.Produk {
	SQL := "UPDATE produks SET name = ?, deskripsi = ?, harga = ?, jumlah_stok = ?, tanggal_masuk = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, produk.Name, produk.Deskripsi, produk.Harga, produk.JumlahStok, produk.TanggalMasuk, produk.Id)
	helper.PanicIfError(err)
	return produk
}
func (p *ProduRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, produk domain.Produk) {
	SQL := "DELETE FROM produks WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, produk.Id)
	helper.PanicIfError(err)
}
