package controller

import (
	"apelsin/models"
	"context"
)

func (ctrl *Controller) CreateCategory(ctx context.Context, model models.Category) error {
	var sessionErr error
	session, ctx, err := ctrl.store.StartSession(ctx)
	if err != nil {
		return err
	}
	defer func() { session.Close(sessionErr) }()
	repo := ctrl.store.Category()

	if _, sessionErr = repo.Create(ctx, model); sessionErr != nil {
		return sessionErr
	}
	return nil
}

func (ctrl *Controller) GetCategory(ctx context.Context, code string) (models.Category, error) {
	repo := ctrl.store.Category()
	categories, err := repo.Read(ctx, code)
	if err != nil {
		return models.Category{}, err
	}
	return categories, nil
}

func (ctrl *Controller) ListCategory(ctx context.Context) ([]models.Category, error) {
	repo := ctrl.store.Category()
	categories, err := repo.ReadAll(ctx)
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (ctrl *Controller) UpdateCategory(ctx context.Context, model models.Category) error {
	var sessionErr error
	session, ctx, err := ctrl.store.StartSession(ctx)
	if err != nil {
		return err
	}
	defer func() { session.Close(sessionErr) }()
	repo := ctrl.store.Category()

	if sessionErr = repo.Update(ctx, model); sessionErr != nil {
		return sessionErr
	}
	return nil
}

func (ctrl *Controller) DeleteCategory(ctx context.Context, id string) error {
	session, ctx, err := ctrl.store.StartSession(ctx)
	if err != nil {
		return err
	}
	var sessionErr error
	defer func() { session.Close(sessionErr) }()
	repo := ctrl.store.Category()
	sessionErr = repo.Delete(ctx, id)
	if sessionErr != nil {
		return sessionErr
	}
	return nil
}
