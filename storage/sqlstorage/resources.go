package sqlstorage

import (
	"apelsin/models"
)

type dbCustomer struct {
	Id      int    `db:"id"`
	Name    string `db:"name"`
	Country string `db:"country"`
	Address string `db:"address"`
	Phone   string `db:"phone"`
}

func (dd dbCustomer) toModel() models.Customer {
	m := models.Customer{
		Id:      dd.Id,
		Name:    dd.Name,
		Country: dd.Country,
		Address: dd.Address,
		Phone:   dd.Phone,
	}
	return m
}

func newDbCustomer(m models.Customer) dbCustomer {
	dd := dbCustomer{
		Id:      m.Id,
		Name:    m.Name,
		Country: m.Country,
		Address: m.Address,
		Phone:   m.Phone,
	}
	return dd
}
