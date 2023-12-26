package storage
 
import ( 
	"testovoe/internal/models"
)

type UserRepository interface {
	Save(user *models.User) error
	FindById(uid string) (*models.User, error)
	FindByLogin(login string) (*models.User, error)
}
