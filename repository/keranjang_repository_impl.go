package repository

import (
	"EcommersAPIHP/helper"
	"EcommersAPIHP/model/domain"
	"context"
	"database/sql"
	"errors"
)

type KeranjangRepositoryImpl struct{}

func NewKeranjangRepository() KeranjangRepository {
	return &KeranjangRepositoryImpl{}
}

func (repo *KeranjangRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, keranjang domain.Keranjang) domain.Keranjang {
	SQL := "INSERT INTO keranjangs (user_id,produk_id,jumlah_produk,tanggal_penambahan) VALUES(?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, keranjang.UserId, keranjang.ProdukId, keranjang.JumlahProduk, keranjang.TanggalPenambahan)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	keranjang.Id = int(id)
	return keranjang
}
func (repo *KeranjangRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Keranjang, map[int]domain.User, map[int]domain.Produk) {
	SQL := "select k.id, k.user_id, k.produk_id, u.id, u.name, u.email, u.pengguna, u.tanggal_buat_akun, p.id, p.name, p.deskripsi , p.harga , p.jumlah_stok , p.tanggal_masuk ,k.jumlah_produk, k.tanggal_penambahan from keranjangs k inner join users u on k.user_id = u.id inner join produks p on k.produk_id = p.id"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var keranjangs []domain.Keranjang
	userMap := make(map[int]domain.User)
	produkMap := make(map[int]domain.Produk)
	for rows.Next() {
		keranjang := domain.Keranjang{}
		user := domain.User{}
		produk := domain.Produk{}
		err := rows.Scan(&keranjang.Id, &keranjang.UserId, &keranjang.ProdukId,
			&user.Id, &user.Name, &user.Email, &user.Pengguna, &user.TanggalBuatAkun,
			&produk.Id, &produk.Name, &produk.Deskripsi, &produk.Harga, &produk.JumlahStok,
			&produk.TanggalMasuk, &keranjang.JumlahProduk, &keranjang.TanggalPenambahan)
		helper.PanicIfError(err)
		keranjangs = append(keranjangs, keranjang)

		userMap[user.Id] = user
		produkMap[produk.Id] = produk
	}
	return keranjangs, userMap, produkMap
}
func (repo *KeranjangRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, keranjangId int) (domain.Keranjang, domain.User, domain.Produk, error) {
	SQL := "select k.id, k.user_id, k.produk_id, u.id, u.name, u.email, u.pengguna, u.tanggal_buat_akun, p.id, p.name, p.deskripsi , p.harga , p.jumlah_stok , p.tanggal_masuk ,k.jumlah_produk, k.tanggal_penambahan from keranjangs k inner join users u on k.user_id = u.id inner join produks p on k.produk_id = p.id where k.id = ?"
	rows, err := tx.QueryContext(ctx, SQL, keranjangId)
	helper.PanicIfError(err)
	defer rows.Close()

	keranjang := domain.Keranjang{}
	user := domain.User{}
	produk := domain.Produk{}
	if rows.Next() {
		err := rows.Scan(&keranjang.Id, &keranjang.UserId, &keranjang.ProdukId,
			&user.Id, &user.Name, &user.Email, &user.Pengguna, &user.TanggalBuatAkun,
			&produk.Id, &produk.Name, &produk.Deskripsi, &produk.Harga, &produk.JumlahStok,
			&produk.TanggalMasuk, &keranjang.JumlahProduk, &keranjang.TanggalPenambahan)
		helper.PanicIfError(err)
		return keranjang, user, produk, nil
	} else {
		return keranjang, user, produk, errors.New("keranjang id not found")
	}
}
func (repo *KeranjangRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, keranjang domain.Keranjang) domain.Keranjang {
	SQL := "UPDATE keranjangs SET user_id = ?,produk_id=?,jumlah_produk=?,tanggal_penambahan=? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, keranjang.UserId, keranjang.ProdukId, keranjang.JumlahProduk, keranjang.TanggalPenambahan, keranjang.Id)
	helper.PanicIfError(err)
	return keranjang
}
func (repo *KeranjangRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, keranjang domain.Keranjang) {
	SQL := "DELETE FROM keranjangs WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, keranjang.Id)
	helper.PanicIfError(err)
}
