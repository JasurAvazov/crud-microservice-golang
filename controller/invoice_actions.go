package controller

import (
	"apelsin/models"
	"context"
)

func (ctrl *Controller) CreateInvoice(ctx context.Context, model models.Invoice) error {
	var sessionErr error
	session, ctx, err := ctrl.store.StartSession(ctx)
	if err != nil {
		return err
	}
	defer func() { session.Close(sessionErr) }()
	repo := ctrl.store.Invoice()

	if _, sessionErr = repo.Create(ctx, model); sessionErr != nil {
		return sessionErr
	}
	return nil
}

func (ctrl *Controller) GetInvoice(ctx context.Context, code string) (models.Invoice, error) {
	repo := ctrl.store.Invoice()
	invoices, err := repo.Read(ctx, code)
	if err != nil {
		return models.Invoice{}, err
	}
	return invoices, nil
}

func (ctrl *Controller) ListInvoice(ctx context.Context) ([]models.Invoice, error) {
	repo := ctrl.store.Invoice()
	invoices, err := repo.ReadAll(ctx)
	if err != nil {
		return nil, err
	}
	return invoices, nil
}

func (ctrl *Controller) UpdateInvoice(ctx context.Context, model models.Invoice) error {
	var sessionErr error
	session, ctx, err := ctrl.store.StartSession(ctx)
	if err != nil {
		return err
	}
	defer func() { session.Close(sessionErr) }()
	repo := ctrl.store.Invoice()

	if sessionErr = repo.Update(ctx, model); sessionErr != nil {
		return sessionErr
	}
	return nil
}

func (ctrl *Controller) DeleteInvoice(ctx context.Context, id string) error {
	session, ctx, err := ctrl.store.StartSession(ctx)
	if err != nil {
		return err
	}
	var sessionErr error
	defer func() { session.Close(sessionErr) }()
	repo := ctrl.store.Invoice()
	sessionErr = repo.Delete(ctx, id)
	if sessionErr != nil {
		return sessionErr
	}
	return nil
}
