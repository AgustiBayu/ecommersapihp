package domain

import "time"

type Keranjang struct {
	Id                int
	UserId            int
	ProdukId          int
	JumlahProduk      int
	TanggalPenambahan time.Time
}
