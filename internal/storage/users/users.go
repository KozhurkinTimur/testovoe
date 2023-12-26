package users

import (
	"database/sql"
	"fmt"
	"testovoe/internal/models"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "user"
	password = "root"
	dbname   = "user"
)

type UserRepository struct {
	db *sql.DB
}

func NewUsersRepository() (*UserRepository, error) {
	//db, err := sql.Open("postgresql", "postgresql://user:root@127.0.0.1:5432/users sslmode=disable") //sslmode?
	psql := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psql)
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	//"postgres", "user=user password=root host=localhost dbname=users sslmode=disable"
	if err != nil {
		//log.Fatalf("Error: Unable to connect to database: %v", err)
		return nil, err
	}

	return &UserRepository{
		db: db,
	}, nil
}

func (r *UserRepository) Save(user *models.User) error {
	insertQuery := "INSERT INTO users (username, password) VALUES ($1, $2)"

	// Создание подготовленного выражения
	insertStmt, err := r.db.Prepare(insertQuery)
	if err != nil {
		return err
	}
	_, err = insertStmt.Exec(user.Login, user.Password)
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

// func (r *PostgresUserRepository) FindByLogin(login string) (*models.User) {
// 	db, err := r.Connect()
// }
