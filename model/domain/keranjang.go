package domain

import "time"

type Keranjang struct {
	Id                int
	UserId            User
	ProdukId          Produk
	JumlahProduk      int
	TanggalPenambahan time.Time
}
