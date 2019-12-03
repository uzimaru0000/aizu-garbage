package usecase

import (
	"context"

	"github.com/uzimaru0000/aizu-garbage/model"
	"github.com/uzimaru0000/aizu-garbage/usecase/port"
)

type PlaceUseCase struct {
	pa port.PlaceAccessor
}

func NewPlaceUseCase(pa port.PlaceAccessor) *PlaceUseCase {
	return &PlaceUseCase{pa: pa}
}

func (pu *PlaceUseCase) GetAll(ctx context.Context) ([]*model.Place, error) {
	return pu.pa.GetPlaces(ctx)
}

func (pu *PlaceUseCase) Get(ctx context.Context, id string) (*model.Place, error) {
	return pu.pa.GetPlace(ctx, id)
}
