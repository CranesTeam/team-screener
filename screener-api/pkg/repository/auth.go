package repository

import (
	"fmt"

	m "github.com/CranesTeam/team-screener/pkg/model"
	"github.com/jmoiron/sqlx"
)

const (
	defaultRole = "USER"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) GetUserRoleId() (int, error) {
	var id int
	query := fmt.Sprintf("select id from %s where name=$1", userRolesTable)
	err := r.db.Get(&id, query, defaultRole)

	return id, err
}

func (r *AuthRepository) CreateUser(user m.User, userInfo m.UserInfo) (string, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return "empty", err
	}

	var userId int
	query := fmt.Sprintf("insert into %s (username, password_hash, role_id) values ($1, $2, $3) returning id", usersTable)
	row := tx.QueryRow(query, user.Username, user.PasswordHash, user.RoleId)
	if err = row.Scan(&userId); err != nil {
		tx.Rollback()
		return "empty", err
	}

	infoQuery := fmt.Sprintf("insert into %s (user_id, name, email) values ($1, $2, $3)", userInfoTable)
	_, err = tx.Exec(infoQuery, userId, userInfo.Name, userInfo.Email)
	if err != nil {
		tx.Rollback()
		return "empty", err
	}

	var userUuid string
	userQuery := fmt.Sprintf("select external_uuid from %s where id=$1", usersTable)
	userRow := tx.QueryRow(userQuery, userId)
	if err = userRow.Scan(&userUuid); err != nil {
		tx.Rollback()
		return "empty", err
	}

	return userUuid, tx.Commit()
}

func (r *AuthRepository) GetUser(username string) (m.User, error) {
	var user m.User
	query := fmt.Sprintf("select u.id, u.external_uuid, u.username, u.password_hash, u.role_id, r.name as role_name from %s u inner join %s r on (u.role_id = r.id) where username=$1", usersTable, userRolesTable)
	err := r.db.Get(&user, query, username)

	return user, err
}
