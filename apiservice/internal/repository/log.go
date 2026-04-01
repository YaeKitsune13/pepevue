package repository

import (
	"apiservice/internal/model"

	"gorm.io/gorm"
)

type LogRepository struct {
	db *gorm.DB
}

func NewLogRepository(db *gorm.DB) *LogRepository {
	return &LogRepository{
		db: db,
	}
}

func (r *LogRepository) CreateLog(log *model.Log) error {
	return r.db.Create(log).Error
}
