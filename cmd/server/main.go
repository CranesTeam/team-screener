package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"time"

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

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15,
		"the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

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
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while runnig http server: %s", err.Error())
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)
	logrus.Println("shutting down")
	os.Exit(0)
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
