package repository

import "database/sql"

type InfoRepository struct {
	db *sql.DB
}

func NewInfoRepository(db *sql.DB) *InfoRepository {
	return &InfoRepository{db}
}

func (r *InfoRepository) GetVersion() (*string, error) {
	var version string

	row := r.db.QueryRow("SELECT version();")
	if row.Err() != nil {
		return nil, row.Err()
	}

	err := row.Scan(&version)
	if err != nil {
		return nil, err
	}

	return &version, nil
}
