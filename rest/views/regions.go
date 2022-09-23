package views

import "apelsin/models"

type region struct {
	Code     string    `json:"code"`
	Title    Languages `json:"title"`
	Rank     string    `json:"rank"`
	IsActive bool      `json:"is_active"`
}

func Region(model models.Region) region {
	return region{
		Code: model.Code,
		Title: Languages{
			Ru: model.Title.GetRu(),
			En: model.Title.GetEn(),
			Uz: model.Title.GetUz(),
		},
		Rank:     model.Rank,
		IsActive: model.ActivatedAt != nil,
	}
}

func Regions(models []models.Region) []region {
	r := make([]region, len(models))
	for i := range models {
		r[i] = Region(models[i])
	}
	return r
}
