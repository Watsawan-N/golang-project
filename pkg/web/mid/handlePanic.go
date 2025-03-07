package mid

import (
	"context"
	"golang-project/pkg/errs"
	"golang-project/pkg/helper"
	"golang-project/pkg/model"
	"golang-project/pkg/service"
	"golang-project/pkg/web"

	"net/http"
	"time"
)

func HandlePanic(common helper.ICommon, errorLogService service.IErrorLogService) web.Middleware {
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) (panicErr error) {

			requestTime := time.Now()
			body := getBody(r)

			defer func() {
				if rec := recover(); rec != nil {

					errMsg := helper.GetErrorMsgFromRecover(rec)

					err := errs.NewInternalServerError(errMsg)
					panicErr = common.HandleErr(&w, err)
					if panicErr != nil {

						ResponseTime := time.Now()
						ResponseStatus := getStatusCode(panicErr)
						ResponseBody := panicErr.Error()

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
				}
			}()

			panicErr = handler(ctx, w, r)

			return panicErr
		}
		return h
	}
	return m
}
