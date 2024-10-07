package repository

import (
	"database/sql"
)

type DBRepository struct {
	db *sql.DB
}

func NewDBRepository(db *sql.DB) *DBRepository {
	return &DBRepository{db: db}
}

func GetProcessorDB() {
	// TODO: Implement logic to load processor from database
}
