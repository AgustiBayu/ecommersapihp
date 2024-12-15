package web

type PesananResponse struct {
	Id             int          `json:"id" validate:"required"`
	User           UserResponse `json:"user"`
	TotalHarga     int          `json:"total_harga" validate:"required"`
	Status         string       `json:"status" validate:"required,oneof=PENDING COMPLETED CANCELLED"`
	TanggalPesanan string       `json:"tanggal_pesanan" validate:"required"`
}
