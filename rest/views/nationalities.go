package views

import "apelsin/models"

type nationality struct {
	Code     string    `json:"code"`
	Title    Languages `json:"title"`
	IsActive bool      `json:"is_active"`
}

func Nationality(model models.Nationality) nationality {
	return nationality{
		Code: model.Code,
		Title: Languages{
			Ru: model.Title.GetRu(),
			En: model.Title.GetEn(),
			Uz: model.Title.GetUz(),
		},
		IsActive: model.ActivatedAt != nil,
	}
}

func Nationalities(models []models.Nationality) []nationality {
	r := make([]nationality, len(models))
	for i := range models {
		r[i] = Nationality(models[i])
	}
	return r
}
