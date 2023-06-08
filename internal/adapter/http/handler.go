package http

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"

	v1 "go-test-api/internal/adapter/http/v1"
	"go-test-api/internal/model"
	"go-test-api/internal/port"

	"github.com/labstack/echo/v4"
)

type svc struct {
	e    *echo.Echo
	addr string
}

var _ port.ProcessStartStopper = (*svc)(nil)

func (s *svc) Start(ctx context.Context, l net.Listener) {
	s.e.Listener = l
	if err := s.e.Start(s.addr); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal("shutting down the server")
	}
}

func (s *svc) Stop(ctx context.Context) {
	if err := s.e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}

// NewHTTPServer
// @title GO DOCUMENT API DOCUMENTATION
// @version 1.0
// @description This is a go document api docs.
// @termsOfService http://swagger.io/terms/
//
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
//
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
//
// @host localhost:8080
// @BasePath /
// @schemes http
func NewHTTPServer(
	ctx context.Context,
	conf model.Config,
	uc port.Usecases,
) *svc {
	app := echo.New()
	svc := &svc{e: app, addr: fmt.Sprintf(":%d", conf.App.Port)}

	v1Group := app.Group("/api/v1")

	v1Group.GET("/custom", func(c echo.Context) error {
		return c.String(200, "hello custom")
	})

	v1.New(v1Group, uc)

	return svc
}
