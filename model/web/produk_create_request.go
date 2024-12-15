package web

type ProdukCreateRequest struct {
	Name         string `json:"name" validate:"required"`
	Deskripsi    string `json:"deskripsi" validate:"required"`
	Harga        int    `json:"harga" validate:"required"`
	JumlahStok   int    `json:"jumlah_stok" validate:"required"`
	TanggalMasuk string `json:"tanggal_masuk" validate:"required"`
}
