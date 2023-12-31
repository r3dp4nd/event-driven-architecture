package commands

import (
	"context"
	"github.com/Sraik25/event-driven-architecture/stores/internal/domain"
)

type (
	EnableParticipation struct {
		ID string
	}

	EnableParticipationHandler struct {
		stores domain.StoreRepository
	}
)

func NewEnableParticipationHandler(stores domain.StoreRepository) EnableParticipationHandler {
	return EnableParticipationHandler{
		stores: stores,
	}
}

func (h EnableParticipationHandler) EnableParticipation(ctx context.Context, cmd EnableParticipation) error {
	store, err := h.stores.Load(ctx, cmd.ID)
	if err != nil {
		return err
	}

	if err = store.EnableParticipation(); err != nil {
		return err
	}

	return h.stores.Save(ctx, store)
}
