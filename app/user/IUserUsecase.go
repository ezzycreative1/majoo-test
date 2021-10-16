package user

import (
	"github.com/ezzycreative1/majoo-test/app/base/handler"
	"github.com/ezzycreative1/majoo-test/external/requests"
	"github.com/ezzycreative1/majoo-test/models"

	"github.com/gin-gonic/gin"
)

// IUserUsecase ..
type IUserUsecase interface {
	VerifyToken(c *gin.Context, token string) error
	Register(c *gin.Context, request requests.RegisterRequest) *handler.BaseError
	Login(c *gin.Context, request requests.LoginRequest) (*models.UserResponse, error)
	LoginRefresh(c *gin.Context, user *models.User) (*models.UserResponse, error)
	RequestLoginValid(req requests.LoginRequest) error
	GetUserID(c *gin.Context) string
	GetUserIDByEmail(email string) string
	GetUserByEmail(email string) (*models.User, error)
	GetProfil(c *gin.Context) (*models.User, error)
	ForgotPassword(email string) *handler.BaseError
	ChangePassword(c *gin.Context, req requests.ChangePasswordParam) error
	RequestChangePasswordValid(pass string) error
	ResetPassword(c *gin.Context, req requests.ResetPassword) error
	GetVerifiedUsers(id string) bool
}
