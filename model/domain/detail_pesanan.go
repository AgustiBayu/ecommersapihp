package domain

type DetailPesanan struct {
	Id                   int
	PesananId            Pesanan
	ProdukId             Produk
	JumlahProduk         int
	HargaProdukPembelian int
}
