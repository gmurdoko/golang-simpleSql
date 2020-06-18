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

func (ps *ProductService) GetProducts(pageNo, totalPerPage int) []*models.Product {
	products, err := models.AllProduct(ps.db, totalPerPage*(pageNo-1), totalPerPage)
	if err != nil {
		return nil
	}
	return products
}

func (ps *ProductService) CreateAProduct(productCode string, productName string, category string) *models.Product {
	product := models.Product{
		ProductCode:     productCode,
		ProductName:     productName,
		ProductCategory: models.Category{CateogryId: category},
	}
	err := models.CreateProduct(ps.db, product)
	if err != nil {
		return nil
	}
	return &product
}
