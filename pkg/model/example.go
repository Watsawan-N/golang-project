package model

import "time"

type Example struct {
	StartDate *time.Time
	EndDate   *time.Time
	IsTest    *bool
	LangCode  *string
	TestName  *string
}
