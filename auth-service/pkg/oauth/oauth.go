package oauth

import (
	"time"

	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/jackc/pgx/v4"

	"github.com/vgarvardt/go-pg-adapter/pgx4adapter"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-oauth2/oauth2/v4/generates"

	pg "github.com/vgarvardt/go-oauth2-pg/v4"
)

func InitManager(pgxConn *pgx.Conn, secretCode string) (*manage.Manager, *pg.ClientStore) {
	manager := manage.NewDefaultManager()
	manager.MapAccessGenerate(generates.NewJWTAccessGenerate("", []byte(secretCode), jwt.SigningMethodHS512))
	adapter := pgx4adapter.NewConn(pgxConn)
	tokenStore, _ := pg.NewTokenStore(adapter, pg.WithTokenStoreGCInterval(time.Minute))
	defer tokenStore.Close()

	clientStore, _ := pg.NewClientStore(adapter)
	manager.MapTokenStorage(tokenStore)
	manager.MapClientStorage(clientStore)
	return manager, clientStore
}
