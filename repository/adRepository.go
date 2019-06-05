package repository

import (
	"ad-crawler/model"
	"database/sql"
)

type AdRepository interface {
	InsertRecords(string, []model.Record) error
	GetRecords(publisher string) ([]model.Record, error)
	CreateTable() error
}

func GetConnection(config model.Configuration) (*sql.DB, error) {
	return sql.Open("sqlite3", config.DBFileLocation)

}
