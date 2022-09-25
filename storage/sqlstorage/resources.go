package sqlstorage

import (
	"apelsin/models"
	"time"
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

type dbOrder struct {
	Id      int       `db:"id"`
	Date    time.Time `db:"date"`
	Cust_id int       `db:"cust_id"`
}

func (dd dbOrder) toModel() models.Order {
	m := models.Order{
		Id:      dd.Id,
		Date:    dd.Date,
		Cust_id: dd.Cust_id,
	}
	return m
}

func newDbOrder(m models.Order) dbOrder {
	dd := dbOrder{
		Id:      m.Id,
		Date:    m.Date,
		Cust_id: m.Cust_id,
	}
	return dd
}

type dbCategory struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

func (dd dbCategory) toModel() models.Category {
	m := models.Category{
		Id:   dd.Id,
		Name: dd.Name,
	}
	return m
}

func newDbCategory(m models.Category) dbCategory {
	dd := dbCategory{
		Id:   m.Id,
		Name: m.Name,
	}
	return dd
}
