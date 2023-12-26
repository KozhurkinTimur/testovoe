package postgresql

import (
	"database/sql"
	"fmt"
	//"fmt"
	//"testovoe/internal/models"
	_ "github.com/lib/pq"
)
const (
	host = "localhost"
	port = 5432
	user = "user"
	password =  "root"
	dbname = "user"
)
type PostgresUserRepository struct{
	*sql.DB
}

func Connect() (*sql.DB, error) {
	//db, err := sql.Open("postgresql", "postgresql://user:root@127.0.0.1:5432/users sslmode=disable") //sslmode?
	psql := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psql)
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	//"postgres", "user=user password=root host=localhost dbname=users sslmode=disable"
	if err != nil {
		//log.Fatalf("Error: Unable to connect to database: %v", err)
		return nil, err
	}
	//defer db.Close() //defer?
	return db, nil
}

func (r *PostgresUserRepository) Save(login, password, id string) {
	// db, err := r.Connect()
	// err = db.Ping()
	// if err != nil {
	// 	panic(err)
	// }
	insertQuery := "INSERT INTO users (username, password, id) VALUES ($1, $2, $3)"

	// Создание подготовленного выражения
	insertStmt, err := r.Prepare(insertQuery)
	if err != nil {
		return
	}

	err = r.Ping()
	if err != nil {
		panic(err)
	}

	_, err = insertStmt.Exec(login, password, id)

	err = r.Ping()
	if err != nil {
		panic(err)
	}

	//defer insertStmt.Close()

	// Выполнение запроса с параметрами
	//  if err != nil {
	//   return fmt.Errorf("error executing insert query: %v", err)
	//  }

	//  return nil
}

// func (r *PostgresUserRepository) FindByLogin(id string) (*models.User) {
// 	db, err := r.Connect()
// }

// func (r *PostgresUserRepository) FindByLogin(login string) (*models.User) {
// 	db, err := r.Connect()
// }
