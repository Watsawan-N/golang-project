package endpoint

import (
	"context"
	"golang-project/pkg/helper"
	"golang-project/pkg/service"
	"net/http"
)

type ExampleEndpoint struct {
	ExampleService service.IExampleService
	Common         helper.ICommon
}

func MakeExampleEndpoint(service *service.IExampleService, common *helper.ICommon) *ExampleEndpoint {
	return &ExampleEndpoint{
		ExampleService: *service,
		Common:         *common,
	}
}

func (e *ExampleEndpoint) GetById(_ context.Context, w http.ResponseWriter, _ *http.Request) error {
	res, err := e.ExampleService.GetById(1)
	if err != nil {
		return e.Common.HandleErr(&w, err)
	}

	e.Common.APIResponse(&w, http.StatusOK, res)

	return nil
}
