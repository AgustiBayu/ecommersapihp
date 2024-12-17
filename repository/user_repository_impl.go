package repository

import (
	"EcommersAPIHP/helper"
	"EcommersAPIHP/model/domain"
	"context"
	"database/sql"
	"errors"
)

type UserRepositoryImpl struct{}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (u *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "INSERT INTO users(name,email,password,pengguna,tanggal_buat_akun) VALUES(?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL, user.Name, user.Email, user.Password, user.Pengguna, user.TanggalBuatAkun)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	user.Id = int(id)
	return user
}

func (u *UserRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error) {
	SQL := "SELECT id,name,email,password,pengguna,tanggal_buat_akun FROM users WHERE email = ?"
	rows, err := tx.QueryContext(ctx, SQL, email)
	helper.PanicIfError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Pengguna, &user.TanggalBuatAkun)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("email not found")
	}
}
func (u *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error) {
	SQL := "SELECT id,name,email,password,pengguna,tanggal_buat_akun FROM users WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Pengguna, &user.TanggalBuatAkun)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("user id not found")
	}
}
