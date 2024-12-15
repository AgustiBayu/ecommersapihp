package domain

import "time"

type StatusPesanan string

const (
	StatusPending   StatusPesanan = "PENDING"
	StatusCompleted StatusPesanan = "COMPLETED"
	StatusCancelled StatusPesanan = "CANCELLED"
)

type Pesanan struct {
	Id             int
	UserId         int
	TotalHarga     int
	Status         StatusPesanan
	TanggalPesanan time.Time
}
