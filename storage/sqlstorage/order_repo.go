package sqlstorage

import (
	"apelsin/errs"
	"apelsin/models"
	"context"
)

type orderRepo struct {
	s *Store
}

func (repo *orderRepo) GetRecords(ctx context.Context) ([]models.Record, error) {
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

func (repo *orderRepo) Create(ctx context.Context, order models.Order) (models.Order, error) {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	q := `INSERT INTO "order" (id,date,cust_id) 
			VALUES ($1,$2,$3)`
	row := newDbOrder(order)
	_, err := sqlClient.Exec(q,
		row.Id, row.Date,
		row.Cust_id,
	)
	if err != nil {
		return order, err
	}
	return order, nil
}

func (repo *orderRepo) Read(ctx context.Context, code string) (models.Order, error) {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	q := `SELECT * FROM "order" WHERE id=$1`
	row := dbOrder{}
	if err := sqlClient.Get(&row, q, code); err != nil {
		return models.Order{}, err
	}
	return row.toModel(), nil
}

func (repo *orderRepo) ReadAll(ctx context.Context) ([]models.Order, error) {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	rows := make([]dbOrder, 0)
	q := `SELECT * FROM "order"`
	if err := sqlClient.Select(&rows, q); err != nil {
		return nil, err
	}
	orders := make([]models.Order, len(rows))
	for i, row := range rows {
		orders[i] = row.toModel()
	}
	return orders, nil
}

func (repo *orderRepo) Update(ctx context.Context, order models.Order) error {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	q := `UPDATE "order" SET 
		    date=$2, cust_id=$3
		 WHERE id=$1`
	row := newDbOrder(order)
	res, err := sqlClient.Exec(q,
		row.Id, row.Date, row.Cust_id)
	if err != nil {
		return err
	}
	if c, _ := res.RowsAffected(); c == 0 {
		return errs.Errf(errs.ErrNotFound, "No rows affected")
	}
	return nil
}

func (repo *orderRepo) Delete(ctx context.Context, id string) error {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	q := `DELETE FROM "order" WHERE id=$1`
	res, err := sqlClient.Exec(q, id)
	if err != nil {
		return err
	}
	if c, _ := res.RowsAffected(); c == 0 {
		return errs.Errf(errs.ErrNotFound, "No rows affected")
	}
	return nil
}

func (repo *orderRepo) ClearAll(ctx context.Context) error {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	q := `DELETE FROM "order" WHERE TRUE`
	if _, err := sqlClient.Exec(q); err != nil {
		return err
	}
	return nil
}
