package models

type Offer struct {
	Code           string
	Name           Languages
	ShouldBeSigned bool
	SignatureType  string
	CreatedAt      string
}

type FullOffer struct {
	Code           string
	VersionCode    string
	Name           string
	ShouldBeSigned bool
	SignatureType  string
	Content        string
}

func (n Offer) GetCode() string {
	return n.Code
}

func (n Offer) GetDescription() Translatable {
	return n.Name
}

func (n Offer) GetExternalCodes() map[string]string {
	return nil
}
