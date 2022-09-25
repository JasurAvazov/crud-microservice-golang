package sqlstorage

import (
	"apelsin/models"
	"database/sql"
	"encoding/json"
)

type dbDistrict struct {
	Code          string          `db:"code"`
	SOATO         string          `db:"soato_code"`
	CodeGNI       string          `db:"gni_code"`
	CodeProvince  string          `db:"province_code"`
	Title         json.RawMessage `db:"title"`
	ActivatedAt   sql.NullTime    `db:"activated_at"`
	DeactivatedAt sql.NullTime    `db:"deactivated_at"`
	State         bool            `db:"state"`
}

func (dd dbDistrict) toModel() models.District {
	m := models.District{
		CodeSOATO:    dd.SOATO,
		Code:         dd.Code,
		CodeGNI:      dd.CodeGNI,
		CodeProvince: dd.CodeProvince,
		Title:        models.Languages{},
		State:        dd.State,
	}
	_ = json.Unmarshal(dd.Title, &m.Title)
	return m
}

func newDbDistrict(d models.District) dbDistrict {
	rawJson, _ := json.Marshal(d.Title)
	dd := dbDistrict{
		SOATO:        d.CodeSOATO,
		Code:         d.Code,
		CodeGNI:      d.CodeGNI,
		CodeProvince: d.CodeProvince,
		Title:        rawJson,
		ActivatedAt: sql.NullTime{
			Valid: d.ActivatedAt != nil,
			Time:  *d.ActivatedAt},
		State: d.State,
	}
	return dd
}
