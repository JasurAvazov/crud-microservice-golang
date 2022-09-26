package controller

import (
	"apelsin/models"
	"context"
)

func (ctrl *Controller) CreatePayment(ctx context.Context, model models.Payment) error {
	var sessionErr error
	session, ctx, err := ctrl.store.StartSession(ctx)
	if err != nil {
		return err
	}
	defer func() { session.Close(sessionErr) }()
	repo := ctrl.store.Payment()

	if _, sessionErr = repo.Create(ctx, model); sessionErr != nil {
		return sessionErr
	}
	return nil
}

func (ctrl *Controller) GetPayment(ctx context.Context, code string) (models.Payment, error) {
	repo := ctrl.store.Payment()
	payments, err := repo.Read(ctx, code)
	if err != nil {
		return models.Payment{}, err
	}
	return payments, nil
}

func (ctrl *Controller) ListPayment(ctx context.Context) ([]models.Payment, error) {
	repo := ctrl.store.Payment()
	payments, err := repo.ReadAll(ctx)
	if err != nil {
		return nil, err
	}
	return payments, nil
}

func (ctrl *Controller) UpdatePayment(ctx context.Context, model models.Payment) error {
	var sessionErr error
	session, ctx, err := ctrl.store.StartSession(ctx)
	if err != nil {
		return err
	}
	defer func() { session.Close(sessionErr) }()
	repo := ctrl.store.Payment()

	if sessionErr = repo.Update(ctx, model); sessionErr != nil {
		return sessionErr
	}
	return nil
}

func (ctrl *Controller) DeletePayment(ctx context.Context, id string) error {
	session, ctx, err := ctrl.store.StartSession(ctx)
	if err != nil {
		return err
	}
	var sessionErr error
	defer func() { session.Close(sessionErr) }()
	repo := ctrl.store.Payment()
	sessionErr = repo.Delete(ctx, id)
	if sessionErr != nil {
		return sessionErr
	}
	return nil
}
