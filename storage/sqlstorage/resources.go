package sqlstorage

import (
	"apelsin/models"
	"database/sql"
	"encoding/json"
	"reflect"
	"time"
)

type dbFullOffer struct {
	Code           string         `db:"code"`
	ShouldBeSigned bool           `db:"should_be_signed"`
	SignatureType  sql.NullString `db:"signature_type"`
	VersionCode    sql.NullString `db:"version_code"`
	Name           sql.NullString `db:"name"`
	Content        sql.NullString `db:"content"`
}

type dbOffer struct {
	Code           string          `db:"code"`
	Name           json.RawMessage `db:"name"`
	ShouldBeSigned bool            `db:"should_be_signed"`
	SignatureType  sql.NullString  `db:"signature_type"`
	CreatedAt      string          `db:"created_at"`
}

func newDbOffer(d models.Offer) dbOffer {
	rawJson, _ := json.Marshal(d.Name)
	dd := dbOffer{
		Code:           d.Code,
		Name:           rawJson,
		ShouldBeSigned: d.ShouldBeSigned,
		SignatureType: sql.NullString{
			String: d.SignatureType,
			Valid:  d.SignatureType != "",
		},
	}

	return dd
}

func (raw *dbOffer) toModel() models.Offer {
	var (
		name models.Languages
	)

	_ = json.Unmarshal(raw.Name, &name)

	result := models.Offer{
		Code:           raw.Code,
		Name:           name,
		ShouldBeSigned: raw.ShouldBeSigned,
		SignatureType:  raw.SignatureType.String,
	}
	return result
}

// structToJson ...
func structToJson(m interface{}) (map[string]interface{}, error) {
	var (
		st  map[string]interface{}
		err error
	)

	byt, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}

	if marErr := json.Unmarshal(byt, &st); marErr != nil {
		return nil, marErr
	}

	return st, nil
}

func (raw *dbOffer) merge(model dbOffer) {
	receiver := reflect.ValueOf(raw).Elem()
	sender := reflect.ValueOf(model)

	rec, _ := structToJson(model)

	for i := 0; i < receiver.Type().NumField(); i++ {
		var slice bool
		recType := receiver.Field(i).Kind().String()
		recValue := rec[receiver.Type().Field(i).Tag.Get("json")]

		if recType == "slice" && cap(recValue.([]interface{})) == 0 {
			slice = true
		}

		if receiver.Field(i).IsZero() || slice {
			f := receiver.FieldByName(receiver.Type().Field(i).Name)
			r := sender.FieldByName(sender.Type().Field(i).Name)
			f.Set(r)
		}

	}
}

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

type dbNationality struct {
	Code          string          `db:"code"`
	Title         json.RawMessage `db:"title"`
	ActivatedAt   sql.NullTime    `db:"activated_at"`
	DeactivatedAt sql.NullTime    `db:"deactivated_at"`
}

func (dn dbNationality) toModel() models.Nationality {
	m := models.Nationality{
		Code:  dn.Code,
		Title: models.Languages{},
	}
	_ = json.Unmarshal(dn.Title, &m.Title)
	return m
}

func newDbNationality(d models.Nationality) dbNationality {
	rawJson, _ := json.Marshal(d.Title)
	dd := dbNationality{
		Code:  d.Code,
		Title: rawJson,
	}
	return dd
}

type dbRegion struct {
	Code          string          `db:"code"`
	Title         json.RawMessage `db:"title"`
	Rank          string          `db:"rank"`
	ActivatedAt   sql.NullTime    `db:"activated_at"`
	DeactivatedAt sql.NullTime    `db:"deactivated_at"`
}

func (dr dbRegion) toModel() models.Region {
	var activatedAt, deactivatedAt *time.Time
	if dr.ActivatedAt.Valid {
		activatedAt = &dr.ActivatedAt.Time
	}
	if dr.DeactivatedAt.Valid {
		deactivatedAt = &dr.DeactivatedAt.Time
	}
	m := models.Region{
		Code:          dr.Code,
		Title:         models.Languages{},
		Rank:          dr.Rank,
		ActivatedAt:   activatedAt,
		DeactivatedAt: deactivatedAt,
	}
	_ = json.Unmarshal(dr.Title, &m.Title)
	return m
}

func newDbRegion(d models.Region) dbRegion {
	rawJson, _ := json.Marshal(d.Title)
	dd := dbRegion{
		Code:  d.Code,
		Title: rawJson,
		Rank:  d.Rank,
	}
	if d.DeactivatedAt != nil {
		dd.DeactivatedAt = sql.NullTime{
			Valid: d.DeactivatedAt != nil,
			Time:  *d.DeactivatedAt}
	}
	if d.ActivatedAt != nil {
		dd.ActivatedAt = sql.NullTime{
			Valid: d.ActivatedAt != nil,
			Time:  *d.ActivatedAt}
	}
	return dd
}
