package sqlstorage

import (
	"apelsin/errs"
	"apelsin/models"
	"context"
)

type detailRepo struct {
	s *Store
}

func (repo *detailRepo) GetRecords(ctx context.Context) ([]models.Record, error) {
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

func (repo *detailRepo) Create(ctx context.Context, detail models.Detail) (models.Detail, error) {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	q := `INSERT INTO detail (id,ord_id,pr_id,quantity) 
			VALUES ($1,$2,$3,$4)`
	row := newDbDetail(detail)
	_, err := sqlClient.Exec(q,
		row.Id, row.OrdId, row.PrId, row.Quantity,
	)
	if err != nil {
		return detail, err
	}
	return detail, nil
}

func (repo *detailRepo) Read(ctx context.Context, code string) (models.Detail, error) {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	q := `SELECT * FROM detail WHERE id=$1`
	row := dbDetail{}
	if err := sqlClient.Get(&row, q, code); err != nil {
		return models.Detail{}, err
	}
	return row.toModel(), nil
}

func (repo *detailRepo) ReadAll(ctx context.Context) ([]models.Detail, error) {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	rows := make([]dbDetail, 0)
	q := `SELECT * FROM detail`
	if err := sqlClient.Select(&rows, q); err != nil {
		return nil, err
	}
	categories := make([]models.Detail, len(rows))
	for i, row := range rows {
		categories[i] = row.toModel()
	}
	return categories, nil
}

func (repo *detailRepo) Update(ctx context.Context, detail models.Detail) error {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	q := `UPDATE detail SET 
		    ord_id=$2,pr_id=$3,quantity=$4
		 WHERE id=$1`
	row := newDbDetail(detail)
	res, err := sqlClient.Exec(q,
		row.Id, row.OrdId, row.PrId, row.Quantity)
	if err != nil {
		return err
	}
	if c, _ := res.RowsAffected(); c == 0 {
		return errs.Errf(errs.ErrNotFound, "No rows affected")
	}
	return nil
}

func (repo *detailRepo) Delete(ctx context.Context, id string) error {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	q := `DELETE FROM detail WHERE id=$1`
	res, err := sqlClient.Exec(q, id)
	if err != nil {
		return err
	}
	if c, _ := res.RowsAffected(); c == 0 {
		return errs.Errf(errs.ErrNotFound, "No rows affected")
	}
	return nil
}

func (repo *detailRepo) ClearAll(ctx context.Context) error {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	q := `DELETE FROM detail WHERE TRUE`
	if _, err := sqlClient.Exec(q); err != nil {
		return err
	}
	return nil
}
