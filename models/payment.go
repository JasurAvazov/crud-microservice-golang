package models

import "time"

type Payment struct {
	Id     int
	Time   time.Time
	Amount float64
	InvId  int
}

func (n Payment) GetCode() string {
	return string(rune(n.Id))
}

func (n Payment) GetDescription() Translatable {
	return nil
}

func (n Payment) GetExternalCodes() map[string]string {
	return nil
}
