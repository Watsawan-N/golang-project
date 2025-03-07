package model

import "time"

type ErrorLog struct {
	Method         *string
	RequestURI     *string
	RequestBody    *string
	RequestTime    *time.Time
	ResponseTime   *time.Time
	ResponseStatus *int
	ResponseBody   *string
}
