package app

import (
	"CarRentalService/internal/handler"
	"CarRentalService/internal/redis"
	"CarRentalService/internal/repository"
	"CarRentalService/internal/service"
	"CarRentalService/pkg/config"
	"CarRentalService/pkg/http_server"
	"CarRentalService/pkg/postgres"
	"context"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type App struct {
	server *http_server.Server
}

func NewApp() *App {
	return & App{
		server: new(http_server.Server),
	}
}

func (a *App) Run() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := config.InitConfig(); err != nil {
		logrus.Fatal(err.Error())
	}
	if err := config.LoadEnv(); err != nil {
		logrus.Fatalf("error of load env %s", err.Error())
	}
	db, err := postgres.NewPostgresDB(config.GetDBconfig())
	if err != nil {
		logrus.Fatal(err.Error())
	}
	redisClient, err := redis.NewClient()
	if err != nil {
		logrus.Fatalf("failed to initialize Redis: %s", err.Error())
	}

	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handler := handler.NewHandler(service, redisClient)

	go func () {
		if err := a.server.Run(config.GetPort(), handler.InitRouts()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()
	logrus.Println("Service started")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<- quit

	logrus.Println("server shutting down")
	if err := a.server.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
	if err := redisClient.Close(); err != nil {
		logrus.Errorf("error occured on redis client connection close: %s", err.Error())
	}
}	