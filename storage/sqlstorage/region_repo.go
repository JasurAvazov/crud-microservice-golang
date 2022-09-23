package sqlstorage

import (
	"apelsin/errs"
	"apelsin/models"
	"context"
)

type regionRepo struct {
	s *Store
}

func (repo *regionRepo) GetRecords(ctx context.Context) ([]models.Record, error) {
	b, err := repo.ReadAll(ctx)
	if err != nil {
		return nil, errs.Errf(errs.ErrSourceInternal, err.Error())
	}
	records := make([]models.Record, len(b))
	for i, v := range b {
		records[i] = models.Record(v)
	}
	return records, nil
}

func (repo *regionRepo) Create(ctx context.Context, region models.Region) (models.Region, error) {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	q := `INSERT INTO regions 
			(code,title,rank,activated_at,deactivated_at) 
			VALUES ($1,$2,$3,$4,$5)`
	row := newDbRegion(region)
	_, err := sqlClient.Exec(q,
		row.Code, row.Title, row.Rank, row.ActivatedAt, row.DeactivatedAt,
	)
	if err != nil {
		return region, errs.Errf(errs.ErrWrongInput, err.Error())
	}
	return row.toModel(), nil
}

func (repo *regionRepo) Read(ctx context.Context, code string) (models.Region, error) {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	q := `SELECT * FROM regions WHERE code=$1`
	row := dbRegion{}
	if err := sqlClient.Get(&row, q, code); err != nil {
		return models.Region{}, err
	}
	return row.toModel(), nil
}

func (repo *regionRepo) ReadAll(ctx context.Context) ([]models.Region, error) {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	rows := make([]dbRegion, 0)
	q := `SELECT * FROM regions`
	if err := sqlClient.Select(&rows, q); err != nil {
		return nil, err
	}
	regions := make([]models.Region, len(rows))
	for i, row := range rows {
		regions[i] = row.toModel()
	}
	return regions, nil
}

func (repo *regionRepo) Update(ctx context.Context, region models.Region) error {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	q := `UPDATE regions SET 
		    title=$1,rank=$2,
		 	activated_at=$3,deactivated_at=$4
		 WHERE code=$5`
	row := newDbRegion(region)
	res, err := sqlClient.Exec(q,
		row.Title, row.Rank, row.ActivatedAt, row.DeactivatedAt, row.Code)
	if err != nil {
		return err
	}
	if c, _ := res.RowsAffected(); c == 0 {
		return errs.Errf(errs.ErrNotFound, "No rows affected")
	}
	return nil
}

func (repo *regionRepo) Delete(ctx context.Context, code string) error {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	q := `DELETE FROM regions WHERE code=$1`
	res, err := sqlClient.Exec(q, code)
	if err != nil {
		return err
	}
	if c, _ := res.RowsAffected(); c == 0 {
		return errs.Errf(errs.ErrNotFound, "No rows affected")
	}
	return nil
}

func (repo *regionRepo) ClearAll(ctx context.Context) error {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	q := `DELETE FROM regions WHERE TRUE`
	if _, err := sqlClient.Exec(q); err != nil {
		return err
	}
	return nil
}
