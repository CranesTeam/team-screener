package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/CranesTeam/team-screener/pkg/handler"
	"github.com/CranesTeam/team-screener/pkg/repository"
	"github.com/CranesTeam/team-screener/pkg/server"
	"github.com/CranesTeam/team-screener/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := InitConfig(); err != nil {
		logrus.Fatalf("error initialisation config %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error while readind env file %s", err.Error())
	}
}

// @title Team screener app
// @version 0.0.1-SNAPSHOT
// @description Skills services

// @BasePath /

// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	db, err := repository.NewPostgresDb(repository.Config{
		Host:     viper.GetString("db.postgres.host"),
		Port:     viper.GetString("db.postgres.port"),
		Username: viper.GetString("db.postgres.username"),
		Database: viper.GetString("db.postgres.database"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMode:  viper.GetString("db.postgres.sslMode"),
	})
	if err != nil {
		logrus.Fatalf("error while init db %s", err.Error())
	}

	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handlers := handler.NewHandler(service)

	srv := new(server.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while runnig http server: %s", err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	<-c

	logrus.Println("shutting down...")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured closing connections: %s", err.Error())
	}

	logrus.Println("done...")
	os.Exit(0)
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
