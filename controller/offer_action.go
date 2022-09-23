package controller

import (
	"apelsin/models"
	"context"
)

func (ctrl *Controller) CreateOffer(ctx context.Context, model models.Offer) error {
	var sessionErr error
	session, ctx, err := ctrl.store.StartSession(ctx)
	if err != nil {
		return err
	}
	defer func() { session.Close(sessionErr) }()
	repo := ctrl.store.Offer()

	if _, sessionErr = repo.Create(ctx, model); sessionErr != nil {
		return sessionErr
	}
	return nil
}

func (ctrl *Controller) ListOffers(ctx context.Context) ([]models.Offer, error) {
	repo := ctrl.store.Offer()
	offers, err := repo.ReadAll(ctx)
	if err != nil {
		return nil, err
	}
	return offers, nil
}

func (ctrl *Controller) GetOffer(ctx context.Context, code string) (models.Offer, error) {
	repo := ctrl.store.Offer()
	offers, err := repo.Read(ctx, code)
	if err != nil {
		return models.Offer{}, err
	}
	return offers, nil
}

func (ctrl *Controller) UpdateOffer(ctx context.Context, model models.Offer) error {
	var sessionErr error
	session, ctx, err := ctrl.store.StartSession(ctx)
	if err != nil {
		return err
	}
	defer func() { session.Close(sessionErr) }()
	repo := ctrl.store.Offer()

	if sessionErr = repo.Update(ctx, model); sessionErr != nil {
		return sessionErr
	}
	return nil
}

func (ctrl *Controller) DeleteOffer(ctx context.Context, code string) error {
	session, ctx, err := ctrl.store.StartSession(ctx)
	if err != nil {
		return err
	}
	var sessionErr error
	defer func() { session.Close(sessionErr) }()
	repo := ctrl.store.Offer()
	sessionErr = repo.Delete(ctx, code)
	if sessionErr != nil {
		return sessionErr
	}
	return nil
}
