package controller

import (
	"apelsin/models"
	"context"
)

func (ctrl *Controller) CreateProduct(ctx context.Context, model models.Product) error {
	var sessionErr error
	session, ctx, err := ctrl.store.StartSession(ctx)
	if err != nil {
		return err
	}
	defer func() { session.Close(sessionErr) }()
	repo := ctrl.store.Product()

	if _, sessionErr = repo.Create(ctx, model); sessionErr != nil {
		return sessionErr
	}
	return nil
}

func (ctrl *Controller) GetProduct(ctx context.Context, code string) (models.Product, error) {
	repo := ctrl.store.Product()
	products, err := repo.Read(ctx, code)
	if err != nil {
		return models.Product{}, err
	}
	return products, nil
}

func (ctrl *Controller) ListProduct(ctx context.Context) ([]models.Product, error) {
	repo := ctrl.store.Product()
	products, err := repo.ReadAll(ctx)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (ctrl *Controller) UpdateProduct(ctx context.Context, model models.Product) error {
	var sessionErr error
	session, ctx, err := ctrl.store.StartSession(ctx)
	if err != nil {
		return err
	}
	defer func() { session.Close(sessionErr) }()
	repo := ctrl.store.Product()

	if sessionErr = repo.Update(ctx, model); sessionErr != nil {
		return sessionErr
	}
	return nil
}

func (ctrl *Controller) DeleteProduct(ctx context.Context, id string) error {
	session, ctx, err := ctrl.store.StartSession(ctx)
	if err != nil {
		return err
	}
	var sessionErr error
	defer func() { session.Close(sessionErr) }()
	repo := ctrl.store.Product()
	sessionErr = repo.Delete(ctx, id)
	if sessionErr != nil {
		return sessionErr
	}
	return nil
}
