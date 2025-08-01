package test

import (
	"errors"
	"inventaris/mocking"
	"inventaris/models"
	"inventaris/repository"
	"inventaris/service"
	"inventaris/web"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateSuccess(t *testing.T) {
	produkMock := new(mocking.ProdukRepositoryMock)
	loggingMock := new(mocking.LoggingMock)
	validate := validator.New()

	svc := service.NewProdukServiceImpl(produkMock, validate, loggingMock)

	req := web.CreateProdukRequest{
		Nama:      "chiki",
		Deskripsi: "makanan",
		Harga:     "10000",
		Kategori:  "snack",
	}

	harga, _ := decimal.NewFromString(req.Harga)

	expected := models.Produk{
		Id:        1,
		Nama:      "chiki",
		Deskripsi: "makanan",
		Harga:     harga,
		Kategori:  "snack",
	}

	produkMock.On("Create", mock.MatchedBy(func(p models.Produk) bool {
		return p.Nama == req.Nama && p.Harga.Equal(harga)
	})).Return(expected, nil)

	produkMock.On("Create", mock.Anything).Return(expected, nil)

	resp, err := svc.Create(req)

	assert.Nil(t, err)
	assert.Equal(t, expected.Nama, resp.Nama)
	assert.Equal(t, expected.Deskripsi, resp.Deskripsi)
	assert.Equal(t, expected.Harga, resp.Harga)
	assert.Equal(t, expected.Kategori, resp.Kategori)
	assert.Equal(t, expected.Gambar, resp.Gambar)

	produkMock.AssertExpectations(t)
}

func TestCreateInvalid(t *testing.T) {
	produkMock := new(mocking.ProdukRepositoryMock)
	loggingMock := new(mocking.LoggingMock)
	validate := validator.New()

	svc := service.NewProdukServiceImpl(produkMock, validate, loggingMock)

	req := web.CreateProdukRequest{
		Nama:      "chiki",
		Deskripsi: "makanan",
		Harga:     "ljflksd",
		Kategori:  "snack",
	}

	resp, err := svc.Create(req)

	assert.NotNil(t, err)
	assert.Equal(t, web.ProdukResponse{}, resp)

	produkMock.AssertNotCalled(t, "Create", mock.Anything)
}

func TestCreateDuplicate(t *testing.T) {
	produkMock := new(mocking.ProdukRepositoryMock)
	loggingMock := new(mocking.LoggingMock)
	validate := validator.New()

	svc := service.NewProdukServiceImpl(produkMock, validate, loggingMock)

	req := web.CreateProdukRequest{
		Nama:      "chiki",
		Deskripsi: "makanan",
		Harga:     "10000",
		Kategori:  "snack",
	}

	err := errors.New("duplicate data")
	produkMock.On("Create", mock.Anything).Return(models.Produk{}, err)

	_, err = svc.Create(req)

	assert.NotNil(t, err)
	assert.Equal(t, "duplicate data", err.Error())

	produkMock.AssertExpectations(t)
}

func TestCreateInternalError(t *testing.T) {
	produkMock := new(mocking.ProdukRepositoryMock)
	loggingMock := new(mocking.LoggingMock)
	validate := validator.New()

	svc := service.NewProdukServiceImpl(produkMock, validate, loggingMock)

	req := web.CreateProdukRequest{
		Nama:      "chiki",
		Deskripsi: "makanan",
		Harga:     "10000",
		Kategori:  "snack",
	}

	err := errors.New("internal error")
	produkMock.On("Create", mock.Anything).Return(models.Produk{}, err)

	_, err = svc.Create(req)

	assert.NotNil(t, err)
	assert.Equal(t, "internal error", err.Error())

	produkMock.AssertExpectations(t)
}

func TestFindByIdSuccess(t *testing.T) {
	produkMock := new(mocking.ProdukRepositoryMock)
	loggingMock := new(mocking.LoggingMock)
	validate := validator.New()

	svc := service.NewProdukServiceImpl(produkMock, validate, loggingMock)

	id := 29
	harga := "10000"
	hrg, _ := decimal.NewFromString(harga)
	expected := models.Produk{
		Id: uint(id),
		Nama:      "chiki",
		Deskripsi: "makanan",
		Harga:     hrg,
		Kategori:  "snack",
	}

	produkMock.On("FindById", id).Return( expected, nil)

	resp, err := svc.FindById(id)

	assert.Nil(t, err)
	assert.Equal(t, resp.ID, uint(id))
	assert.Equal(t, resp.Harga, hrg)
	assert.Equal(t, resp.Deskripsi, "makanan")
	assert.Equal(t, resp.Kategori, "snack")

	produkMock.AssertExpectations(t)
}

func TestFindByIdNotFound(t *testing.T) {
	produkMock := new(mocking.ProdukRepositoryMock)
	loggingMock := new(mocking.LoggingMock)
	validate := validator.New()

	svc := service.NewProdukServiceImpl(produkMock, validate, loggingMock)

	id := 29

	produkMock.On("FindById", id).Return(models.Produk{}, repository.ErrorIdNotFound)

	_, err := svc.FindById(id)

	assert.NotNil(t, err)
	assert.Equal(t, repository.ErrorIdNotFound, err)

	produkMock.AssertExpectations(t)
}

func TestFindByIdInternalError(t *testing.T) {
	produkMock := new(mocking.ProdukRepositoryMock)
	loggingMock := new(mocking.LoggingMock)
	validate := validator.New()

	svc := service.NewProdukServiceImpl(produkMock, validate, loggingMock)

	id := 29

	err := errors.New("internal error")
	produkMock.On("FindById", id).Return(models.Produk{}, err)

	_, err = svc.FindById(id)

	assert.NotNil(t, err)
	assert.Equal(t, "internal error", err.Error())

	produkMock.AssertExpectations(t)
}