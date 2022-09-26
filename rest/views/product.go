package views

import "apelsin/models"

type product struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	CategoryId  int     `json:"category_id"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Photo       string  `json:"photo"`
}

func Product(model models.Product) product {
	return product{
		Id:          model.Id,
		Name:        model.Name,
		CategoryId:  model.CategoryId,
		Description: model.Description,
		Price:       model.Price,
		Photo:       model.Photo,
	}
}

func Products(models []models.Product) []product {
	r := make([]product, len(models))
	for i := range models {
		r[i] = Product(models[i])
	}
	return r
}
