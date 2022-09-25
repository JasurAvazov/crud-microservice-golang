package sqlstorage

import (
	"apelsin/errs"
	"apelsin/models"
	"context"
)

type customerRepo struct {
	s *Store
}

func (repo *customerRepo) GetRecords(ctx context.Context) ([]models.Record, error) {
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

func (repo *customerRepo) Create(ctx context.Context, customer models.Customer) (models.Customer, error) {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	q := `INSERT INTO customer (id,name,country,address,
                       phone) 
			VALUES ($1,$2,$3,$4,$5)`
	row := newDbCustomer(customer)
	_, err := sqlClient.Exec(q,
		row.Id, row.Name,
		row.Country, row.Address,
		row.Phone,
	)
	if err != nil {
		return customer, err
	}
	return customer, nil
}

func (repo *customerRepo) Read(ctx context.Context, code string) (models.Customer, error) {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	q := `SELECT * FROM customer WHERE id=$1`
	row := dbCustomer{}
	if err := sqlClient.Get(&row, q, code); err != nil {
		return models.Customer{}, err
	}
	return row.toModel(), nil
}

func (repo *customerRepo) ReadAll(ctx context.Context) ([]models.Customer, error) {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	rows := make([]dbCustomer, 0)
	q := `SELECT * FROM customer`
	if err := sqlClient.Select(&rows, q); err != nil {
		return nil, err
	}
	customers := make([]models.Customer, len(rows))
	for i, row := range rows {
		customers[i] = row.toModel()
	}
	return customers, nil
}

func (repo *customerRepo) Update(ctx context.Context, customer models.Customer) error {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	q := `UPDATE customer SET 
		    name=$2, country=$3,
			address=$4, phone=$5
		 WHERE id=$1`
	row := newDbCustomer(customer)
	res, err := sqlClient.Exec(q,
		row.Id, row.Name, row.Country, row.Address, row.Phone)
	if err != nil {
		return err
	}
	if c, _ := res.RowsAffected(); c == 0 {
		return errs.Errf(errs.ErrNotFound, "No rows affected")
	}
	return nil
}

func (repo *customerRepo) Delete(ctx context.Context, id string) error {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	q := `DELETE FROM customer WHERE id=$1`
	res, err := sqlClient.Exec(q, id)
	if err != nil {
		return err
	}
	if c, _ := res.RowsAffected(); c == 0 {
		return errs.Errf(errs.ErrNotFound, "No rows affected")
	}
	return nil
}

func (repo *customerRepo) ClearAll(ctx context.Context) error {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	q := `DELETE FROM customer WHERE TRUE`
	if _, err := sqlClient.Exec(q); err != nil {
		return err
	}
	return nil
}
