package mocking

import (
	"inventaris/models"

	"github.com/stretchr/testify/mock"
)

type ProdukRepositoryMock struct{
	mock.Mock
}

func(m *ProdukRepositoryMock) Create(produk models.Produk) (models.Produk, error){
	args := m.Called(produk)
	return args.Get(0).(models.Produk), args.Error(1)
}

func(m *ProdukRepositoryMock) Update(produk models.Produk) (models.Produk, error){
	
	return models.Produk{}, nil
}

func(m *ProdukRepositoryMock) FindById(produkId int) (models.Produk, error){
	args := m.Called(produkId)
	return args.Get(0).(models.Produk), args.Error(1)
}	

func(m *ProdukRepositoryMock) Delete(produkId int) error{
	args := m.Called(produkId)
	return  args.Error(1)
}

func(m *ProdukRepositoryMock) FindAll() ([]models.Produk, error){
	
	return nil, nil
}	

func(m *ProdukRepositoryMock) UpdateImage(produkId int, gambar string) (models.Produk, error){
	args := m.Called(produkId)
	return args.Get(0).(models.Produk), args.Error(1)
}	