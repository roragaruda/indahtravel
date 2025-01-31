package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	NamaPeserta   string `gorm:"not null" form:"nama-peserta"`
	NoTelp        string `gorm:"not null" form:"no-telp"`
	JmlPeserta    uint   `gorm:"not null" form:"jml-peserta"`
	FotoBukti     string
	KodePembelian string `gorm:"not null" form:"kode-pembelian"`
	Status        string `gorm:"not null default:pending" form:"status"`
	ProductID     uint   `gorm:"not null" form:"product-id"`
}
