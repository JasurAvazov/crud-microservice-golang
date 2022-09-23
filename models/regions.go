package models

import "time"

type Region struct {
	Code          string
	Title         Languages
	Rank          string
	ActivatedAt   *time.Time
	DeactivatedAt *time.Time
	State         bool
}

func (n Region) GetCode() string {
	return n.Code
}

func (n Region) GetDescription() Translatable {
	return n.Title
}

func (n Region) GetExternalCodes() map[string]string {
	return nil
}
