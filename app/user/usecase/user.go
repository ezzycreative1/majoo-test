package usecase

import (
	"unicode"

	"github.com/ezzycreative1/majoo-test/app/base/handler"
	"github.com/ezzycreative1/majoo-test/helpers/response_mapping"

	//"encoding/json"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/ezzycreative1/majoo-test/app/user"
	userInterfaces "github.com/ezzycreative1/majoo-test/app/user"
	"github.com/ezzycreative1/majoo-test/cache"
	"github.com/ezzycreative1/majoo-test/external/requests"
	"github.com/ezzycreative1/majoo-test/helpers"
	"github.com/ezzycreative1/majoo-test/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	repo user.Repository
}

// NewUserUsecase ..
func NewUserUsecase(repo user.Repository) userInterfaces.IUserUsecase {
	return &userUsecase{
		repo: repo,
	}
}

func (us *userUsecase) GetUserByEmail(email string) (*models.User, error) {
	emails := helpers.EncryptBase64(email)

	user, err := us.repo.Get(emails)
	if err != nil {
		return nil, errors.New("Wrong Email")
	}

	return user, nil
}

func (us *userUsecase) GetVerifiedUsers(id string) bool {
	if !us.repo.IsUserVerified(id) {
		// return nil, errors.New("User is not verified yet!")
		return false
	}

	return true
}

// Login ..
func (us *userUsecase) Login(c *gin.Context, request requests.LoginRequest) (*models.UserResponse, error) {
	response := &models.UserResponse{}
	if request.Email == "" || request.Password == "" {
		return nil, errors.New("Email Or Password Cannot Blank")
	}

	checkEmail := helpers.ValidEmail(request.Email)
	if !checkEmail {
		return nil, errors.New("Format Email Wrong")
	}

	//check  email
	request.Email = helpers.EncryptBase64(request.Email)
	user, err := us.repo.Get(request.Email)
	if err != nil {
		return nil, errors.New("Wrong Email")
	}

	//check password
	checkPassword := checkPasswordHash(request.Password, user.Password)
	if checkPassword == false {
		return nil, errors.New("Wrong Password")
	}

	if !us.repo.IsUserVerified(user.ID) {
		return nil, errors.New("Your account is not verified. Please verify email first before you login.")

	}

	getToken, err := GenerateToken(user.ID, user.Email)
	if err != nil {
		return nil, errors.New("Genarete Token Failed")
	}
	response.Token = getToken

	return response, nil

}

func (us *userUsecase) LoginRefresh(c *gin.Context, user *models.User) (*models.UserResponse, error) {
	response := &models.UserResponse{}

	if !us.repo.IsUserVerified(user.ID) {
		// return nil, errors.New("User is not verified yet!")
		return nil, errors.New("Your account is not verified. Please verify email first before you login.")

	}

	getToken, err := GenerateToken(user.ID, user.Email)
	if err != nil {
		return nil, errors.New("Genarete Token Failed")
	}
	response.Token = getToken

	return response, nil

}

