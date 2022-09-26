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

type dbProduct struct {
	Id          int     `db:"id"`
	Name        string  `db:"name"`
	CategoryId  int     `db:"category_id"`
	Description string  `db:"description"`
	Price       float64 `db:"price"`
	Photo       string  `db:"photo"`
}

func (dd dbProduct) toModel() models.Product {
	m := models.Product{
		Id:          dd.Id,
		Name:        dd.Name,
		CategoryId:  dd.CategoryId,
		Description: dd.Description,
		Price:       dd.Price,
		Photo:       dd.Photo,
	}
	return m
}

func newDbProduct(m models.Product) dbProduct {
	dd := dbProduct{
		Id:          m.Id,
		Name:        m.Name,
		CategoryId:  m.CategoryId,
		Description: m.Description,
		Price:       m.Price,
		Photo:       m.Photo,
	}
	return dd
}

type dbDetail struct {
	Id       int   `db:"id"`
	OrdId    int   `db:"ord_id"`
	PrId     int   `db:"pr_id"`
	Quantity int16 `db:"quantity"`
}

func (dd dbDetail) toModel() models.Detail {
	m := models.Detail{
		Id:       dd.Id,
		OrdId:    dd.OrdId,
		PrId:     dd.PrId,
		Quantity: dd.Quantity,
	}
	return m
}

func newDbDetail(m models.Detail) dbDetail {
	dd := dbDetail{
		Id:       m.Id,
		OrdId:    m.OrdId,
		PrId:     m.PrId,
		Quantity: m.Quantity,
	}
	return dd
}

type dbInvoice struct {
	Id     int       `db:"id"`
	OrdId  int       `db:"ord_id"`
	Amount float64   `db:"amount"`
	Issued time.Time `db:"issued"`
	Due    time.Time `db:"due"`
}

func (dd dbInvoice) toModel() models.Invoice {
	m := models.Invoice{
		Id:     dd.Id,
		OrdId:  dd.OrdId,
		Amount: dd.Amount,
		Issued: dd.Issued,
		Due:    dd.Due,
	}
	return m
}

func newDbInvoice(m models.Invoice) dbInvoice {
	dd := dbInvoice{
		Id:     m.Id,
		OrdId:  m.OrdId,
		Amount: m.Amount,
		Issued: m.Issued,
		Due:    m.Due,
	}
	return dd
}
