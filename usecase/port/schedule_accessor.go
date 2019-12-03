package port

import (
	"context"

	"github.com/uzimaru0000/aizu-garbage/model"
)

type ScheduleAccessor interface {
	GetSchedules(ctx context.Context, placeID string) ([]*model.Schedule, error)
	SaveSchedule(ctx context.Context, placeID string, schedule *model.Schedule) error
	DeleteSchedules(ctx context.Context) error
}

type ScheduleFetcher interface {
	FetchSchedules(placeID string) ([]*model.Schedule, error)
}
