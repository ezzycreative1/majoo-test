package user

import (
	"github.com/ezzycreative1/majoo-test/external/requests"
	"github.com/ezzycreative1/majoo-test/models"

	"github.com/gin-gonic/gin"
)

// Repository ..
type Repository interface {
	Insert(*models.User) (*string, error)
	Update(*models.User) error
	Get(string) (*models.User, error)
	GetById(string) (*models.User, error)
	GetUsers() (*[]models.User, error)
	DeleteUser(string) (*models.User, error)
	DeleteUsers([]string) error
	UpdateUser(req requests.UpdateUser) error
	ChangePhoto(userid string, c *gin.Context) error
	ChangeName(req requests.ChangeName) error
	CheckVerifyTokenExist(token, userID string) bool
	GetUserForSendEmailRegister(userID, token string) (models.User, error)
	VerifyAccount(token, userID string) bool
	UpdatePassword(req requests.ChangePasswordParam) error
	CreateNewAdmin(user models.User) error
	FetchAllAdmins(offset int) ([]models.User, error)
	FetchAdminByID(ID string) (models.User, error)
	UpdateAdminBySuperuser(user models.User) (bool, error)
	DeleteAdminBySuperuser(ID string) (bool, error)
	GetUserById(userID string) (user models.User, err error)
	GetUserIDByEmail(email string) (string, error)
	IsUserVerified(userID string) bool
	ResetPassword(token string, userID string, pass string) error
	GetPasswordUser(userID string) (string, error)
}
