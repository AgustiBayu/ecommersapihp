package repository

import (
	"EcommersAPIHP/model/domain"
	"context"
	"database/sql"
)

type UserRespository interface {
	Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	FindAll(ctx context.Context, tx *sql.Tx) []domain.User
	FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error)
	Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Delete(ctx context.Context, tx *sql.Tx, user domain.User)
}
