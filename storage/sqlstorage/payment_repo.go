package sqlstorage

import (
	"apelsin/errs"
	"apelsin/models"
	"context"
)

type paymentRepo struct {
	s *Store
}

func (repo *paymentRepo) GetRecords(ctx context.Context) ([]models.Record, error) {
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

func (repo *paymentRepo) Create(ctx context.Context, payment models.Payment) (models.Payment, error) {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	q := `INSERT INTO payment (id,time,amount,inv_id) 
			VALUES ($1,$2,$3,$4)`
	row := newDbPayment(payment)
	_, err := sqlClient.Exec(q,
		row.Id, row.Time, row.Amount, row.InvId,
	)
	if err != nil {
		return payment, err
	}
	return payment, nil
}

func (repo *paymentRepo) Read(ctx context.Context, code string) (models.Payment, error) {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	q := `SELECT * FROM payment WHERE id=$1`
	row := dbPayment{}
	if err := sqlClient.Get(&row, q, code); err != nil {
		return models.Payment{}, err
	}
	return row.toModel(), nil
}

func (repo *paymentRepo) ReadAll(ctx context.Context) ([]models.Payment, error) {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	rows := make([]dbPayment, 0)
	q := `SELECT * FROM payment`
	if err := sqlClient.Select(&rows, q); err != nil {
		return nil, err
	}
	payments := make([]models.Payment, len(rows))
	for i, row := range rows {
		payments[i] = row.toModel()
	}
	return payments, nil
}

func (repo *paymentRepo) Update(ctx context.Context, payment models.Payment) error {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	q := `UPDATE payment SET 
		    time=$2,amount=$3,inv_id=$4
		 WHERE id=$1`
	row := newDbPayment(payment)
	res, err := sqlClient.Exec(q,
		row.Id, row.Time, row.Amount, row.InvId)
	if err != nil {
		return err
	}
	if c, _ := res.RowsAffected(); c == 0 {
		return errs.Errf(errs.ErrNotFound, "No rows affected")
	}
	return nil
}

func (repo *paymentRepo) Delete(ctx context.Context, id string) error {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	q := `DELETE FROM payment WHERE id=$1`
	res, err := sqlClient.Exec(q, id)
	if err != nil {
		return err
	}
	if c, _ := res.RowsAffected(); c == 0 {
		return errs.Errf(errs.ErrNotFound, "No rows affected")
	}
	return nil
}

func (repo *paymentRepo) ClearAll(ctx context.Context) error {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	q := `DELETE FROM payment WHERE TRUE`
	if _, err := sqlClient.Exec(q); err != nil {
		return err
	}
	return nil
}
