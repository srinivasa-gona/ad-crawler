package repository

import (
	"database/sql"
	"log"

	"ad-crawler/model"
)

type AdRepositoryImpl struct {
	db *sql.DB
}

func NewAdRepositoryImpl(d *sql.DB) AdRepository {
	return AdRepositoryImpl{
		db: d,
	}
}

//InsertRecords - Insert records in to ad_txt table
func (ari AdRepositoryImpl) InsertRecords(publisherName string, records []model.Record) error {

	log.Printf("Insert records called with record count %d", len(records))

	statement, _ := ari.db.Prepare("BEGIN TRANSACTION")
	_, err := statement.Exec()

	statement, _ = ari.db.Prepare("DELETE FROM ads_txt where publisher = ?")
	_, err = statement.Exec(publisherName)

	if err != nil {
		log.Printf("Error in deleting ad text for publisher %s is %v", publisherName, err)
	}

	for _, record := range records {
		statement, _ := ari.db.Prepare("INSERT into ads_txt values (?, ? , ? , ?, ?)")
		_, err := statement.Exec(publisherName, record.DomainName, record.PublisherActId, record.ActType, record.CertAuthId)
		if err != nil {
			log.Printf("Error in Inserting Records %v", err)
			return err
		}

	}

	statement, _ = ari.db.Prepare("COMMIT")
	_, err = statement.Exec()
	return nil
}

//CreateTable - Creates ads_txt table if doesn't exist
func (ari AdRepositoryImpl) CreateTable() error {
	statement, err := ari.db.Prepare("CREATE TABLE IF NOT EXISTS ads_txt (publisher TEXT, domain_name TEXT, publisher_act_id TEXT,act_type TEXT,cert_auth_id TEXT)")
	if err != nil {
		log.Printf("Error in prepate statement for create table %v", err)
		return err
	}

	_, err = statement.Exec()

	if err != nil {
		log.Printf("Error in executing create table statement %v", err)
		return err
	}

	return nil

}

//GetRecords - Returns ad records for given publisher
func (ari AdRepositoryImpl) GetRecords(publisher string) ([]model.Record, error) {
	recordList := []model.Record{}
	query := "SELECT domain_name,publisher_act_id,act_type,cert_auth_id FROM ads_txt WHERE publisher = '" + publisher + "'"
	log.Printf("Get Records Query is %s", query)
	result, err := ari.db.Query(query, publisher)
	if err != nil {
		log.Printf("Error in getting ads list by publisher : %v", err)
		return recordList, err
	}

	for result.Next() {
		var record model.Record
		err := result.Scan(&record.DomainName, &record.PublisherActId, &record.ActType, &record.CertAuthId)
		if err != nil {
			log.Printf("Error in converting ads list by publisher : %v", err)

		}

		recordList = append(recordList, record)
	}
	return recordList, nil

}
