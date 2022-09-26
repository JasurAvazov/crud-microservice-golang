package models

import "time"

type Invoice struct {
	Id     int
	OrdId  int
	Amount float64
	Issued time.Time
	Due    time.Time
}

func (n Invoice) GetCode() string {
	return string(rune(n.Id))
}

func (n Invoice) GetDescription() Translatable {
	return nil
}

func (n Invoice) GetExternalCodes() map[string]string {
	return nil
}
