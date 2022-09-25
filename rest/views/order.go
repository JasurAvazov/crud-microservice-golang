package views

import (
	"apelsin/models"
	"time"
)

type order struct {
	Id      int       `json:"id"`
	Date    time.Time `json:"date"`
	Cust_id int       `json:"cust_id"`
}

func Order(model models.Order) order {
	return order{
		Id:      model.Id,
		Date:    model.Date,
		Cust_id: model.Cust_id,
	}
}

func Orders(models []models.Order) []order {
	r := make([]order, len(models))
	for i := range models {
		r[i] = Order(models[i])
	}
	return r
}
