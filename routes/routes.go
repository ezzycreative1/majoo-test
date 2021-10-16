package app

import (
	HCHandler "github.com/ezzycreative1/majoo-test/app/healthcheck/handler"
	UsInterfaces "github.com/ezzycreative1/majoo-test/app/user"
	UserHandler "github.com/ezzycreative1/majoo-test/app/user/handler"
	"github.com/gin-gonic/gin"

	"github.com/ezzycreative1/majoo-test/middleware"
)

// HealthCheckHTTPHandler routes
func HealthCheckHTTPHandler(router *gin.Engine) {
	handler := &HCHandler.HealthCheckHandler{}
	router.GET("/check", handler.Check)
}

func NewUserHandler(router *gin.Engine, uc UsInterfaces.IUserUsecase) {
	handler := &UserHandler.UserHandler{
		RUsecase: uc,
	}

	//router.Use(cors.Default())
	user := router.Group("/user")
	{
		user.POST("/login", handler.Login)
		user.POST("/refresh", middleware.AuthorizeUserMiddleware(), handler.Refresh)
		user.GET("/userID", handler.ExtractToken)
		user.GET("/profile", handler.GetProfil)
		user.POST("/register", handler.Register)
		user.GET("/verify", handler.VerifyToken)
		user.POST("/forgot/password", handler.ForgotPassword)
		user.POST("/change-password", middleware.AuthorizeUserMiddleware(), handler.ChangePasswordHandler)
		user.POST("/reset-password", handler.ResetPassword)
	}
}
