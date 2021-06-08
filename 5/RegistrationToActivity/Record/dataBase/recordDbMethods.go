package dataBase

import (
	"RegistrationToActivity/Record/models"
	"database/sql"
)

type RecordModel struct {
	Db *sql.DB
}

func (model *RecordModel) Insert(dbRecord models.DbRecord) (sql.Result, error) {
	query := "INSERT INTO records VALUES (?, ?, ?, ?)"

	res, err := model.Db.Exec(query, dbRecord.Name, dbRecord.LastName, dbRecord.Email, dbRecord.Birthday)
	return res, err
}

func (model *RecordModel) GetAll() ([]models.DbRecord, error) {
	var records []models.DbRecord

	query := "SELECT * FROM records ORDER BY birthday"

	rows, err := model.Db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		result := models.DbRecord{}
		err = rows.Scan(&result.Name, &result.LastName, &result.Email, &result.Birthday)
		if err != nil {
			return nil, err
		}
		records = append(records, result)
	}

	return records, err
}