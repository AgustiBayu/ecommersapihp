package web

type DetailPesananResponse struct {
	Id                   int             `json:"id" validate:"required"`
	Pesanan              PesananResponse `json:"pesanan" validate:"required"`
	Produk               ProdukResponse  `json:"produk" validate:"required"`
	JumlahProduk         int             `json:"jumlah_produk" validate:"required"`
	HargaProdukPembelian int             `json:"harga_produk_pembelian" validate:"required"`
}
