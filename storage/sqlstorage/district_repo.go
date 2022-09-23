package sqlstorage

import (
	"apelsin/errs"
	"apelsin/models"
	"context"
)

type districtRepo struct {
	s *Store
}

func (repo *districtRepo) GetRecords(ctx context.Context) ([]models.Record, error) {
	b, err := repo.ReadAll(ctx)
	if err != nil {
		return nil, err
	}
	records := make([]models.Record, len(b))
	for i, v := range b {
		records[i] = models.Record(v)
	}
	return records, nil
}

func (repo *districtRepo) Create(ctx context.Context, district models.District) (models.District, error) {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	q := `INSERT INTO districts (code,soato_code,gni_code,province_code,
                       title,activated_at,deactivated_at) 
			VALUES ($1,$2,$3,$4,$5,$6,$7)`
	row := newDbDistrict(district)
	_, err := sqlClient.Exec(q,
		row.Code, row.SOATO,
		row.CodeGNI, row.CodeProvince,
		row.Title, row.ActivatedAt, row.DeactivatedAt,
	)
	if err != nil {
		return district, err
	}
	return district, nil
}

func (repo *districtRepo) Read(ctx context.Context, code string) (models.District, error) {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	q := `SELECT * FROM districts WHERE code=$1`
	row := dbDistrict{}
	if err := sqlClient.Get(&row, q, code); err != nil {
		return models.District{}, err
	}
	return row.toModel(), nil
}

func (repo *districtRepo) ReadAll(ctx context.Context) ([]models.District, error) {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	rows := make([]dbDistrict, 0)
	q := `SELECT * FROM districts`
	if err := sqlClient.Select(&rows, q); err != nil {
		return nil, err
	}
	districts := make([]models.District, len(rows))
	for i, row := range rows {
		districts[i] = row.toModel()
	}
	return districts, nil
}

func (repo *districtRepo) Update(ctx context.Context, district models.District) error {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	q := `UPDATE districts SET 
		    soato_code=$1,gni_code=$2,
		 	province_code=$3,title=$4,
		 	activated_at=$5,deactivated_at=$6
		 WHERE code=$7`
	row := newDbDistrict(district)
	res, err := sqlClient.Exec(q,
		row.SOATO, row.CodeGNI, row.CodeProvince, row.Title, row.ActivatedAt, row.DeactivatedAt, row.Code)
	if err != nil {
		return err
	}
	if c, _ := res.RowsAffected(); c == 0 {
		return errs.Errf(errs.ErrNotFound, "No rows affected")
	}
	return nil
}

func (repo *districtRepo) Delete(ctx context.Context, code string) error {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	q := `DELETE FROM districts WHERE code=$1`
	res, err := sqlClient.Exec(q, code)
	if err != nil {
		return err
	}
	if c, _ := res.RowsAffected(); c == 0 {
		return errs.Errf(errs.ErrNotFound, "No rows affected")
	}
	return nil
}

func (repo *districtRepo) ClearAll(ctx context.Context) error {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	q := `DELETE FROM districts WHERE TRUE`
	if _, err := sqlClient.Exec(q); err != nil {
		return err
	}
	return nil
}