func checkPasswordHash(password, hash string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (us *userUsecase) RequestLoginValid(req requests.LoginRequest) error {

	var v = validator.New()
	var errString []string

	req.Email = strings.TrimSpace(req.Email)

	if e := v.Var(req.Email, "required,email"); e != nil {
		errString = append(errString, "email")
	}

	if req.Password == "" {
		errString = append(errString, "password")
	}

	if len(errString) != 0 {
		return fmt.Errorf("field %s required", strings.Join(errString, ", "))
	}

	return nil
}

func (us *userUsecase) RequestUpdateUserValid(req requests.UpdateUser) error {

	var errString []string

	req.Name = strings.TrimSpace(req.Name)
	req.NoHp = strings.TrimSpace(req.NoHp)

	if req.Name == "" {
		errString = append(errString, "field name required")
	}

	if req.NoHp == "" {

		errString = append(errString, "no hp required")

	} else if _, e := strconv.Atoi(req.NoHp); e != nil {

		errString = append(errString, "no hp must be number")

	} else {

		if len(req.NoHp) < 10 {
			errString = append(errString, "no hp minimum 10 digit")
		}

		if len(req.NoHp) > 15 {
			errString = append(errString, "no hp maximum 15 digit")
		}

	}

	return nil
}

func (us *userUsecase) GetUserID(c *gin.Context) string {
	userID := GetUserID(c)
	return userID
}

func (us *userUsecase) GetUserIDByEmail(email string) string {
	userid, err := us.repo.GetUserIDByEmail(email)

	if err != nil {
		return ""
		//log.Fatalln(err.Error())
	}

	return userid
}

func (us *userUsecase) GetProfil(c *gin.Context) (*models.User, error) {
	userID := GetUserID(c)
	dataUser, err := us.repo.GetById(userID)

	dataUser.Password = ""
	dataUser.Email = helpers.DecryptBase64(dataUser.Email)

	if err != nil {
		return nil, err
	}
	return dataUser, nil
}

func (us *userUsecase) Register(c *gin.Context, request requests.RegisterRequest) *handler.BaseError {
	var errString []string

	v := validator.New()
	if err := v.Struct(&request); err != nil {
		return &handler.BaseError{
			Error: err,
			Code:  response_mapping.InvalidParam,
		}
	}

	if _, e := strconv.Atoi(request.PhoneNumber); e != nil {

		errString = append(errString, "phone_number must be number")

	} else {

		if len(request.PhoneNumber) < 10 {
			errString = append(errString, "phone_number minimum 10 digit")
		}

		if len(request.PhoneNumber) > 15 {
			errString = append(errString, "phone_number maximum 15 digit")
		}
	}

	if len(errString) != 0 {
		return &handler.BaseError{
			Error: fmt.Errorf("field %s", strings.Join(errString, ", ")),
			Code:  response_mapping.InvalidParam,
		}
	}

	// Generate Random uuID
	u, err := uuid.NewRandom()
	if err != nil {
		return &handler.BaseError{
			Error: fmt.Errorf("failed to generate UUID: %w", err),
			Code:  response_mapping.GeneralError,
		}
	}

	userId := u.String()
	emailEncrypt := helpers.EncryptBase64(request.Email)
	modelUser := &models.User{}

	// Duplicate Email Validation
	_, err = us.repo.Get(emailEncrypt)
	if err == nil {
		return &handler.BaseError{
			Error: errors.New("email duplicate error"),
			Code:  response_mapping.RegisterEmailDuplicate,
		}
	}

	password := hashPasswordUser(request.Password)

	modelUser.ID = userId
	modelUser.Fullname = strings.Title(request.Fullname)
	modelUser.Email = emailEncrypt
	modelUser.Password = password
	modelUser.PhoneNumber = request.PhoneNumber
	modelUser.Status = "active"

	_, err = us.repo.Insert(modelUser)

	if err != nil {
		return &handler.BaseError{
			Error: err,
			Code:  response_mapping.GeneralError,
		}
	}

	return nil
}

func hashPasswordUser(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	if err != nil {
		log.Println(err.Error())
	}
	return string(bytes)
}

// Verifying user token ..
func (us *userUsecase) VerifyToken(c *gin.Context, token string) error {
	key := fmt.Sprintf("email_verif_token:%s", token)
	result, err := cache.Get(key)
	if err != nil {
		return errors.New("Token is invalid")
	}
	userid := strings.Split(result, ":")[1]

	if us.repo.IsUserVerified(userid) {
		return errors.New("User has already verified")
	}

	if !us.repo.VerifyAccount(token, userid) {
		return err
	}

	if err = cache.Del(key); err != nil {
		return errors.New("Can't delete key from redis")
	}

	user, err := us.repo.GetById(userid)
	if err != nil {
		return err
	}

	err = us.repo.Update(user)
	if err != nil {
		return err
	}

	return nil
}

func (us *userUsecase) ForgotPassword(email string) *handler.BaseError {

	// check user exist
	user, err := us.repo.Get(email)

	if err != nil {
		return &handler.BaseError{
			Error: err,
			Code:  response_mapping.GeneralError,
		}
	}

	key := helpers.GenerateRandString(60)
	if err := cache.Set(key, user.ID, 30*time.Minute); err != nil {
		fmt.Println(err.Error())
		return &handler.BaseError{
			Error: err,
			Code:  response_mapping.GeneralError,
		}
	}

	return nil
}

func (us *userUsecase) ChangePassword(c *gin.Context, param requests.ChangePasswordParam) error {
	var err error

	err = us.RequestChangePasswordValid(param.OldPassword)

	if err != nil {
		return err
	}

	err = us.RequestChangePasswordValid(param.NewPassword)

	if err != nil {
		return err
	}

	if param.OldPassword == param.NewPassword {
		return errors.New("Passwords may not be the same as before")
	}

	hashPassword, err := hashPassword(param.NewPassword)

	if err != nil {
		return err
	}

	old_pass, err := us.repo.GetPasswordUser(param.UserID)

	if err != nil {
		return err
	}

	if hashPassword == old_pass {
		return errors.New("new password same with old password")
	}

	confirm_old := checkPasswordHash(param.OldPassword, old_pass)

	if !confirm_old {
		return errors.New("old_password not match")
	}

	confirm := checkPasswordHash(param.ConfirmPassword, hashPassword)

	if !confirm {
		return errors.New("New Password and Confirm Password Not Match")
	}

	param.NewPassword = hashPassword

	err = us.repo.UpdatePassword(param)

	if err != nil {
		return errors.New("Failed Change Password")
	}

	return err
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	return string(bytes), err
}

func (us *userUsecase) RequestChangePasswordValid(pass string) error {

	var errString []string

	var isStringAlphabetic = regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString

	// check password falid
	var upper, lower, number, symbol bool = false, false, false, false
	for _, r := range pass {

		if !isStringAlphabetic(string(r)) {
			symbol = true
			continue
		}

		if unicode.IsUpper(r) {
			upper = true
			continue
		}

		if unicode.IsLower(r) {
			lower = true
			continue
		}

		if unicode.IsDigit(r) {
			number = true
			continue
		}
	}

	if len(pass) < 8 {
		errString = append(errString, "must be 8 character")
	}

	if !symbol {
		errString = append(errString, "must be contain symbol")
	}

	if !upper {
		errString = append(errString, "must be contain uppercase")
	}

	if !lower {
		errString = append(errString, "must be contain lowercase")
	}

	if !number {
		errString = append(errString, "must be contain number")
	}

	if len(errString) > 0 {
		return fmt.Errorf("new password or old password %s", strings.Join(errString, ", "))
	}

	return nil
}

func (us *userUsecase) ResetPassword(c *gin.Context, req requests.ResetPassword) error {
	key := fmt.Sprintf("email_reset_password_token:%s", req.Token)
	result, err := cache.Get(key)
	if err != nil {
		return errors.New("Token is invalid")
	}
	userid := strings.Split(result, ":")[1]

	err = us.RequestChangePasswordValid(req.NewPassword)

	if err != nil {
		return err
	}

	hashPassword, err := hashPassword(req.NewPassword)

	if err != nil {
		return err
	}

	confirm := checkPasswordHash(req.ConfirmPassword, hashPassword)

	if !confirm {
		return errors.New("New Password and Confirm Password Not Match")
	}

	if err = us.repo.ResetPassword(req.Token, userid, hashPassword); err != nil {
		return err
	}

	if err = cache.Del(key); err != nil {
		return errors.New("Can't delete key from redis")
	}
	return nil
}
