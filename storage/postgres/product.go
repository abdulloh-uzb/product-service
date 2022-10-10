package postgres

import (
	pb "product-service/genproto"

	"github.com/jmoiron/sqlx"
)

type productRepo struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) *productRepo {
	return &productRepo{db: db}
}

func (r *productRepo) CreateProduct(req *pb.ProductRequest) (*pb.Product, error) {
	productResp := pb.Product{}
	err := r.db.QueryRow(`insert into products (name, price, type, category) values($1,$2,$3,$4) returning id, name, price, type, category`,
		req.Name, req.Price, req.TypeId, req.CategoryId).Scan(&productResp.Id, &productResp.Name, &productResp.Price, &productResp.TypeId, &productResp.CategoryId)
	if err != nil {
		return &pb.Product{}, err
	}
	return &productResp, nil
}

func (r *productRepo) UpdateProduct(req *pb.Product) (*pb.Product, error) {
	productResp := pb.Product{}

	err := r.db.QueryRow(`
	UPDATE products
	SET name = $1, price = $2, type = $3, category = $4 
	WHERE id = $5 
	returning id, name, price, type, category`,
		req.Name, req.Price, req.TypeId, req.CategoryId, req.Id).
		Scan(&productResp.Id, &productResp.Name, &productResp.Price, &productResp.TypeId, &productResp.CategoryId)
	if err != nil {
		return &pb.Product{}, err
	}

	return &productResp, nil

}
