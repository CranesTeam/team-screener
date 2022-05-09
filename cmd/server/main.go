package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/CranesTeam/team-screener/pkg/handler"
	"github.com/CranesTeam/team-screener/pkg/repository"
	"github.com/CranesTeam/team-screener/pkg/server"
	"github.com/CranesTeam/team-screener/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15,
		"the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	if err := InitConfig(); err != nil {
		log.Fatalf("error initialisation config %s", err.Error())
	}

	log.Println("init hander and start server")
	repo := repository.NewRepository()
	service := service.NewService(repo)
	handlers := handler.NewHandler(service)

	srv := new(server.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while runnig http server: %s", err.Error())
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}

func InitConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
