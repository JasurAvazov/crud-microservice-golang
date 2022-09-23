package views

import "apelsin/models"

type offer struct {
	Code           string    `json:"code"`
	Name           Languages `json:"name"`
	ShouldBeSigned bool      `json:"should_be_signed"`
	SignatureType  string    `json:"signature_type"`
}

func Offer(model models.Offer) offer {
	return offer{
		Code: model.Code,
		Name: Languages{
			Ru: model.Name.GetRu(),
			En: model.Name.GetEn(),
			Uz: model.Name.GetUz(),
		},
		ShouldBeSigned: model.ShouldBeSigned,
		SignatureType:  model.SignatureType,
	}
}

func Offers(models []models.Offer) []offer {
	r := make([]offer, len(models))
	for i := range models {
		r[i] = Offer(models[i])
	}
	return r
}
