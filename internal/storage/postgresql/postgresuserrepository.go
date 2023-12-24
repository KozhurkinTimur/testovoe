package postgresql

import ( 
	"database/sql" 
	//"github.com/lib/pq" 
	//"testovoe/internal/models"
)

type PostgresUserRepository struct {}

func (r *PostgresUserRepository) Connect() (*sql.DB, error) {
	db, err := sql.Open("postgres", "user=user password=root host=localhost dbname=users sslmode=disable") //sslmode?
	if err != nil {
		//log.Fatalf("Error: Unable to connect to database: %v", err) 
		return nil, err
	}
	defer db.Close() //defer?
	return db, nil
}

// func (r *PostgresUserRepository) Save(user *models.User) {
// 	db, err := r.Connect()
// }

// func (r *PostgresUserRepository) FindById(id string) (*models.User) {
// 	db, err := r.Connect()
// }

// func (r *PostgresUserRepository) FindByLogin(login string) (*models.User) {
// 	db, err := r.Connect()
// }
