package repository

import (
	"context"
	"database/sql"
	"github.com/uzimaru0000/aizu-garbage/model"
	"github.com/uzimaru0000/aizu-garbage/usecase/port"
)

type placeRepository struct {
	db *sql.DB
}

func NewPlaceRepository(db *sql.DB) port.PlaceAccessor {
	return &placeRepository{db: db}
}

func (repo *placeRepository) GetPlaces(ctx context.Context) ([]*model.Place, error) {
	conn, err := repo.db.Conn(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.QueryContext(ctx, `
		SELECT * from place
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	places := []*model.Place{}
	for rows.Next() {
		place, err := mapPlace(rows)
		if err != nil {
			return nil, err
		}
		places = append(places, place)
	}

	if len(places) == 0 {
		return nil, nil
	}

	return places, nil
}

func (repo *placeRepository) GetPlace(ctx context.Context, id string) (*model.Place, error) {
	conn, err := repo.db.Conn(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rows, err := conn.QueryContext(ctx, `
		SELECT * from place WHERE id = ?
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var place *model.Place
	for rows.Next() {
		place, err = mapPlace(rows)
		if err != nil {
			return nil, err
		}
		break
	}

	return place, nil
}

func mapPlace(rows *sql.Rows) (*model.Place, error) {
	place := &model.Place{}
	err := rows.Scan(
		&place.ID,
		&place.Name,
	)

	if err != nil {
		return nil, err
	}

	return place, nil
}
