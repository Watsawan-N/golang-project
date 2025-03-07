package api

import (
	"golang-project/pkg/api/infrastructure"
	"golang-project/pkg/web"
	"golang-project/pkg/web/mid"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type APIConfig struct {
	DB      *gorm.DB
	Timeout time.Duration
}

func APIMux(config APIConfig) http.Handler {

	createHelper := infrastructure.CreateHelper()
	createRepo := infrastructure.CreateRepository(config.DB)
	createService := infrastructure.CreateService(createRepo, createHelper)
	createEndpoint := infrastructure.CreateEndpoint(createService, createHelper)

	app := web.NewApp(
		mid.HandlePanic(createHelper.ICommon, createService.IErrorLogService),
		mid.Logger(createService.IErrorLogService),
		mid.Cors("*"),
	)

	createRouting(app, createEndpoint, createService, createHelper)

	return app.Mux
}

func createRouting(
	app *web.App,
	endpoint infrastructure.Endpoint,
	service infrastructure.Service,
	helper infrastructure.Helper,
) {
	app.Mux.Methods("OPTIONS").HandlerFunc(handlerOptionAllowCrossCors)
	app.Handle("GET", "/example", endpoint.ExampleEndpoint.GetById)
}

func handlerOptionAllowCrossCors(w http.ResponseWriter, r *http.Request) {
	mid.SetHeaderAllowCors(&w, "*")
}
