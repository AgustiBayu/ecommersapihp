package web

type KeranjangResponse struct {
	Id                int            `json:"id" validate:"required"`
	User              UserResponse   `json:"user"`
	Produk            ProdukResponse `json:"produk"`
	JumlahProduk      int            `json:"jumlah_produk" validate:"required"`
	TanggalPenambahan string         `json:"tanggal_penambahan" validate:"required"`
}
