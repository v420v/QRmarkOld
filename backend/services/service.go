package services

import "database/sql"

type QrmarkAPIService struct {
	DB *sql.DB
}

func NewQrmarkAPIService(db *sql.DB) *QrmarkAPIService {
	return &QrmarkAPIService{DB: db}
}
