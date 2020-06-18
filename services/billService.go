package services

import (
	"database/sql"
	"github.com/edwardsuwirya/simpleSql/models"
)

type BillService struct {
	db *sql.DB
}

func NewBillService(db *sql.DB) *BillService {
	return &BillService{db}
}

func (bs *BillService) CreateABill(billId int, productId int, sales float64, tax float64) *models.Bill {
	bill := models.Bill{
		BillId:    billId,
		ProductId: productId,
		Sales:     sales,
		Tax:       tax,
	}
	err := models.CreateBill(bs.db, bill)
	if err != nil {
		return nil
	}
	return &bill
}
func (bs *BillService) GetAllBill() []*models.Bill {
	bills, err := models.AllBill(bs.db)
	if err != nil {
		return nil
	}
	return bills
}

func (bs *BillService) TotalSales() float64 {
	bills, err := models.TotalSales(bs.db)
	if err != nil {
		return 0
	}
	return (*bills).Summary

}
