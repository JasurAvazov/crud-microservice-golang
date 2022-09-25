package sqlstorage

import (
	"apelsin/errs"
	"apelsin/models"
	"context"
)

type categoryRepo struct {
	s *Store
}

func (repo *categoryRepo) GetRecords(ctx context.Context) ([]models.Record, error) {
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

func (repo *categoryRepo) Create(ctx context.Context, category models.Category) (models.Category, error) {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	q := `INSERT INTO category (id,name) 
			VALUES ($1,$2)`
	row := newDbCategory(category)
	_, err := sqlClient.Exec(q,
		row.Id, row.Name,
	)
	if err != nil {
		return category, err
	}
	return category, nil
}

func (repo *categoryRepo) Read(ctx context.Context, code string) (models.Category, error) {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	q := `SELECT * FROM category WHERE id=$1`
	row := dbCategory{}
	if err := sqlClient.Get(&row, q, code); err != nil {
		return models.Category{}, err
	}
	return row.toModel(), nil
}

func (repo *categoryRepo) ReadAll(ctx context.Context) ([]models.Category, error) {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	rows := make([]dbCategory, 0)
	q := `SELECT * FROM category`
	if err := sqlClient.Select(&rows, q); err != nil {
		return nil, err
	}
	categories := make([]models.Category, len(rows))
	for i, row := range rows {
		categories[i] = row.toModel()
	}
	return categories, nil
}

func (repo *categoryRepo) Update(ctx context.Context, category models.Category) error {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	q := `UPDATE category SET 
		    name=$2
		 WHERE id=$1`
	row := newDbCategory(category)
	res, err := sqlClient.Exec(q,
		row.Id, row.Name)
	if err != nil {
		return err
	}
	if c, _ := res.RowsAffected(); c == 0 {
		return errs.Errf(errs.ErrNotFound, "No rows affected")
	}
	return nil
}

func (repo *categoryRepo) Delete(ctx context.Context, id string) error {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	q := `DELETE FROM category WHERE id=$1`
	res, err := sqlClient.Exec(q, id)
	if err != nil {
		return err
	}
	if c, _ := res.RowsAffected(); c == 0 {
		return errs.Errf(errs.ErrNotFound, "No rows affected")
	}
	return nil
}

func (repo *categoryRepo) ClearAll(ctx context.Context) error {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	q := `DELETE FROM category WHERE TRUE`
	if _, err := sqlClient.Exec(q); err != nil {
		return err
	}
	return nil
}
