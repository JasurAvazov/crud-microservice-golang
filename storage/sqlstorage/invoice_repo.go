package sqlstorage

import (
	"apelsin/errs"
	"apelsin/models"
	"context"
)

type invoiceRepo struct {
	s *Store
}

func (repo *invoiceRepo) GetRecords(ctx context.Context) ([]models.Record, error) {
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

func (repo *invoiceRepo) Create(ctx context.Context, invoice models.Invoice) (models.Invoice, error) {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	q := `INSERT INTO invoice (id,ord_id,amount,issued,due) 
			VALUES ($1,$2,$3,$4,$5)`
	row := newDbInvoice(invoice)
	_, err := sqlClient.Exec(q,
		row.Id, row.OrdId, row.Amount, row.Issued, row.Due,
	)
	if err != nil {
		return invoice, err
	}
	return invoice, nil
}

func (repo *invoiceRepo) Read(ctx context.Context, code string) (models.Invoice, error) {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	q := `SELECT * FROM invoice WHERE id=$1`
	row := dbInvoice{}
	if err := sqlClient.Get(&row, q, code); err != nil {
		return models.Invoice{}, err
	}
	return row.toModel(), nil
}

func (repo *invoiceRepo) ReadAll(ctx context.Context) ([]models.Invoice, error) {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	rows := make([]dbInvoice, 0)
	q := `SELECT * FROM invoice`
	if err := sqlClient.Select(&rows, q); err != nil {
		return nil, err
	}
	invoices := make([]models.Invoice, len(rows))
	for i, row := range rows {
		invoices[i] = row.toModel()
	}
	return invoices, nil
}

func (repo *invoiceRepo) Update(ctx context.Context, invoice models.Invoice) error {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	q := `UPDATE invoice SET 
		    ord_id=$2,amount=$3,issued=$4,due=$5
		 WHERE id=$1`
	row := newDbInvoice(invoice)
	res, err := sqlClient.Exec(q,
		row.Id, row.OrdId, row.Amount, row.Issued, row.Due)
	if err != nil {
		return err
	}
	if c, _ := res.RowsAffected(); c == 0 {
		return errs.Errf(errs.ErrNotFound, "No rows affected")
	}
	return nil
}

func (repo *invoiceRepo) Delete(ctx context.Context, id string) error {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	q := `DELETE FROM invoice WHERE id=$1`
	res, err := sqlClient.Exec(q, id)
	if err != nil {
		return err
	}
	if c, _ := res.RowsAffected(); c == 0 {
		return errs.Errf(errs.ErrNotFound, "No rows affected")
	}
	return nil
}

func (repo *invoiceRepo) ClearAll(ctx context.Context) error {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	q := `DELETE FROM invoice WHERE TRUE`
	if _, err := sqlClient.Exec(q); err != nil {
		return err
	}
	return nil
}
