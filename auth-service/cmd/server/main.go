package main

import (
	"context"
	"net/http"
	"time"

	"github.com/CranesTeam/team-screener/auth/pkg/repository"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-oauth2/oauth2/v4/generates"

	pg "github.com/vgarvardt/go-oauth2-pg/v4"
	"github.com/vgarvardt/go-pg-adapter/pgx4adapter"
)

var (
	portvar int
)

func init() {
	log.SetFormatter(new(log.JSONFormatter))

	if err := InitConfig(); err != nil {
		log.Fatalf("error initialisation config %s", err.Error())
	}
}

func main() {
	logrus.Info("Starting OAuth2 server...")
	pgxConn, _ := repository.NewPostgresDb(context.TODO(), repository.Config{
		Host:     viper.GetString("db.postgres.host"),
		Port:     viper.GetString("db.postgres.port"),
		Username: viper.GetString("db.postgres.username"),
		Database: viper.GetString("db.postgres.database"),
		Password: viper.GetString("db.postgres.password"),
	})

	manager := manage.NewDefaultManager()
	// todo: set code by config
	manager.MapAccessGenerate(generates.NewJWTAccessGenerate("", []byte("secret_team_code"), jwt.SigningMethodHS512))
	adapter := pgx4adapter.NewConn(pgxConn)
	tokenStore, _ := pg.NewTokenStore(adapter, pg.WithTokenStoreGCInterval(time.Minute))
	defer tokenStore.Close()

	clientStore, _ := pg.NewClientStore(adapter)
	manager.MapTokenStorage(tokenStore)
	manager.MapClientStorage(clientStore)

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		logrus.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		logrus.Println("Response Error:", re.Error.Error())
	})

	http.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
		err := srv.HandleAuthorizeRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		srv.HandleTokenRequest(w, r)
	})

	// todo: add some logic and read POST request
	http.HandleFunc("/registration", func(w http.ResponseWriter, r *http.Request) {
		uuid := uuid.New().String()[:8]
		clientStore.Create(&models.Client{
			ID:     uuid,
			Secret: "test",
			Domain: "test",
			UserID: "test",
		})
		logrus.Info("find after saving...")
		logrus.Info(clientStore.GetByID(context.TODO(), uuid))
	})

	// todo: added graceful shutdown
	logrus.Fatal(http.ListenAndServe(viper.GetString("port"), nil))
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
