package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	httpHandler "go-test-api/internal/adapter/http"

	"go-test-api/internal/adapter/repository"
	"go-test-api/internal/adapter/usecase"
	"go-test-api/internal/model"

	"github.com/soheilhy/cmux"
)

func main() {
	ctx := context.Background()
	repo := repository.NewRepository(ctx, model.PostgresClient{Db: "fal_db", Host: "localhost", Username: "postgres", Password: "Cyanogenmod@123", Port: "5432"})
	customerRepo := repo.GetCustomerRepository()

	uc := usecase.NewUsecases(model.Config{}, customerRepo)

	httpServer := httpHandler.NewHTTPServer(ctx, model.Config{App: model.App{Env: "localhost", Port: 8080}}, uc)

	// creating a listener for server
	nl, err := net.Listen("tcp", fmt.Sprintf(":%d", model.Config{App: model.App{Env: "localhost", Port: 8080}}.App.Port))
	if err != nil {
		fmt.Printf("tcp connection failure - %v", err)
	}

	m := cmux.New(nl)
	// a different listener for HTTP1
	httpL := m.Match(cmux.HTTP1Fast())

	// start http server and proxy calls to gRPC server endpoint
	go func() {
		httpServer.Start(ctx, httpL)
	}()

	// simple graceful shutdown
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	wg := &sync.WaitGroup{}
	wg.Add(1)
	ctx, fn := context.WithTimeout(context.Background(), model.Config{App: model.App{Env: "localhost", Port: 8080}}.App.GracefulTimeout)
	defer fn()
	go func() {
		defer wg.Done()
		httpServer.Stop(ctx)
	}()

	wg.Wait()
}
