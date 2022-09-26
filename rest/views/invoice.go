package views

import (
	"apelsin/models"
	"time"
)

type invoice struct {
	Id     int       `json:"id"`
	OrdId  int       `json:"ord_id"`
	Amount float64   `json:"amount"`
	Issued time.Time `json:"issued"`
	Due    time.Time `json:"due"`
}

func Invoice(model models.Invoice) invoice {
	return invoice{
		Id:     model.Id,
		OrdId:  model.OrdId,
		Amount: model.Amount,
		Issued: model.Issued,
		Due:    model.Due,
	}
}

func Invoices(models []models.Invoice) []invoice {
	r := make([]invoice, len(models))
	for i := range models {
		r[i] = Invoice(models[i])
	}
	return r
}
