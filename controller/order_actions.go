package controller

import (
	"apelsin/models"
	"context"
)

func (ctrl *Controller) CreateOrder(ctx context.Context, model models.Order) error {
	var sessionErr error
	session, ctx, err := ctrl.store.StartSession(ctx)
	if err != nil {
		return err
	}
	defer func() { session.Close(sessionErr) }()
	repo := ctrl.store.Order()

	if _, sessionErr = repo.Create(ctx, model); sessionErr != nil {
		return sessionErr
	}
	return nil
}

func (ctrl *Controller) GetOrder(ctx context.Context, code string) (models.Order, error) {
	repo := ctrl.store.Order()
	orders, err := repo.Read(ctx, code)
	if err != nil {
		return models.Order{}, err
	}
	return orders, nil
}

func (ctrl *Controller) ListOrder(ctx context.Context) ([]models.Order, error) {
	repo := ctrl.store.Order()
	orders, err := repo.ReadAll(ctx)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (ctrl *Controller) UpdateOrder(ctx context.Context, model models.Order) error {
	var sessionErr error
	session, ctx, err := ctrl.store.StartSession(ctx)
	if err != nil {
		return err
	}
	defer func() { session.Close(sessionErr) }()
	repo := ctrl.store.Order()

	if sessionErr = repo.Update(ctx, model); sessionErr != nil {
		return sessionErr
	}
	return nil
}

func (ctrl *Controller) DeleteOrder(ctx context.Context, id string) error {
	session, ctx, err := ctrl.store.StartSession(ctx)
	if err != nil {
		return err
	}
	var sessionErr error
	defer func() { session.Close(sessionErr) }()
	repo := ctrl.store.Order()
	sessionErr = repo.Delete(ctx, id)
	if sessionErr != nil {
		return sessionErr
	}
	return nil
}
