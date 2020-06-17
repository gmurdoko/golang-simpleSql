package models

import "database/sql"

type Bill struct {
	BillId    int
	ProductId int
	Sales     float64
	Tax       float64
}

func AllBill(db *sql.DB) ([]*Bill, error) {
	rows, err := db.Query("SELECT * FROM bill")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	bills := make([]*Bill, 0)

	for rows.Next() {
		//new => reserve 1 memory allocation with certain data type
		b := new(Bill)
		err := rows.Scan(&b.BillId, &b.ProductId, &b.Sales, &b.Tax)
		if err != nil {
			return nil, err
		}
		bills = append(bills, b)
	}

	return bills, nil

}
