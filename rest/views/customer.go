package views

import "apelsin/models"

type customer struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

func Customer(model models.Customer) customer {
	return customer{
		Id:      model.Id,
		Name:    model.Name,
		Country: model.Country,
		Address: model.Address,
		Phone:   model.Phone,
	}
}

func Customers(models []models.Customer) []customer {
	r := make([]customer, len(models))
	for i := range models {
		r[i] = Customer(models[i])
	}
	return r
}
