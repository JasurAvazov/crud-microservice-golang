package models

type Customer struct {
	Id      int
	Name    string
	Country string
	Address string
	Phone   string
}

func (n Customer) GetCode() string {
	return string(rune(n.Id))
}

func (n Customer) GetDescription() Translatable {
	return nil
}

func (n Customer) GetExternalCodes() map[string]string {
	return nil
}
