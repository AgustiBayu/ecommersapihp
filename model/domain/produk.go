package domain

import "time"

type Produk struct {
	Id           int
	Name         string
	Deskripsi    string
	Harga        int
	JumlahStok   int
	TanggalMasuk time.Time
}
