package mid

import (
	"bytes"
	"context"
	"golang-project/pkg/errs"
	"golang-project/pkg/model"
	"golang-project/pkg/service"
	"golang-project/pkg/web"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Logger(errorLogService service.IErrorLogService) web.Middleware {

	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			requestTime := time.Now()
			body := getBody(r)

			err := handler(ctx, w, r)

			if err != nil {
				ResponseTime := time.Now()
				ResponseStatus := getStatusCode(err)
				ResponseBody := err.Error()

				input := model.ErrorLog{
					Method:         &r.Method,
					RequestURI:     &r.RequestURI,
					RequestBody:    &body,
					RequestTime:    &requestTime,
					ResponseTime:   &ResponseTime,
					ResponseStatus: &ResponseStatus,
					ResponseBody:   &ResponseBody,
				}
				saveLog(errorLogService, input)
			}

			return err
		}
		return h
	}
	return m
}

func getBody(r *http.Request) string {
	if r.Body == nil {
		return ""
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return ""
	}

	r.Body = ioutil.NopCloser(bytes.NewReader(body))

	return string(body)
}

func getStatusCode(err error) int {
	switch e := err.(type) {
	case errs.Error:
		return e.StatusCode
	default:
		return 0
	}
}

func saveLog(errorLogService service.IErrorLogService, input model.ErrorLog) {
	err := errorLogService.Create(input)
	if err != nil {
		log.Println("Save error log failed:", err.Err.Error())
	}
}
