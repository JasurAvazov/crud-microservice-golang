package models

type Category struct {
	Id   int
	Name string
}

func (n Category) GetCode() string {
	return string(rune(n.Id))
}

func (n Category) GetDescription() Translatable {
	return nil
}

func (n Category) GetExternalCodes() map[string]string {
	return nil
}
