package service

import (
	"errors"
	"inventaris/helper"
	log "inventaris/logging"
	"inventaris/models"
	"inventaris/repository"
	"inventaris/web"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type ProdukServiceImpl struct {
	ProdukRepository repository.ProdukRepository
	Validate         *validator.Validate
	Logging log.Logging
}

func NewProdukServiceImpl(produkRepository repository.ProdukRepository, validate *validator.Validate, logging log.Logging) *ProdukServiceImpl {
	return &ProdukServiceImpl{
		ProdukRepository: produkRepository,
		Validate:         validate,
		Logging: logging,
	}
}

func (p *ProdukServiceImpl) Create(r web.CreateProdukRequest) (web.ProdukResponse, error) {
	err := p.Validate.Struct(r)
	if err != nil {
		p.Logging.ErrInfo(err, "validation failed")
		return web.ProdukResponse{}, err
	}

	hargacvt, errs := decimal.NewFromString(r.Harga)
	if errs != nil {
		p.Logging.ErrInfo(err, "failed convert to decimal")
		return web.ProdukResponse{}, errs
	}

	produk := models.Produk{
		Nama:      r.Nama,
		Deskripsi: r.Deskripsi,
		Harga:     hargacvt,
		Kategori:  r.Kategori,
	}

	result, err := p.ProdukRepository.Create(produk)
	if err != nil {
		if errors.Is(err, gorm.ErrInvalidData){
			p.Logging.ErrInfo(err, "invalid input request" )
			return web.ProdukResponse{}, errors.New("bad request")
		}

		p.Logging.ErrInfo(err, "internal error")
		return web.ProdukResponse{}, err
	}
	p.Logging.MsgInfo("Created")
	response := helper.ToProdukResponse(result)

	return response, nil
}

func (p *ProdukServiceImpl) Update(r web.UpdateProdukRequest) (web.ProdukResponse, error) {

	err := p.Validate.Struct(r)
	if err != nil {
		return web.ProdukResponse{}, err
	}

	produk, errs := p.ProdukRepository.FindById(r.Id)
	if errs != nil {
		if errors.Is(errs, gorm.ErrRecordNotFound){
			return web.ProdukResponse{}, repository.ErrorIdNotFound
		}
		return web.ProdukResponse{}, errs
	}

	hargacvt, _ := decimal.NewFromString(r.Harga)

	produk.Nama = r.Nama
	produk.Deskripsi = r.Deskripsi
	produk.Harga = hargacvt
	produk.Kategori = r.Kategori

	result, err := p.ProdukRepository.Update(produk)
	if err != nil {
		return web.ProdukResponse{}, err
	}
	response := helper.ToProdukResponse(result)

	return response, nil
}

func (p *ProdukServiceImpl) Delete(produkId int) error {
	produk, err := p.ProdukRepository.FindById(produkId)
	if err != nil {
		return err
	}

	p.ProdukRepository.Delete(produkId)

	if produk.Gambar != "" {
		os.Remove("uploads/" + produk.Gambar)
	}

	return nil
}

func (p *ProdukServiceImpl) FindById(produkId int) (web.ProdukResponse, error) {
	produk, err := p.ProdukRepository.FindById(produkId)
	if err != nil {
		if errors.Is(err, repository.ErrorIdNotFound) {
			p.Logging.ErrInfo(err, "id not found")
			return web.ProdukResponse{}, repository.ErrorIdNotFound
		}
		p.Logging.ErrInfo(err, "internal error")
		return web.ProdukResponse{}, err
	}

	p.Logging.MsgInfo("success find id")
	response := helper.ToProdukResponse(produk)

	return response, nil
}

func (p *ProdukServiceImpl) FindAll() ([]web.ProdukResponse, error) {
	produk, err := p.ProdukRepository.FindAll()
	if err != nil {
		return []web.ProdukResponse{}, err
	}

	var responses []web.ProdukResponse
	for _, value := range produk {
		response := helper.ToProdukResponse(value)

		responses = append(responses, response)
	}

	return responses, nil
}

func (p *ProdukServiceImpl) UpdateImage(produkId int, gambar string) (web.ProdukResponse, error) {
	_, err := p.ProdukRepository.FindById(produkId)
	if err != nil {
		return web.ProdukResponse{}, err
	}

	result, err := p.ProdukRepository.UpdateImage(produkId, gambar)
	if err != nil {
		return web.ProdukResponse{}, err
	}

	response :=helper.ToProdukResponse(result)

	return response, nil
}
