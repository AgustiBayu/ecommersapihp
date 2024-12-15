package domain

import "time"

type StatusPengguna string

const (
	Admin    StatusPengguna = "ADMIN"
	Pengguna StatusPengguna = "PENGGUNA"
)

type User struct {
	Id              int
	Name            string
	Email           string
	Password        string
	Pengguna        StatusPengguna
	TanggalBuatAkun time.Time
}
