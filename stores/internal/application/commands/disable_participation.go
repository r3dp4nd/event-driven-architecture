package commands

import (
	"context"
	"github.com/Sraik25/event-driven-architecture/stores/internal/domain"
)

type (
	DisableParticipation struct {
		ID string
	}

	DisableParticipationHandler struct {
		stores domain.StoreRepository
	}
)

func NewDisableParticipationHandler(stores domain.StoreRepository) DisableParticipationHandler {
	return DisableParticipationHandler{stores: stores}
}

func (h DisableParticipationHandler) DisableParticipation(ctx context.Context, cmd DisableParticipation) error {
	store, err := h.stores.Find(ctx, cmd.ID)
	if err != nil {
		return err
	}

	err = store.DisableParticipation()
	if err != nil {
		return err
	}

	return h.stores.Update(ctx, store)
}
