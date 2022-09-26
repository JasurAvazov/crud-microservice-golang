package views

import (
	"apelsin/models"
	"time"
)

type payment struct {
	Id     int       `json:"id"`
	Time   time.Time `json:"time"`
	Amount float64   `json:"amount"`
	InvId  int       `json:"inv_id"`
}

func Payment(model models.Payment) payment {
	return payment{
		Id:     model.Id,
		Time:   model.Time,
		Amount: model.Amount,
		InvId:  model.InvId,
	}
}

func Payments(models []models.Payment) []payment {
	r := make([]payment, len(models))
	for i := range models {
		r[i] = Payment(models[i])
	}
	return r
}
