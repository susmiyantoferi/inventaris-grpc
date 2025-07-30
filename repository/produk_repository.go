package repository

import "inventaris/models"

type ProdukRepository interface {
	Create(produk models.Produk) models.Produk
	Update(produk models.Produk) (models.Produk, error)
	Delete(produkId int) error
	FindById(produkId int) (models.Produk, error)
	FindAll() ([]models.Produk, error)
	UpdateImage(produkId int, gambar string) (models.Produk, error)
}
