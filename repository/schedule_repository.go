package repository

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uzimaru0000/aizu-garbage/model"
	"github.com/uzimaru0000/aizu-garbage/usecase/port"
)

type scheduleRepository struct {
	db *sql.DB
}

func NewScheduleRepository(db *sql.DB) port.ScheduleAccessor {
	return &scheduleRepository{db: db}
}

func (repo *scheduleRepository) GetSchedules(ctx context.Context, placeID string) ([]*model.Schedule, error) {
	conn, err := repo.db.Conn(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.QueryContext(ctx, `
		SELECT type, date, created_at from schedule where place_id = ?
	`, placeID)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	}
	defer rows.Close()

	schedules := []*model.Schedule{}
	for rows.Next() {
		schedule := &model.Schedule{}
		err = rows.Scan(
			&schedule.Type,
			&schedule.Date,
			&schedule.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		schedules = append(schedules, schedule)
	}

	if len(schedules) == 0 {
		return nil, nil
	}

	return schedules, nil
}

func (repo *scheduleRepository) SaveSchedule(ctx context.Context, placeID string, schedule *model.Schedule) error {
	conn, err := repo.db.Conn(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecContext(ctx, `
		INSERT INTO schedule(type, date, place_id) values (?, ?, ?)
	`, schedule.Type, schedule.Date, placeID)

	return err
}

func (repo *scheduleRepository) DeleteSchedules(ctx context.Context) error {
	conn, err := repo.db.Conn(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.ExecContext(ctx, `
		DELETE FROM schedule
	`)
	if err != nil {
		return err
	}

	return nil
}
