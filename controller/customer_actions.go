package controller

import (
	"apelsin/models"
	"context"
)

func (ctrl *Controller) DeleteCustomer(ctx context.Context, id string) error {
	session, ctx, err := ctrl.store.StartSession(ctx)
	if err != nil {
		return err
	}
	var sessionErr error
	defer func() { session.Close(sessionErr) }()
	repo := ctrl.store.Customer()
	sessionErr = repo.Delete(ctx, id)
	if sessionErr != nil {
		return sessionErr
	}
	return nil
}

func (ctrl *Controller) GetCustomer(ctx context.Context, code string) (models.Customer, error) {
	repo := ctrl.store.Customer()
	customers, err := repo.Read(ctx, code)
	if err != nil {
		return models.Customer{}, err
	}
	return customers, nil
}

func (ctrl *Controller) CreateCustomer(ctx context.Context, model models.Customer) error {
	var sessionErr error
	session, ctx, err := ctrl.store.StartSession(ctx)
	if err != nil {
		return err
	}
	defer func() { session.Close(sessionErr) }()
	repo := ctrl.store.Customer()

	if _, sessionErr = repo.Create(ctx, model); sessionErr != nil {
		return sessionErr
	}
	return nil
}

func (ctrl *Controller) UpdateCustomer(ctx context.Context, model models.Customer) error {
	var sessionErr error
	session, ctx, err := ctrl.store.StartSession(ctx)
	if err != nil {
		return err
	}
	defer func() { session.Close(sessionErr) }()
	repo := ctrl.store.Customer()

	if sessionErr = repo.Update(ctx, model); sessionErr != nil {
		return sessionErr
	}
	return nil
}

func (ctrl *Controller) ListCustomer(ctx context.Context) ([]models.Customer, error) {
	repo := ctrl.store.Customer()
	customers, err := repo.ReadAll(ctx)
	if err != nil {
		return nil, err
	}
	return customers, nil
}
