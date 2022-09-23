package controller

import (
	"apelsin/errs"
	"io"

	"github.com/xuri/excelize/v2"
)

func (ctrl *Controller) _parseFile(file io.Reader, validator headingValidator) ([][]string, error) {
	xls, err := excelize.OpenReader(file)
	if err != nil {
		return nil, errs.Errf(errs.ErrWrongInput, err.Error())
	}
	sheet := xls.GetSheetName(0)
	if sheet == "" {
		return nil, errs.Errf(errs.ErrWrongInput, "No sheets found")
	}
	rows, err := xls.GetRows(sheet)
	if err != nil {
		return nil, errs.Errf(errs.ErrWrongInput, err.Error())
	}
	if err := validator.Validate(rows[0]); err != nil {
		return nil, err
	}
	return rows, nil
}
