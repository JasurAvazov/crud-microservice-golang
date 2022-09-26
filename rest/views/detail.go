package views

import "apelsin/models"

type detail struct {
	Id       int   `json:"id"`
	OrdId    int   `json:"ord_id"`
	PrId     int   `json:"pr_id"`
	Quantity int16 `json:"quantity"`
}

func Detail(model models.Detail) detail {
	return detail{
		Id:       model.Id,
		OrdId:    model.OrdId,
		PrId:     model.PrId,
		Quantity: model.Quantity,
	}
}

func Details(models []models.Detail) []detail {
	r := make([]detail, len(models))
	for i := range models {
		r[i] = Detail(models[i])
	}
	return r
}
