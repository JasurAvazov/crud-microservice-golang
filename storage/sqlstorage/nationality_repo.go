package sqlstorage

import (
	"apelsin/errs"
	"apelsin/models"
	"context"
)

type nationalityRepo struct {
	s *Store
}

func (repo *nationalityRepo) GetRecords(ctx context.Context) ([]models.Record, error) {
	b, err := repo.ReadAll(context.Background())
	if err != nil {
		return nil, err
	}
	records := make([]models.Record, len(b))
	for i, v := range b {
		records[i] = models.Record(v)
	}
	return records, nil
}

func (repo *nationalityRepo) Create(ctx context.Context, nat models.Nationality) (models.Nationality, error) {
	rep := repo.s.sqlClientByCtx(ctx)
	q := `INSERT INTO nationalities 
			(code,title,activated_at,deactivated_at) 
			VALUES ($1,$2,$3,$4)`
	row := newDbNationality(nat)
	_, err := rep.Exec(q,
		row.Code, row.Title, row.ActivatedAt, row.DeactivatedAt,
	)
	if err != nil {
		return nat, err
	}
	return row.toModel(), nil
}

func (repo *nationalityRepo) Read(ctx context.Context, code string) (models.Nationality, error) {
	rep := repo.s.sqlClientByCtx(ctx)
	q := `SELECT * FROM nationalities WHERE code=$1`
	row := dbNationality{}
	if err := rep.Get(&row, q, code); err != nil {
		return models.Nationality{}, err
	}
	return row.toModel(), nil
}

func (repo *nationalityRepo) ReadAll(ctx context.Context) ([]models.Nationality, error) {
	rep := repo.s.sqlClientByCtx(ctx)
	rows := make([]dbNationality, 0)
	q := `SELECT * FROM nationalities`
	if err := rep.Select(&rows, q); err != nil {
		return nil, err
	}
	nationalities := make([]models.Nationality, len(rows))
	for i, row := range rows {
		nationalities[i] = row.toModel()
	}
	return nationalities, nil
}

func (repo *nationalityRepo) Update(ctx context.Context, nationality models.Nationality) error {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	q := `UPDATE nationalities SET 
		    title=$1, activated_at=$2,deactivated_at=$3
		 WHERE code=$4`
	row := newDbNationality(nationality)
	res, err := sqlClient.Exec(q,
		row.Title, row.ActivatedAt, row.DeactivatedAt, row.Code)
	if err != nil {
		return err
	}
	if c, _ := res.RowsAffected(); c == 0 {
		return errs.Errf(errs.ErrNotFound, "No rows affected")
	}
	return nil
}

func (repo *nationalityRepo) Delete(ctx context.Context, code string) error {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	q := `DELETE FROM nationalities WHERE code=$1`
	res, err := sqlClient.Exec(q, code)
	if err != nil {
		return err
	}
	if c, _ := res.RowsAffected(); c == 0 {
		return errs.Errf(errs.ErrNotFound, "No rows affected")
	}
	return nil
}

func (repo *nationalityRepo) ClearAll(ctx context.Context) error {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	q := `DELETE FROM nationalities WHERE TRUE`
	if _, err := sqlClient.Exec(q); err != nil {
		return err
	}
	return nil
}
