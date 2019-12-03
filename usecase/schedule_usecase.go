package usecase

import (
	"context"

	"github.com/uzimaru0000/aizu-garbage/model"
	"github.com/uzimaru0000/aizu-garbage/usecase/port"
)

type ScheduleUseCase struct {
	sa port.ScheduleAccessor
	sf port.ScheduleFetcher
}

func NewScheduleUseCase(sa port.ScheduleAccessor, sf port.ScheduleFetcher) *ScheduleUseCase {
	return &ScheduleUseCase{
		sa: sa,
		sf: sf,
	}
}

func (su *ScheduleUseCase) Get(ctx context.Context, place *model.Place) ([]*model.Schedule, error) {
	schedules, err := su.sa.GetSchedules(ctx, place.ID)
	if err != nil {
		return nil, err
	}

	if schedules == nil && err == nil {
		schedules, err = su.sf.FetchSchedules(place.ID)
		if err != nil {
			return nil, err
		}

		err = su.Save(ctx, place, schedules)
		if err != nil {
			return nil, err
		}
	}

	return schedules, nil
}

func (su *ScheduleUseCase) Save(ctx context.Context, place *model.Place, schedules []*model.Schedule) error {
	for _, schedule := range schedules {
		err := su.sa.SaveSchedule(ctx, place.ID, schedule)
		if err != nil {
			return err
		}
	}

	return nil
}

func (su *ScheduleUseCase) DeleteAll(ctx context.Context) error {
	return su.sa.DeleteSchedules(ctx)
}
