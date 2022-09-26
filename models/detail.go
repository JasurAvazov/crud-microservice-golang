package models

type Detail struct {
	Id       int
	OrdId    int
	PrId     int
	Quantity int16
}

func (n Detail) GetCode() string {
	return string(rune(n.Id))
}

func (n Detail) GetDescription() Translatable {
	return nil
}

func (n Detail) GetExternalCodes() map[string]string {
	return nil
}
