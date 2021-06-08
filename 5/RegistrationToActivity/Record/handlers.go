package main

import (
	"RegistrationToActivity/Record/models"
	"RegistrationToActivity/api/auth"
	"RegistrationToActivity/api/record"
	"context"
	"encoding/json"
)

func (s *server) SetRecord(ctx context.Context, in *record.SetRecordRequest) (*record.SetRecordResponse, error) {
	var newRecord models.DbRecord

	err := json.Unmarshal([]byte(in.RecordJson), &newRecord)
	if err != nil {
		return &record.SetRecordResponse{IsSuccess: false}, err
	}

	_, err = app.Db.Insert(newRecord)

	if err == nil {
		return &record.SetRecordResponse{IsSuccess: true}, nil
	} else {
		return &record.SetRecordResponse{IsSuccess: false}, err
	}
}

func (s *server) GetRecords(ctx context.Context, in *record.TokenRequest) (*record.GetRecordsResponse, error) {
	res, err := authClient.IsAdmin(ctx, &auth.IsAdminRequest{Token: in.Token})
	if  res == nil || res.IsAdmin == false || err != nil {
		return &record.GetRecordsResponse{IsSuccess: false}, err
	}

	records, err := app.Db.GetAll()
	if err != nil {
		return &record.GetRecordsResponse{IsSuccess: false}, err
	}

	var recordsJson []string

	for _, elem := range records {
		newRecord, err := json.Marshal(elem)
		if err != nil {
			return &record.GetRecordsResponse{IsSuccess: false}, err
		}

		recordsJson = append(recordsJson, string(newRecord))
	}

	return &record.GetRecordsResponse{IsSuccess: true, RecordsJson: recordsJson}, nil
}
