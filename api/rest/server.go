package rest

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	cfg "github.com/Laem20957/records-app/configurations"
	"github.com/Laem20957/records-app/internal/logger"
	"github.com/sirupsen/logrus"
)

type HttpServer struct {
	server *http.Server
	logs   *logrus.Logger
}

func (hs *HttpServer) HttpServerSettings() *HttpServer {
	initRoutes := InitRoutes()
	logger := logger.CreateLogs()

	return &HttpServer{
		logs: logger.Log(),
		server: &http.Server{
			Addr: fmt.Sprintf(
				"%s:%d",
				cfg.InitConfigs().LocalServerHost,
				cfg.InitConfigs().LocalServerPort,
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
			hs.logs.Fatal("Error while server startup:", err)
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
		hs.logs.Fatal("Error while server stop:", err)
	} else {
		hs.logs.Info("Stoping server...")
	}
	hs.logs.Info("Server was stoped")
}
