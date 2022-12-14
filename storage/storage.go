package storage

import (
	"product-service/storage/postgres"
	"product-service/storage/repo"

	"github.com/jmoiron/sqlx"
)

type IStorage interface {
	Product() repo.ProductStorageI
}

type storagePg struct {
	db          *sqlx.DB
	productRepo repo.ProductStorageI
}

func NewStorage(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:          db,
		productRepo: postgres.NewProductRepo(db),
	}
}

func (s storagePg) Product() repo.ProductStorageI {
	return s.productRepo
}
