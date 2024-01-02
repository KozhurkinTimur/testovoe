package users

import (
	"database/sql"
	"fmt"
	"log"
	"testovoe/internal/models"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "user"
	password = "root"
	dbname   = "postgres"
)

type UserRepository struct {
	db *sql.DB
}

func NewUsersRepository() (*UserRepository, error) {
	psql := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psql)
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	
	if err != nil {
		log.Fatalf("Error: Unable to connect to database: %v", err)
		return nil, err
	}

	return &UserRepository{
		db: db,
	}, nil
}

func (r *UserRepository) Save(user *models.User) error {
	insertQuery := "INSERT INTO users (uid, login, password) VALUES ($1, $2, $3)"

	insertStmt, err := r.db.Prepare(insertQuery)
	if err != nil {
		return err
	}
	_, err = insertStmt.Exec(user.Id ,user.Login, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) FindById(uid string) (*models.User, error) {
	return &models.User{}, nil
}

func (r *UserRepository) FindByLogin(login string) (*models.User, error) {
	return &models.User{}, nil
}

