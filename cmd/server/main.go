package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
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

// @title Team screener app
// @version 0.0.1-SNAPSHOT
// @description Skills services

// @BasePath /

// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

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

	<-ctx.Done()
	stop()
	logrus.Println("shutting down gracefully, press Ctrl+C again to force")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logrus.Fatal("Server forced to shutdown: ", err)
	}

	logrus.Println("Server exiting")
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
