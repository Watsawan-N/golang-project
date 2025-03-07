package helper

import (
	"encoding/json"
	"fmt"
	"golang-project/pkg/errs"
	"net/http"
)

func Help() {
	fmt.Println("This is a helper function")
}

type Common struct {
	UseCase ICommon
}

func MakeICommon() ICommon {
	common := &Common{}
	common.UseCase = common
	return common
}

type ICommon interface {
	APIResponse(w *http.ResponseWriter, statusCode int, data interface{})
	HandleErr(w *http.ResponseWriter, err *errs.Error) errs.Error
	HandlePanic(r interface{}, w *http.ResponseWriter) error
}

func (c Common) APIResponse(w *http.ResponseWriter, statusCode int, data interface{}) {
	apiResponse(w, statusCode, data)
}

func apiResponse(w *http.ResponseWriter, statusCode int, data interface{}) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).WriteHeader(statusCode)
	_ = json.NewEncoder(*w).Encode(data)
}

func (c Common) HandleErr(w *http.ResponseWriter, err *errs.Error) errs.Error {
	return handleErr(w, err)
}

func handleErr(w *http.ResponseWriter, err *errs.Error) errs.Error {
	apiResponse(w, err.StatusCode, err.Err.Error())
	return *err
}

func (c Common) HandlePanic(r interface{}, w *http.ResponseWriter) error {
	if r != nil {
		errMsg := "Internal Server Exception: "
		switch x := r.(type) {
		case string:
			errMsg = errMsg + x
		case error:
			errMsg = errMsg + x.Error()
		default:
			errMsg = errMsg + "unexpected panic"
		}
		err := errs.NewInternalServerError(errMsg)
		return handleErr(w, err)
	}

	return nil
}

func GetErrorMsgFromRecover(r interface{}) string {
	errMsg := "Internal Server Exception: "

	switch x := r.(type) {
	case string:
		errMsg = errMsg + x
	case error:
		errMsg = errMsg + x.Error()
	default:
		errMsg = errMsg + "unexpected panic"
	}

	return errMsg
}
