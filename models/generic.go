package models

type Translatable interface {
	GetRu() string
	GetEn() string
	GetUz() string
}

type Record interface {
	GetCode() string
	GetDescription() Translatable
	GetExternalCodes() map[string]string
}
