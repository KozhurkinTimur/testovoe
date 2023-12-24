package storage
 
import ( 
	"testovoe/internal/models"
)

type UserRepository interface {
	Save(user *models.User)
	FindById(uid string) *models.User
	FindByLogin(login string) *models.User
}
