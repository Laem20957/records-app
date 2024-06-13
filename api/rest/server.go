package rest

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"records-app/internal/logger"
	"records-app/settings"

	"github.com/sirupsen/logrus"
)

type HttpServer struct {
	logs   *logrus.Logger
	server *http.Server
}

func (hs *HttpServer) NewHttpServer() *HttpServer {
	settings := settings.GetSettings()
	initRoutes := InitRoutes()
	logger := logger.CreateLogs()
	return &HttpServer{
		logs: logger.Log(),
		server: &http.Server{
			Addr: fmt.Sprintf(
				"%s:%d",
				settings.AppHost,
				settings.AppPort,
			),
			Handler:        initRoutes,
			MaxHeaderBytes: 1 << 20,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
		},
	}
}

func (hs *HttpServer) HttpServerStart() {
	go func() {
		err := hs.server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			hs.logs.Fatal(err)
		}
	}()
	hs.logs.Info("Server was started")
}

func (hs *HttpServer) HttpServerStop() {
	serverStop := make(chan os.Signal, 1)
	signal.Notify(serverStop, os.Interrupt, syscall.SIGTERM)
	<-serverStop

	serverContext, serverCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer serverCancel()

	err := hs.server.Shutdown(serverContext)
	if err != nil {
		hs.logs.Fatal(err)
	} else {
		hs.logs.Info("Stopping server...")
	}
	hs.logs.Info("Server was stopped")
}
