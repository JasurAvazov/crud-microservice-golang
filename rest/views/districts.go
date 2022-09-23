package views

import "apelsin/models"

type district struct {
	CodeSOATO    string    `json:"code_soato"`
	Code         string    `json:"code"`
	CodeGNI      string    `json:"code_gni"`
	CodeProvince string    `json:"code_province"`
	Title        Languages `json:"title"`
	IsActive     bool      `json:"is_active"`
}

func District(model models.District) district {
	return district{
		CodeSOATO:    model.CodeSOATO,
		Code:         model.Code,
		CodeGNI:      model.CodeGNI,
		CodeProvince: model.CodeProvince,
		Title: Languages{
			Ru: model.Title.GetRu(),
			En: model.Title.GetEn(),
			Uz: model.Title.GetUz(),
		},
	}
}

func Districts(models []models.District) []district {
	r := make([]district, len(models))
	for i := range models {
		r[i] = District(models[i])
	}
	return r
}
