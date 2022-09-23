package models

import "time"

type Nationality struct {
	Code          string
	Title         Languages
	ActivatedAt   *time.Time
	DeactivatedAt *time.Time
	State         bool
}

func (n Nationality) GetCode() string {
	return n.Code
}

func (n Nationality) GetDescription() Translatable {
	return n.Title
}

func (n Nationality) GetExternalCodes() map[string]string {
	return nil
}
