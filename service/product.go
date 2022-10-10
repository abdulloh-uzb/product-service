package service

import (
	"context"
	"product-service/storage"

	l "product-service/pkg/logger"

	pb "product-service/genproto"

	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductService struct {
	storage storage.IStorage
	logger  l.Logger
}

func NewProductService(db *sqlx.DB, log l.Logger) *ProductService {
	return &ProductService{
		storage: storage.NewStorage(db),
		logger:  log,
	}
}

func (s *ProductService) CreateProduct(ctx context.Context, req *pb.ProductRequest) (*pb.Product, error) {
	product, err := s.storage.Product().CreateProduct(req)
	if err != nil {
		s.logger.Error("error insert", l.Any("error while insert product", err))
		return &pb.Product{}, status.Error(codes.Internal, "something went wrong, please check product create func")

	}
	return product, nil
}

func (s *ProductService) UpdateProduct(ctx context.Context, req *pb.Product) (*pb.Product, error) {
	product, err := s.storage.Product().UpdateProduct(req)
	if err != nil {
		s.logger.Error("error insert", l.Any("error while insert product", err))
		return &pb.Product{}, status.Error(codes.Internal, "something went wrong, please check product update func")
	}
	return product, nil
}
