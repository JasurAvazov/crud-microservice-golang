package models

import "time"

type District struct {
	CodeSOATO    string
	Code         string
	CodeGNI      string
	CodeProvince string
	Title        Languages
	State        bool
	ActivatedAt  *time.Time
}

func (n District) GetCode() string {
	return n.Code
}

func (n District) GetDescription() Translatable {
	return n.Title
}

func (n District) GetExternalCodes() map[string]string {
	return nil
}
