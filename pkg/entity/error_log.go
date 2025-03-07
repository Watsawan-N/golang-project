package entity

import (
	"time"

	"gorm.io/gorm"
)

type ErrorLog struct {
	gorm.Model
	Method         *string
	RequestURI     *string
	RequestBody    *string
	RequestTime    *time.Time
	ResponseTime   *time.Time
	ResponseStatus *int
	ResponseBody   *string
}

func (ErrorLog) TableName() string {
	return "system_error_logs"
}
