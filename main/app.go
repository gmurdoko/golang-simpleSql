package main

import (
	"database/sql"
	"fmt"
	"github.com/edwardsuwirya/simpleSql/models"
	"log"
)

type Env struct {
	db *sql.DB
}

func main() {
	db, err := models.InitDB("root:P@ssw0rd@/enigma")
	if err != nil {
		log.Panic(err)
	}
	env := &Env{db: db}
	bills, err := models.AllBill(env.db)
	if err != nil{
		log.Panic(err)
	}
	for _, bill := range bills {
		fmt.Printf("%d, %d, %f, %f\n", bill.BillId, bill.ProductId, bill.Sales, bill.Tax)
	}
}
