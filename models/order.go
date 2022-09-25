package models

import "time"

type Order struct {
	Id      int
	Date    time.Time
	Cust_id int
}

func (n Order) GetCode() string {
	return string(rune(n.Id))
}

func (n Order) GetDescription() Translatable {
	return nil
}

func (n Order) GetExternalCodes() map[string]string {
	return nil
}
