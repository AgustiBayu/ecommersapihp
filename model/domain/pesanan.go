package domain

import "time"

type StatusPesanan string

type Pesanan struct {
	Id             int
	UserId         User
	TotalHarga     int
	Status         StatusPesanan
	TanggalPesanan time.Time
}
