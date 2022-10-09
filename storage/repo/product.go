package repo

import (
	pb "product-service/genproto"
)

type ProductStorageI interface {
	CreateProduct(*pb.ProductRequest) (*pb.Product, error)
}
