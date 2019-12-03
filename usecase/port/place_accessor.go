package port

import (
	"context"

	"github.com/uzimaru0000/aizu-garbage/model"
)

type PlaceAccessor interface {
	GetPlaces(ctx context.Context) ([]*model.Place, error)
	GetPlace(ctx context.Context, id string) (*model.Place, error)
}
