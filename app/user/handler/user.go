package handler

import (

	// "net/smtp"

	"encoding/json"
	"errors"
	"fmt"
	"strings"

	BaseHandler "github.com/ezzycreative1/majoo-test/app/base/handler"
	UsInterface "github.com/ezzycreative1/majoo-test/app/user"
	middleware "github.com/ezzycreative1/majoo-test/app/user/usecase"
	"github.com/ezzycreative1/majoo-test/helpers/response_mapping"

	"github.com/ezzycreative1/majoo-test/external/requests"

	//"time"

	//UsInterface "backend/app/user"
	//"backend/requests"

	"github.com/ezzycreative1/majoo-test/helpers"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// UserHandler ..
type UserHandler struct {
	RUsecase UsInterface.IUserUsecase
}

// Login ..
func (u *UserHandler) Login(c *gin.Context) {
	requestBody := requests.LoginRequest{}
	err := c.BindJSON(&requestBody)
	if err != nil {

		BaseHandler.ResponseFailed(c, BaseHandler.BaseError{
			Error: err,
			Code:  response_mapping.InvalidParam,
		})
		return
	}

	// validate
	if err := u.RUsecase.RequestLoginValid(requestBody); err != nil {
		BaseHandler.FailedResponseBackend(c, err)
		return
	}

	user, err := u.RUsecase.GetUserByEmail(requestBody.Email)
	// -- validate if user not
	if user == nil {
		BaseHandler.ResponseFailed(c, BaseHandler.BaseError{
			Error: errors.New("Email not register."),
			Code:  response_mapping.LoginUserNotFound,
		})
		return
	}

	if !u.RUsecase.GetVerifiedUsers(user.ID) {
		BaseHandler.ResponseFailed(c, BaseHandler.BaseError{
			Error: errors.New("Your account is not verified. Please verify email first before you login."),
			Code:  response_mapping.AccountSuspend,
		})
		return
	}

	//return
	resLogin, err := u.RUsecase.Login(c, requestBody)
	//userid := u.RUsecase.GetUserIDByEmail(helpers.EncryptBase64(requestBody.Email))

	resJSON, _ := json.Marshal(resLogin)
	var m map[string]json.RawMessage
	if err := json.Unmarshal([]byte(resJSON), &m); err != nil {
		BaseHandler.ResponseFailed(c, BaseHandler.BaseError{
			Error: err,
			Code:  response_mapping.GeneralError,
		})
		return
	}
	delete(m, "token")

	BaseHandler.RespondSuccess(c, "success", resLogin)
	return
}

func (u *UserHandler) Refresh(c *gin.Context) {

	user, err := u.RUsecase.GetProfil(c)

	// -- validate if user not
	if user == nil {
		BaseHandler.ResponseFailed(c, BaseHandler.BaseError{
			Error: errors.New("Email not register."),
			Code:  response_mapping.LoginUserNotFound,
		})
		return
	}

	if !u.RUsecase.GetVerifiedUsers(user.ID) {
		BaseHandler.ResponseFailed(c, BaseHandler.BaseError{
			Error: errors.New("Your account is not verified. Please verify email first before you login."),
			Code:  response_mapping.AccountSuspend,
		})
		return
	}

	//return
	resLogin, err := u.RUsecase.LoginRefresh(c, user)
	if err != nil {
		//save to log error
		BaseHandler.ResponseFailed(c, BaseHandler.BaseError{
			Error: err,
			Code:  response_mapping.InvalidParam,
		})
		return
	}
	BaseHandler.RespondSuccess(c, "success", resLogin)
	return
}

func (u *UserHandler) ExtractToken(c *gin.Context) {
	userid := u.RUsecase.GetUserID(c)
	BaseHandler.RespondSuccess(c, "", userid)

}

func (u *UserHandler) GetProfil(c *gin.Context) {
	profil, err := u.RUsecase.GetProfil(c)
	if err != nil {
		BaseHandler.FailedResponseBackend(c, err)
		return
	}
	BaseHandler.RespondSuccess(c, "", profil)
}

func (u *UserHandler) Register(c *gin.Context) {
	requestBody := requests.RegisterRequest{}
	err := c.ShouldBind(&requestBody)
	if err != nil {
		c.Error(err)
		BaseHandler.ResponseFailed(c, BaseHandler.BaseError{
			Error: err,
			Code:  response_mapping.InvalidParam,
		})
		return
	}

	// register
	errE := u.RUsecase.Register(c, requestBody)
	if errE != nil {
		BaseHandler.ResponseFailed(c, *errE)
		return
	}

	BaseHandler.ResponseSuccess(c, "please check your email to complete your registration", nil)
	return
}

func (u *UserHandler) VerifyToken(c *gin.Context) {
	token := c.Query("token")

	// verify token
	if err := u.RUsecase.VerifyToken(c, token); err != nil {
		BaseHandler.FailedResponseBackend(c, err)
		return
	}

	BaseHandler.RespondSuccess(c, "verify account success", nil)
	return
}

// ForgotPassword ..
func (u *UserHandler) ForgotPassword(c *gin.Context) {

	// validate requiest
	var req requests.ForgotPassword
	if err := c.ShouldBindJSON(&req); err != nil {
		c.Error(err)
		BaseHandler.ResponseFailed(c, BaseHandler.BaseError{
			Error: err,
			Code:  response_mapping.EmailNotRegister,
		})
		return
	}

	// check user exist
	if err := u.RUsecase.ForgotPassword(helpers.EncryptBase64(req.Email)); err != nil {
		BaseHandler.ResponseFailed(c, *err)
		return
	}

	BaseHandler.RespondSuccess(c, "please check your email to validate your account", "")
	return
}

// ChangePassword ..
func (u *UserHandler) ChangePasswordHandler(c *gin.Context) {

	requestBody := requests.ChangePasswordRequest{}
	fmt.Println(requestBody.OldPassword)
	err := c.BindJSON(&requestBody)
	if err != nil {
		c.Error(err)
		BaseHandler.ResponseFailed(c, BaseHandler.BaseError{
			Error: err,
			Code:  response_mapping.InvalidParam,
		})
		return
	}

	userID := middleware.GetUserID(c)
	fmt.Println(userID)

	var param = requests.ChangePasswordParam{}
	param.UserID = userID
	param.OldPassword = strings.TrimSpace(requestBody.OldPassword)
	param.NewPassword = strings.TrimSpace(requestBody.NewPassword)
	param.ConfirmPassword = strings.TrimSpace(requestBody.ConfirmPassword)

	if requestBody.OldPassword == "" || requestBody.NewPassword == "" || requestBody.ConfirmPassword == "" {
		c.Error(err)
		BaseHandler.ResponseFailed(c, BaseHandler.BaseError{
			Error: err,
			Code:  response_mapping.ChangePasswordInvalidParam,
		})
		return
	}

	if err := u.RUsecase.RequestChangePasswordValid(requestBody.OldPassword); err != nil {
		BaseHandler.FailedResponseBackend(c, err)
		return
	}

	if err := u.RUsecase.RequestChangePasswordValid(requestBody.NewPassword); err != nil {
		BaseHandler.FailedResponseBackend(c, err)
		return
	}

	if err := u.RUsecase.ChangePassword(c, param); err != nil {
		BaseHandler.FailedResponseBackend(c, err)
		return
	}

	BaseHandler.RespondSuccess(c, "", "Success Updated Password")
	return
}

// EmailResetPassword ..
func (u *UserHandler) EmailResetPassword(c *gin.Context) {

	type Request struct {
		Email string `json:"email"`
	}

	var req Request

	if err := c.ShouldBindJSON(&req); err != nil {
		// save to log
		e := fmt.Errorf("cannot read request body")
		c.Error(e)
		BaseHandler.FailedResponseBackend(c, e)
		return
	}

	req.Email = strings.TrimSpace(req.Email)

	v := validator.New()
	if err := v.Var(req.Email, "required,email"); err != nil {
		c.Error(err)
		BaseHandler.FailedResponseBackend(c, err)
		return
	}

	if err := u.RUsecase.ForgotPassword(req.Email); err != nil {
		var err = fmt.Errorf("email not found")
		c.Error(err)
		BaseHandler.FailedResponseBackend(c, err)
		return
	}

	BaseHandler.RespondSuccess(c, "", "please check your email")
	return
}

func (u *UserHandler) ResetPassword(c *gin.Context) {

	requestBody := requests.ResetPassword{}

	err := c.BindJSON(&requestBody)
	if err != nil {
		c.Error(err)
		BaseHandler.FailedResponseBackend(c, err)
		return
	}

	err = u.RUsecase.ResetPassword(c, requestBody)
	if err != nil {
		c.Error(err)
		BaseHandler.FailedResponseBackend(c, err)
		return
	}

	BaseHandler.RespondSuccess(c, "", "success reset password")
	return
}

// func (u *UserHandler) EditUser(c *gin.Context) {
// 	requestBody := requests.EditRequest{}
// 	err := c.ShouldBind(&requestBody)
// 	if err != nil {
// 		c.Error(err)
// 		BaseHandler.ResponseFailed(c, BaseHandler.BaseError{
// 			Error: err,
// 			Code:  response_mapping.InvalidParam,
// 		})
// 		return
// 	}

// 	// register
// 	errE := u.RUsecase.EditUser(c, requestBody)
// 	if errE != nil {
// 		BaseHandler.ResponseFailed(c, *errE)
// 		return
// 	}

// 	BaseHandler.ResponseSuccess(c, "success edit user", nil)
// 	return
// }
