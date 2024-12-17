package repository

import (
	"EcommersAPIHP/helper"
	"EcommersAPIHP/model/domain"
	"context"
	"database/sql"
	"errors"
)

type PesananRepositoryImpl struct{}

func NewPesananRepository() PesananRepository {
	return &PesananRepositoryImpl{}
}

func (p *PesananRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, pesanan domain.Pesanan) domain.Pesanan {
	SQL := "INSERT INTO pesanans(user_id,total_harga,status,tanggal_pesanan) VALUES(?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, pesanan.UserId, pesanan.TotalHarga, pesanan.Status, pesanan.TanggalPesanan)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	pesanan.Id = int(id)
	return pesanan
}
func (p *PesananRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Pesanan, map[int]domain.User) {
	SQL := "select p.id,p.user_id,u.id ,u.name,u.email,u.pengguna,u.tanggal_buat_akun,p.total_harga,p.status,p.tanggal_pesanan from pesanans p inner join users u on p.user_id = u.id"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var pesanans []domain.Pesanan
	usersMap := make(map[int]domain.User)

	for rows.Next() {
		pesanan := domain.Pesanan{}
		user := domain.User{}
		err := rows.Scan(&pesanan.Id, &pesanan.UserId, &user.Id, &user.Name, &user.Email, &user.Pengguna, &user.TanggalBuatAkun, &pesanan.TotalHarga, &pesanan.Status, &pesanan.TanggalPesanan)
		helper.PanicIfError(err)
		pesanans = append(pesanans, pesanan)
		usersMap[user.Id] = user
	}
	return pesanans, usersMap
}
func (p *PesananRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, pesananId int) (domain.Pesanan, domain.User, error) {
	SQL := "select p.id,p.user_id,u.id ,u.name,u.email,u.pengguna,u.tanggal_buat_akun,p.total_harga,p.status,p.tanggal_pesanan from pesanans p inner join users u on p.user_id = u.id WHERE p.id = ?"
	rows, err := tx.QueryContext(ctx, SQL, pesananId)
	helper.PanicIfError(err)
	defer rows.Close()

	pesanan := domain.Pesanan{}
	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&pesanan.Id, &pesanan.UserId, &user.Id, &user.Name, &user.Email, &user.Pengguna, &user.TanggalBuatAkun, &pesanan.TotalHarga, &pesanan.Status, &pesanan.TanggalPesanan)
		helper.PanicIfError(err)
		return pesanan, user, nil
	}
	return pesanan, user, errors.New("pesanan id not found")
}
func (p *PesananRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, pesanan domain.Pesanan) domain.Pesanan {
	SQL := "UPDATE pesanans SET user_id = ?, total_harga = ?, status = ?, tanggal_pesanan = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, pesanan.UserId, pesanan.TotalHarga, pesanan.Status, pesanan.TanggalPesanan, pesanan.Id)
	helper.PanicIfError(err)
	return pesanan
}
func (p *PesananRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, pesanan domain.Pesanan) {
	SQL := "DELETE FROM pesanans WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, pesanan.Id)
	helper.PanicIfError(err)
}
