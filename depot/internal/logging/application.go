package logging

import (
	"context"

	"github.com/Sraik25/event-driven-architecture/depot/internal/application"
	"github.com/Sraik25/event-driven-architecture/depot/internal/application/commands"
	"github.com/Sraik25/event-driven-architecture/depot/internal/application/queries"
	"github.com/Sraik25/event-driven-architecture/depot/internal/domain"
	"github.com/rs/zerolog"
)

type Application struct {
	application.App
	logger zerolog.Logger
}

var _ application.App = (*Application)(nil)

func LogApplicationAccess(application application.App, logger zerolog.Logger) Application {
	return Application{
		App:    application,
		logger: logger,
	}
}

func (a Application) CreateShoppingList(ctx context.Context, cmd commands.CreateShoppingList) (err error) {
	a.logger.Info().Msg("--> Depot.CreateShoppingList")
	defer func() { a.logger.Info().Err(err).Msg("<-- Depot.CreateShoppingList") }()
	return a.App.CreateShoppingList(ctx, cmd)
}

func (a Application) CancelShoppingList(ctx context.Context, cmd commands.CancelShoppingList) (err error) {
	a.logger.Info().Msg("--> Depot.CancelShoppingList")
	defer func() { a.logger.Info().Err(err).Msg("<-- Depot.CancelShoppingList") }()
	return a.App.CancelShoppingList(ctx, cmd)
}

func (a Application) AssignShoppingList(ctx context.Context, cmd commands.AssignShoppingList) (err error) {
	a.logger.Info().Msg("--> Depot.AssignShoppingList")
	defer func() { a.logger.Info().Err(err).Msg("<-- Depot.AssignShoppingList") }()
	return a.App.AssignShoppingList(ctx, cmd)
}

func (a Application) CompleteShoppingList(ctx context.Context, cmd commands.CompleteShoppingList) (err error) {
	a.logger.Info().Msg("--> Depot.CompleteShoppingList")
	defer func() { a.logger.Info().Err(err).Msg("<-- Depot.CompleteShoppingList") }()
	return a.App.CompleteShoppingList(ctx, cmd)
}

func (a Application) GetShoppingList(ctx context.Context, query queries.GetShoppingList) (list *domain.ShoppingList,
	err error,
) {
	a.logger.Info().Msg("--> Depot.GetShoppingList")
	defer func() { a.logger.Info().Err(err).Msg("<-- Depot.GetShoppingList") }()
	return a.App.GetShoppingList(ctx, query)
}
