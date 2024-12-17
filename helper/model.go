package helper

import (
	"EcommersAPIHP/model/domain"
	"EcommersAPIHP/model/web"
)

func ToPesananResponses(pesanans []domain.Pesanan, users map[int]domain.User) []web.PesananResponse {
	var pesananResponses []web.PesananResponse
	for _, pesanan := range pesanans {
		user, exists := users[pesanan.UserId]
		if !exists {
			user = domain.User{} // Gunakan user kosong jika tidak ditemukan
		}
		pesananResponses = append(pesananResponses, ToPesananResponse(pesanan, user))
	}
	return pesananResponses
}

func ToPesananResponse(pesanan domain.Pesanan, user domain.User) web.PesananResponse {
	return web.PesananResponse{
		Id: pesanan.Id,
		User: web.UserResponse{
			Id:              user.Id,
			Name:            user.Name,
			Email:           user.Email,
			Pengguna:        string(user.Pengguna),
			TanggalBuatAkun: FormatTanggal(user.TanggalBuatAkun),
		},
		TotalHarga:     pesanan.TotalHarga,
		Status:         string(pesanan.Status),
		TanggalPesanan: FormatTanggal(pesanan.TanggalPesanan),
	}
}
func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		Id:              user.Id,
		Name:            user.Name,
		Email:           user.Email,
		Pengguna:        string(user.Pengguna),
		TanggalBuatAkun: FormatTanggal(user.TanggalBuatAkun),
	}
}

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
