package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

type HttpServer struct {
	logger *logrus.Logger
	server *http.Server
}

func Logger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	return logger
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func HttpServerSettings() *HttpServer {
	logger := Logger()
	handler := http.HandlerFunc(Handler)

	return &HttpServer{
		logger: logger,
		server: &http.Server{
			Addr:           ":8080",
			Handler:        handler,
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
			hs.logger.Fatal("Ошибка при запуске сервера:", err)
		}
	}()

	hs.logger.Info("Сервер запущен")
}

func (hs *HttpServer) HttpServerStop() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := hs.server.Shutdown(ctx)
	if err != nil {
		hs.logger.Fatal("Ошибка при остановке сервера:", err)
	} else {
		hs.logger.Info("Остановка сервера...")
	}

	hs.logger.Info("Сервер остановлен")
}
