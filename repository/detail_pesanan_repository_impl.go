package repository

import (
	"EcommersAPIHP/helper"
	"EcommersAPIHP/model/domain"
	"context"
	"database/sql"
	"errors"
)

type DetailPesananRepositoryImpl struct{}

func NewDetailPesananRepository() DetailPesananRepository {
	return &DetailPesananRepositoryImpl{}
}

func (d *DetailPesananRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, detailPesanan domain.DetailPesanan) domain.DetailPesanan {
	SQL := "INSERT INTO detail_pesanans(pesanan_id, produk_id, jumlah_produk,harga_produk_pembelian) VALUES(?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, detailPesanan.PesananId, detailPesanan.ProdukId, detailPesanan.JumlahProduk, detailPesanan.HargaProdukPembelian)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	detailPesanan.Id = int(id)
	return detailPesanan
}
func (d *DetailPesananRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.DetailPesanan, map[int]domain.Pesanan, map[int]domain.User, map[int]domain.Produk) {
	SQL := "select d.id, d.pesanan_id, d.produk_id, p.id, p.user_id,u.id,u.name,u.email,u.pengguna,u.tanggal_buat_akun, p.total_harga, p.status, p.tanggal_pesanan,p2.id, p2.name, p2.deskripsi ,p2.harga , p2.jumlah_stok , p2.tanggal_masuk, d.jumlah_produk, d.harga_produk_pembelian from detail_pesanans d inner join pesanans p on d.pesanan_id = p.id inner join produks p2 on d.produk_id = p2.id inner join users u on p.user_id = u.id"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var details []domain.DetailPesanan
	pesananMap := make(map[int]domain.Pesanan)
	userMap := make(map[int]domain.User)
	produkMap := make(map[int]domain.Produk)

	for rows.Next() {
		detail := domain.DetailPesanan{}
		pesanan := domain.Pesanan{}
		user := domain.User{}
		produk := domain.Produk{}
		err := rows.Scan(&detail.Id, &detail.PesananId, &detail.ProdukId, &pesanan.Id, &pesanan.UserId, &user.Id, &user.Name, &user.Email, &user.Pengguna, &user.TanggalBuatAkun, &pesanan.TotalHarga,
			&pesanan.Status, &pesanan.TanggalPesanan, &produk.Id, &produk.Name, &produk.Deskripsi, &produk.Harga, &produk.JumlahStok,
			&produk.TanggalMasuk, &detail.JumlahProduk, &detail.HargaProdukPembelian)
		helper.PanicIfError(err)
		details = append(details, detail)
		pesananMap[pesanan.Id] = pesanan
		userMap[user.Id] = user
		produkMap[produk.Id] = produk
	}
	return details, pesananMap, userMap, produkMap
}
func (d *DetailPesananRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, detailPesananId int) (domain.DetailPesanan, domain.Pesanan, domain.User, domain.Produk, error) {
	SQL := "select d.id, d.pesanan_id, d.produk_id, p.id, p.user_id,u.id,u.name,u.email,u.pengguna,u.tanggal_buat_akun, p.total_harga, p.status, p.tanggal_pesanan,p2.id, p2.name, p2.deskripsi ,p2.harga , p2.jumlah_stok , p2.tanggal_masuk, d.jumlah_produk, d.harga_produk_pembelian from detail_pesanans d inner join pesanans p on d.pesanan_id = p.id inner join produks p2 on d.produk_id = p2.id inner join users u on p.user_id = u.id WHERE d.id = ?"
	rows, err := tx.QueryContext(ctx, SQL, detailPesananId)
	helper.PanicIfError(err)
	defer rows.Close()

	detail := domain.DetailPesanan{}
	pesanan := domain.Pesanan{}
	produk := domain.Produk{}
	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&detail.Id, &detail.PesananId, &detail.ProdukId, &pesanan.Id, &pesanan.UserId, &user.Id, &user.Name, &user.Email, &user.Pengguna, &user.TanggalBuatAkun, &pesanan.TotalHarga,
			&pesanan.Status, &pesanan.TanggalPesanan, &produk.Id, &produk.Name, &produk.Deskripsi, &produk.Harga, &produk.JumlahStok,
			&produk.TanggalMasuk, &detail.JumlahProduk, &detail.HargaProdukPembelian)
		helper.PanicIfError(err)
		return detail, pesanan, user, produk, nil
	} else {
		return detail, pesanan, user, produk, errors.New("detail pesanan id not found")
	}
}
func (d *DetailPesananRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, detailPesanan domain.DetailPesanan) domain.DetailPesanan {
	SQL := "UPDATE detail_pesanans SET pesanan_id=?, produk_id=?, jumlah_produk=?,harga_produk_pembelian=? WHERE id =?"
	_, err := tx.ExecContext(ctx, SQL, detailPesanan.PesananId, detailPesanan.ProdukId, detailPesanan.JumlahProduk, detailPesanan.HargaProdukPembelian, detailPesanan.Id)
	helper.PanicIfError(err)
	return detailPesanan
}
func (d *DetailPesananRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, detailPesanan domain.DetailPesanan) {
	SQL := "DELETE FROM detail_pesanan WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, detailPesanan)
	helper.PanicIfError(err)
}
