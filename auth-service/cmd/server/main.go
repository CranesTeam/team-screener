package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/CranesTeam/team-screener/auth/pkg/model"
	"github.com/CranesTeam/team-screener/auth/pkg/oauth"
	"github.com/CranesTeam/team-screener/auth/pkg/repository"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
)

func init() {
	log.SetFormatter(new(log.JSONFormatter))

	if err := InitConfig(); err != nil {
		log.Fatalf("error initialisation config %s", err.Error())
	}
}

func main() {
	log.Info("Starting OAuth2 server...")
	pgxConn, err := repository.NewPostgresDb(context.TODO(), repository.Config{
		Host:     viper.GetString("db.postgres.host"),
		Port:     viper.GetString("db.postgres.port"),
		Username: viper.GetString("db.postgres.username"),
		Database: viper.GetString("db.postgres.database"),
		Password: viper.GetString("db.postgres.password"),
	})
	if err != nil {
		log.Fatalf("error occurred while connect to db:%s", err.Error())
	}

	manager, clientStore := oauth.InitManager(pgxConn, viper.GetString("auth.server.server"))

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
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
		uuid := uuid.New().String()[:16]
		requestBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			handleException(w, "Problem with body!", http.StatusInternalServerError, err)
		}

		var user model.UserRequest
		json.Unmarshal(requestBody, &user)
		log.Info("user:%s", user)

		clientStore.Create(&models.Client{
			ID:     uuid,
			Secret: user.Sercet,
			Domain: user.Domain,
			UserID: user.UserId,
		})

		resultEncoder(w, map[string]string{"clientId": uuid})
	})

	log.Fatal(http.ListenAndServe(viper.GetString("port"), nil))
}

func InitConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func handleException(w http.ResponseWriter, message string, code int, err error) {
	w.Header().Set("Content-Type", "application/text")
	w.WriteHeader(code)
	log.Println(message, err)
	http.Error(w, message, code)
}

func resultEncoder(w http.ResponseWriter, obj interface{}) {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(obj)
}
