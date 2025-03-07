package web

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Values struct {
	TraceID string
	Now     time.Time
}

type App struct {
	Middleware []Middleware
	Mux        *mux.Router
}

func GetValues(ctx context.Context) (*Values, error) {
	v, ok := ctx.Value("key").(*Values)
	if !ok {
		return nil, errors.New("web value missing from context")
	}
	return v, nil
}

type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

func NewApp(mdw ...Middleware) *App {
	return &App{
		Middleware: mdw,
		Mux:        mux.NewRouter(),
	}
}

func (app *App) Handle(method string, path string, handler Handler, mw ...Middleware) {
	handler = wrapMiddleware(app.Middleware, handler)
	handler = wrapMiddleware(mw, handler)

	h := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		v := Values{
			Now: time.Now(),
		}

		ctx = context.WithValue(ctx, "key", &v)

		if err := handler(ctx, w, r); err != nil {
			return
		}
	}
	app.Mux.Methods(method).Path(path).HandlerFunc(h)
}
