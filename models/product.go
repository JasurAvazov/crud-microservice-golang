package models

type Product struct {
	Id          int
	Name        string
	CategoryId  int
	Description string
	Price       float64
	Photo       string
}

func (n Product) GetCode() string {
	return string(rune(n.Id))
}

func (n Product) GetDescription() Translatable {
	return nil
}

func (n Product) GetExternalCodes() map[string]string {
	return nil
}
