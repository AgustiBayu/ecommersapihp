package domain

import "time"

type User struct {
	Id              int
	Name            string
	Email           string
	Password        string
	Pengguna        string
	TanggalBuatAkun time.Time
}
