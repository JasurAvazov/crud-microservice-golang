package sqlstorage

import (
	"apelsin/errs"
	"apelsin/models"
	"context"
	"database/sql"
	"errors"
)

type offerRepo struct {
	s *Store
}

func (repo *offerRepo) GetRecords(ctx context.Context) ([]models.Record, error) {
	c, err := repo.ReadAll(ctx)
	if err != nil {
		return nil, err
	}
	records := make([]models.Record, len(c))
	for i, v := range c {
		records[i] = models.Record(v)
	}
	return records, nil
}

func (repo *offerRepo) Create(ctx context.Context, offer models.Offer) (models.Offer, error) {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	query := `INSERT INTO offers
    (code, name, should_Be_Signed, signature_type)
    VALUES ( $1, $2, $3, $4)`
	row := newDbOffer(offer)
	_, err := sqlClient.Exec(query,
		row.Code, row.Name, row.ShouldBeSigned, row.SignatureType,
	)
	if err != nil {
		return offer, errs.Errf(errs.ErrWrongInput, err.Error())
	}
	return row.toModel(), nil
}

func (repo *offerRepo) Read(ctx context.Context, code string) (models.Offer, error) {
	rep := repo.s.sqlClientByCtx(ctx)
	query := `SELECT * FROM offers WHERE code = $1`

	var model dbOffer

	if err := rep.QueryRowx(query, code).StructScan(&model); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Offer{}, errs.Errf(errs.ErrNotFound, err.Error())
		}
		return models.Offer{}, err
	}

	return model.toModel(), nil
}

func (repo *offerRepo) ReadAll(ctx context.Context) ([]models.Offer, error) {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	rows := make([]dbOffer, 0)
	q := `SELECT * FROM offers`
	if err := sqlClient.Select(&rows, q); err != nil {
		return nil, err
	}
	offers := make([]models.Offer, len(rows))
	for i := range rows {
		offers[i] = rows[i].toModel()
	}
	return offers, nil
}

func (repo *offerRepo) Update(ctx context.Context, offer models.Offer) error {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	q := `UPDATE offers
		SET name = $1 , should_Be_Signed = $2 , signature_type = $3
		WHERE code = $4`
	row := newDbOffer(offer)
	res, err := sqlClient.Exec(q,
		row.Name, row.ShouldBeSigned, row.SignatureType, row.Code)
	if err != nil {
		return err
	}
	if c, _ := res.RowsAffected(); c == 0 {
		return errs.Errf(errs.ErrNotFound, "No rows affected")
	}
	return nil
}

func (repo *offerRepo) Delete(ctx context.Context, code string) error {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	q := `DELETE FROM offers WHERE code=$1`
	res, err := sqlClient.Exec(q, code)
	if err != nil {
		return err
	}
	if c, _ := res.RowsAffected(); c == 0 {
		return errs.Errf(errs.ErrNotFound, "No rows affected")
	}
	return nil
}

func (repo *offerRepo) GetFullOffer(ctx context.Context, code, lang string) (models.FullOffer, error) {
	sqlClient := repo.s.sqlClientByCtx(ctx)
	row := dbFullOffer{}
	query := `SELECT 
		t1.code as code, 
		t2.id as version_code,
		t1.should_be_signed as should_be_signed, 
		t1.signature_type as signature_type, 
		t3.name as name,
		t3.content as content
	FROM offers t1
	INNER JOIN offer_versions t2 ON t1.code = t2.offer_code
	LEFT JOIN offer_version_translations t3 ON t2.id = t3.offer_version_id
	WHERE
		t1.code=$1 AND t3.locale = $2 AND
		t2.activated_at IS NOT NULL AND t2.activated_at < NOW() AND 
		(t2.deactivated_at IS NULL OR t2.deactivated_at > NOW())`
	if err := sqlClient.Get(&row, query, code, lang); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.FullOffer{}, errs.Errf(errs.ErrNotFound, err.Error())
		}
		return models.FullOffer{}, errs.Errf(errs.ErrSourceInternal, err.Error())
	}
	return models.FullOffer{
		Code:           row.Code,
		VersionCode:    row.VersionCode.String,
		Name:           row.Name.String,
		Content:        row.Content.String,
		ShouldBeSigned: row.ShouldBeSigned,
		SignatureType:  row.SignatureType.String,
	}, nil
}
