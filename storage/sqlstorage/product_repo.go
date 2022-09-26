package sqlstorage

import (
	"apelsin/errs"
	"apelsin/models"
	"context"
)

type productRepo struct {
	s *Store
}

func (repo *productRepo) GetRecords(ctx context.Context) ([]models.Record, error) {
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

func (repo *productRepo) Create(ctx context.Context, product models.Product) (models.Product, error) {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	q := `INSERT INTO product (id,name,category_id,description,
                       price,photo) 
			VALUES ($1,$2,$3,$4,$5,$6)`
	row := newDbProduct(product)
	_, err := sqlClient.Exec(q,
		row.Id, row.Name, row.CategoryId, row.Description, row.Price, row.Photo,
	)
	if err != nil {
		return product, err
	}
	return product, nil
}

func (repo *productRepo) Read(ctx context.Context, code string) (models.Product, error) {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	q := `SELECT * FROM product WHERE id=$1`
	row := dbProduct{}
	if err := sqlClient.Get(&row, q, code); err != nil {
		return models.Product{}, err
	}
	return row.toModel(), nil
}

func (repo *productRepo) ReadAll(ctx context.Context) ([]models.Product, error) {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	rows := make([]dbProduct, 0)
	q := `SELECT * FROM product`
	if err := sqlClient.Select(&rows, q); err != nil {
		return nil, err
	}
	products := make([]models.Product, len(rows))
	for i, row := range rows {
		products[i] = row.toModel()
	}
	return products, nil
}

func (repo *productRepo) Update(ctx context.Context, product models.Product) error {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	q := `UPDATE product SET 
		    name=$2,category_id=$3,
			description=$4,price=$5,photo=$6
		 WHERE id=$1`
	row := newDbProduct(product)
	res, err := sqlClient.Exec(q,
		row.Id, row.Name, row.CategoryId, row.Description, row.Price, row.Photo)
	if err != nil {
		return err
	}
	if c, _ := res.RowsAffected(); c == 0 {
		return errs.Errf(errs.ErrNotFound, "No rows affected")
	}
	return nil
}

func (repo *productRepo) Delete(ctx context.Context, id string) error {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	q := `DELETE FROM product WHERE id=$1`
	res, err := sqlClient.Exec(q, id)
	if err != nil {
		return err
	}
	if c, _ := res.RowsAffected(); c == 0 {
		return errs.Errf(errs.ErrNotFound, "No rows affected")
	}
	return nil
}

func (repo *productRepo) ClearAll(ctx context.Context) error {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	q := `DELETE FROM product WHERE TRUE`
	if _, err := sqlClient.Exec(q); err != nil {
		return err
	}
	return nil
}
