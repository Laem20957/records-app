package servers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	config "github.com/Laem20957/records-app/configs"
	"github.com/Laem20957/records-app/internal/transport/rest"
	"github.com/Laem20957/records-app/pkg/logger"
	"github.com/sirupsen/logrus"
)

type HttpServer struct {
	server  *http.Server
	logs    *logrus.Logger
	handler *rest.Handler
}

func (hs *HttpServer) HttpServerSettings() *HttpServer {
	init_routes := hs.handler.InitRoutes()
	logger := logger.CreateLogs()

	return &HttpServer{
		logs: logger,
		server: &http.Server{
			Addr: fmt.Sprintf(
				"%s:%d",
				config.InitConfigs().LocalServerHost,
				config.InitConfigs().LocalServerPort,
			),
			Handler:        init_routes,
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
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	context, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := hs.server.Shutdown(context)
	if err != nil {
		hs.logs.Fatal("Error while server stop:", err)
	} else {
		hs.logs.Info("Stoping server...")
	}

	hs.logs.Info("Server was stoped")
}
