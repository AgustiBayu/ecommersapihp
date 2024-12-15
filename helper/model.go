package helper

import (
	"EcommersAPIHP/model/domain"
	"EcommersAPIHP/model/web"
)

func ToProdukResponse(produk domain.Produk) web.ProdukResponse {
	return web.ProdukResponse{
		Id:           produk.Id,
		Name:         produk.Name,
		Deskripsi:    produk.Deskripsi,
		Harga:        produk.Harga,
		JumlahStok:   produk.JumlahStok,
		TanggalMasuk: FormatTanggal(produk.TanggalMasuk),
	}
}

func ToProdukResponses(produks []domain.Produk) []web.ProdukResponse {
	var produkRespon []web.ProdukResponse
	for _, produk := range produks {
		produkRespon = append(produkRespon, ToProdukResponse(produk))
	}
	return produkRespon
}
