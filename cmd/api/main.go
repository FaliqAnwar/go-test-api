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
	"go-test-api/internal/logzap"

	"go-test-api/internal/adapter/repository"
	"go-test-api/internal/adapter/usecase"
	"go-test-api/internal/model"

	"github.com/soheilhy/cmux"
)

func main() {
	config := model.Config{
		App: model.App{
			Env:         "local",
			Port:        8080,
			Name:        "go-test-api",
			LogOption:   "console",
			LogLevel:    "debug",
			RPCAddress:  "localhost:8080",
			RPCInsecure: true,
		},
		PostgresClient: model.PostgresClient{
			Host: "localhost",
			//Db:       "go_document",
			//Username: "cp_eng",
			//Password: "H921hK7Dv20Al",
			Db:       "postgres",
			Username: "postgres",
			Password: "Cyanogenmod@123",
			Port:     "5432",
		},
	}

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	logger, sugarLogger := logzap.InitLogger(config)
	defer logger.Sync()
	defer sugarLogger.Sync()

	repo := repository.NewRepository(ctx, config.PostgresClient)
	customerRepo := repo.GetCustomerRepository()

	uc := usecase.NewUsecases(config, customerRepo)

	httpServer := httpHandler.NewHTTPServer(ctx, config, logger, uc)

	// creating a listener for server
	nl, err := net.Listen("tcp", fmt.Sprintf(":%d", config.App.Port))
	if err != nil {
		sugarLogger.Fatalf("tcp connection failure - %v", err)
	}

	m := cmux.New(nl)
	// a different listener for HTTP1
	httpL := m.Match(cmux.HTTP1Fast())

	// start http server and proxy calls to gRPC server endpoint
	go func() {
		httpServer.Start(ctx, httpL)
	}()

	go func() {
		if err := m.Serve(); err != nil {
			sugarLogger.Fatalf("serve cmux failure - %v", err)
		}
	}()

	// simple graceful shutdown
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	wg := &sync.WaitGroup{}
	wg.Add(1)
	ctx, fn := context.WithTimeout(context.Background(), config.App.GracefulTimeout)
	defer fn()
	go func() {
		defer wg.Done()
		httpServer.Stop(ctx)
	}()

	wg.Wait()
	sugarLogger.Info("all server stopped!")
}
