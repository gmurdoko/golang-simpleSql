package services

import (
	"database/sql"
	"github.com/edwardsuwirya/simpleSql/models"
)

type ProductService struct {
	db *sql.DB
}

func NewProductService(db *sql.DB) *ProductService {
	return &ProductService{db}
}

func (ps *ProductService) GetAllProduct(pageNo, totalPerPage int) []*models.Product {
	products, err := models.GetAllProduct(ps.db, totalPerPage*(pageNo-1), totalPerPage)
	if err != nil {
		return nil
	}
	return products
}
