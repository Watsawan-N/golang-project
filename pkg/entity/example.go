package entity

import (
	"time"
)

type Example struct {
	DefaultField
	StartDate  *time.Time
	EndDate    *time.Time
	IsTest     *bool
	TestString *string
	TestBool   *bool
	TestDate   *time.Time
}
