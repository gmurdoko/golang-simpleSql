package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/edwardsuwirya/simpleSql/models"
	"github.com/edwardsuwirya/simpleSql/services"
	"github.com/edwardsuwirya/simpleSql/utils"
)

type Env struct {
	db *sql.DB
}

var (
	DB_USER,
	DB_PASSWORD,
	DB_HOST,
	DB_PORT,
	SCHEMA_NAME string
)

func main() {
	DB_USER = utils.GetEnv("DB_USER", "root")
	DB_PASSWORD = utils.GetEnv("DB_PASSWORD", "toor")
	DB_HOST = utils.GetEnv("DB_HOST", "localhost")
	DB_PORT = utils.GetEnv("DB_PORT", "3306")
	SCHEMA_NAME = utils.GetEnv("DB_SCHEMA", "schema")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, SCHEMA_NAME)
	db, err := models.InitDB(dataSourceName)
	if err != nil {
		log.Panic(err)
	}
	env := &Env{db: db}

	//billService := services.NewBillService(env.db)
	//newBill := billService.CreateABill(8, 4, 12000, 8)
	//if newBill != nil {
	//	log.Print(*newBill)
	//}
	//sales := billService.TotalSales()
	//log.Printf("%v", humanize.Commaf(sales))

	productService := services.NewProductService(env.db)
	products := productService.GetAllProduct(1, 2)
	for _, p := range products {
		log.Printf("%v %v %v %v %v", p.ProductId, p.ProductCode, p.ProductName, p.ProductCategory.CateogryId, p.ProductCategory.CategoryName)
	}
}
