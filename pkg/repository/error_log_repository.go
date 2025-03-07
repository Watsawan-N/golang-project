package repository

import (
	"golang-project/pkg/entity"
	"strconv"

	"gorm.io/gorm"
)

type IErrorLogRepository interface {
	Create(input entity.ErrorLog) error
	DeleteOverDueDate(day int) error
}

type ErrorLogRepository struct {
	DB *gorm.DB
}

func MakeIErrorLogRepository(db *gorm.DB) IErrorLogRepository {
	return &ErrorLogRepository{
		DB: db,
	}
}

func (r ErrorLogRepository) Create(input entity.ErrorLog) error {
	tx := r.DB.Create(&input)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (r ErrorLogRepository) DeleteOverDueDate(day int) error {
	sqlWhere := "DATEDIFF(day, created_at, SYSDATETIMEOFFSET()) > " + strconv.Itoa(day)
	tx := r.DB.Unscoped().Where(sqlWhere).Delete(&entity.ErrorLog{})
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
