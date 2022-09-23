package models

type Languages struct {
	Uz string `json:"uz"`
	Ru string `json:"ru"`
	En string `json:"en"`
}

func (l Languages) GetRu() string {
	return l.Ru
}

func (l Languages) GetEn() string {
	return l.En
}

func (l Languages) GetUz() string {
	return l.Uz
}
