package models

import (
	"database/sql"
	guuid "github.com/google/uuid"
	"log"
	"strings"
)

type Product struct {
	ProductId       string
	ProductCode     string
	ProductName     string
	ProductCategory Category
}

func AllProduct(db *sql.DB, pageNo, totalPerPage int) ([]*Product, error) {
	//it is a good practice to always use the LIMIT clause with the ORDER BY clause to constraint the result rows in unique order.
	rows, err := db.Query(`
		SELECT p.id,p.product_code,p.product_name,p.category_id,c.category_name
		FROM product p join category c on p.category_id = c.id 
		ORDER BY p.id 
		LIMIT ?,?
		`, pageNo, totalPerPage)
	if err != nil {
		log.Fatalf("%v", err)
		return nil, err
	}

	defer rows.Close()

	products := make([]*Product, 0)

	for rows.Next() {
		//new => reserve 1 memory allocation with certain data type pb := new(Product)
		p := new(Product)
		//c := new(Category)
		err := rows.Scan(&p.ProductId, &p.ProductCode, &p.ProductName, &p.ProductCategory.CateogryId, &p.ProductCategory.CategoryName)
		if err != nil {
			log.Fatalf("%v", err)
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func FindProductIn(db *sql.DB, ids []string) ([]*Product, error) {
	sql := `
		SELECT id,product_code,product_name
		FROM product 
		WHERE product_code IN(?` + strings.Repeat(",?", len(ids)-1) + `)`
	stmt, err := db.Prepare(sql)

	params := make([]interface{}, len(ids))
	for i, v := range ids {
		params[i] = v
	}
	rows, err := stmt.Query(params...)
	if err != nil {
		log.Fatalf("%v", err)
		return nil, err
	}

	defer stmt.Close()

	products := make([]*Product, 0)

	for rows.Next() {
		p := new(Product)
		err := rows.Scan(&p.ProductId, &p.ProductCode, &p.ProductName)
		if err != nil {
			log.Fatalf("%v", err)
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}
func CreateProduct(db *sql.DB, product Product) error {
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("%v", err)
		return err
	}
	stmt, err := db.Prepare("INSERT INTO product(id,product_code,product_name,category_id)  VALUES(?, ?, ?, ?)")
	if err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return err
	}
	defer stmt.Close()

	id := guuid.New()
	if _, err := stmt.Exec(id, product.ProductCode, product.ProductName, product.ProductCategory.CateogryId); err != nil {
		tx.Rollback()
		log.Fatalf("%v", err)
		return err
	}
	return tx.Commit()
}
