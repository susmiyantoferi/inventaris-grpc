package controller

import (
	"context"
	"errors"
	"inventaris/inventarispb/produkpb"
	"inventaris/repository"
	"inventaris/service"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProdukGRPCServer struct {
	produkpb.UnimplementedProdukServiceServer
	ProdukService service.ProdukService
}

func NewProdukGRPCServer(produkService service.ProdukService) *ProdukGRPCServer{
	return &ProdukGRPCServer{
		ProdukService: produkService,
	}
} 

func(s *ProdukGRPCServer) GetProdukById(ctx context.Context, r *produkpb.GetProdukReq) (*produkpb.Produk, error){
	produk, err := s.ProdukService.FindById(int(r.GetId()))
	if err != nil {
		if errors.Is(err, repository.ErrorIdNotFound) {
			return nil, status.Errorf(codes.NotFound, "produk id:%v not found", r.GetId())
		}
		return nil, status.Errorf(codes.Internal, "internal error %v", err)
	}

	hargaCvt := produk.Harga.InexactFloat64()

	produkMap := &produkpb.Produk{
		Id: int64(produk.ID),
		Nama: produk.Nama,
		Deskripsi: produk.Deskripsi,
		Harga: hargaCvt,
		Kategori: produk.Kategori,
		Gambar: produk.Gambar,
	}

	return produkMap, nil
	
}

func(s *ProdukGRPCServer) GetAllProduk(ctx context.Context, r *produkpb.Empty) (*produkpb.ProdukList, error){
	produk, err := s.ProdukService.FindAll()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch produk: %v", err)
	}

	var produkList []*produkpb.Produk
	for _, v := range produk {
		pbProduk := &produkpb.Produk{
			Id: int64(v.ID),
			Nama: v.Nama,
			Deskripsi: v.Deskripsi,
			Harga: v.Harga.InexactFloat64(),
			Kategori: v.Kategori,
			Gambar: v.Gambar,
		}

		produkList = append(produkList, pbProduk)
	}

	return &produkpb.ProdukList{
		Data: produkList,
	}, nil
	
}