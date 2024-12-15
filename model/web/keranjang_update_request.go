package web

type KeranjangUpdateRequest struct {
	Id                int    `json:"id" validate:"required"`
	UserId            int    `json:"user_id" validate:"required"`
	ProdukId          int    `json:"produk_id" validate:"required"`
	JumlahProduk      int    `json:"jumlah_produk" validate:"required"`
	TanggalPenambahan string `json:"tanggal_penambahan" validate:"required"`
}
