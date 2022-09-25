package views

import "apelsin/models"

type category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func Category(model models.Category) category {
	return category{
		Id:   model.Id,
		Name: model.Name,
	}
}

func Categories(models []models.Category) []category {
	r := make([]category, len(models))
	for i := range models {
		r[i] = Category(models[i])
	}
	return r
}
