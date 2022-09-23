package controller

import (
	"apelsin/errs"
	"strings"
)

type headingValidator interface {
	Validate(fields []string) error
}

type xlsValidator struct {
	fields map[int]string
}

var regionValidator = &xlsValidator{fields: map[int]string{
	0: "Code",
	1: "NameRu",
	2: "NameEn",
	3: "NameUz",
	4: "Rank",
}}
var nationalityValidator = &xlsValidator{fields: map[int]string{
	0: "Code",
	1: "NameRu",
	2: "NameEn",
	3: "NameUz",
}}
var districtValidator = &xlsValidator{fields: map[int]string{
	0: "CodeSOATO",
	1: "Code",
	2: "CodeGNI",
	3: "CodeProvince",
	4: "NameRu",
	5: "NameEn",
	6: "NameUz",
}}
var countryValidator = &xlsValidator{fields: map[int]string{
	0: "Code",
	1: "NameRu",
	2: "NameEn",
	3: "NameUz",
	4: "AlphaDuo",
	5: "AlphaTriple",
	6: "CurrencyCode",
	7: "LocationSign",
}}

func (r *xlsValidator) Validate(fields []string) error {
	if len(fields) < len(r.fields) {
		return errs.Errf(errs.ErrWrongInput, "Not validated")
	}
	for i, str := range fields {
		if i == len(r.fields) {
			break
		}
		if strings.ToLower(str) != strings.ToLower(r.fields[i]) {
			return errs.Errf(errs.ErrWrongInput, "Wrong Heading: %+v", fields)
		}
	}
	return nil
}
