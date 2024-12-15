package web

type PesananCreateRequest struct {
	UserId         int    `json:"user_id" validate:"required"`
	TotalHarga     int    `json:"total_harga" validate:"required"`
	Status         string `json:"status" validate:"required,oneof=PENDING COMPLETED CANCELLED"`
	TanggalPesanan string `json:"tanggal_pesanan" validate:"required"`
}
