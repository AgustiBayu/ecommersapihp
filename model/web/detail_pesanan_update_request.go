package web

type DetailPesananUpdateRequest struct {
	Id                   int `json:"id" validate:"required"`
	PesananId            int `json:"pesanan_id" validate:"required"`
	ProdukId             int `json:"produk_id" validate:"required"`
	JumlahProduk         int `json:"jumlah_produk" validate:"required"`
	HargaProdukPembelian int `json:"harga_produk_pembelian" validate:"required"`
}
