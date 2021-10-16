package repository

import (
	"fmt"
	"time"

	"github.com/ezzycreative1/majoo-test/app/user"
	"github.com/ezzycreative1/majoo-test/external/requests"
	"github.com/ezzycreative1/majoo-test/helpers"
	"github.com/ezzycreative1/majoo-test/models"

	"github.com/gin-gonic/gin"
	gorm "github.com/jinzhu/gorm"
	//"gorm.io/gorm"
)

type userRepository struct {
	Conn *gorm.DB
}

// NewUserRepository ..
func NewUserRepository(Conn *gorm.DB) user.Repository {
	return &userRepository{Conn}
}

func (ur *userRepository) GetUserById(userID string) (user models.User, err error) {
	return user, ur.Conn.Model(user).Where("id = ?", userID).First(&user).Error
}

func (ur *userRepository) GetUserIDByEmail(email string) (string, error) {
	var user models.User

	if err := ur.Conn.Model(user).Where("email = ?", email).First(&user).Error; err != nil {
		return "", err
	}
	return user.ID, nil
}

func (ur *userRepository) Insert(userDoc *models.User) (*string, error) {

	if err := ur.Conn.Create(&userDoc).Error; err != nil {
		return nil, err
	}

	return nil, nil
}

func (ur *userRepository) Update(user *models.User) error {
	if err := ur.Conn.Model(&user).Where("id = ?", user.ID).Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) Get(email string) (*models.User, error) {
	var (
		user models.User
	)
	if err := ur.Conn.Where("email = ?", email).Find(&user).Error; err != nil {
		return &user, err
	}

	return &user, nil

}

func (ur *userRepository) GetUsers() (*[]models.User, error) {
	var (
		users []models.User
	)
	if err := ur.Conn.Find(&users).Error; err != nil {
		return &users, nil
	}

	return &users, nil
}

func (ur *userRepository) DeleteUser(id string) (*models.User, error) {
	var (
		user models.User
	)
	if err := ur.Conn.Where("id = ?", id).Delete(&user).Error; err != nil {
		return &user, nil
	}

	return &user, nil
}

//DeleteUsers
func (ur *userRepository) DeleteUsers(listID []string) error {
	var (
		user models.User
	)
	if err := ur.Conn.Where("id in (?)", listID).Delete(&user).Error; err != nil {
		return err
	}

	return nil
}

// UpdateUser ..
func (ur *userRepository) UpdateUser(req requests.UpdateUser) error {
	var (
		err  error
		user models.User
	)

	updateUser := map[string]interface{}{"Name": req.Name}

	if err = ur.Conn.Model(&user).Where("id = ?", req.Id).Updates(updateUser).Error; err != nil {
		return nil
	}

	return err
}

func (ur *userRepository) CheckVerifyTokenExist(token, userID string) bool {
	var User models.User

	if err := ur.Conn.Model(&User).
		Where("id = ?", userID).
		Where("verify_token = ?", token).
		First(&User).Error; err != nil {
		fmt.Println(err.Error())
		return false
	}

	return true
}

func (ur *userRepository) GetUserForSendEmailRegister(userID, token string) (models.User, error) {
	var user models.User

	if err := ur.Conn.Model(&user).
		Where("id = ?", userID).
		Where("verify_token = ?", token).
		First(&user).Error; err != nil {
		fmt.Println(err.Error())
		return models.User{}, err
	}

	return user, nil
}

func (ur *userRepository) VerifyAccount(token, userID string) bool {

	var User models.User
	now := time.Now().Format("2006-01-02T15:04:05Z07:00")

	updates := map[string]interface{}{"verifed_at": fmt.Sprintf("%v", now), "verify_token": token, "status": "active"}

	if err := ur.Conn.Model(&User).
		Where("id = ?", userID).
		Updates(updates).Error; err != nil {
		return false
	}

	return true
}

func (ur *userRepository) UpdatePassword(req requests.ChangePasswordParam) error {
	var (
		err  error
		user models.User
	)

	updatePassword := map[string]interface{}{"password": req.NewPassword}

	if err = ur.Conn.Model(&user).Where("id = ?", req.UserID).Updates(updatePassword).Error; err != nil {
		return nil
	}

	return err
}

// ChangePhoto ..
func (ur *userRepository) ChangePhoto(userid string, c *gin.Context) error {
	var (
		err  error
		user models.User
	)

	err, filename := helpers.FileUpload(c)

	if err != nil {
		return err
	}
	updateUser := map[string]interface{}{"photo": filename}

	if err = ur.Conn.Model(&user).Where("id = ?", userid).Updates(updateUser).Error; err != nil {
		return nil
	}

	return err
}

// ChangeName ..
func (ur *userRepository) ChangeName(req requests.ChangeName) error {
	var (
		err  error
		user models.User
	)

	updateUser := map[string]interface{}{"Name": req.Name}

	if err = ur.Conn.Model(&user).Where("id = ?", req.Id).Updates(updateUser).Error; err != nil {
		return nil
	}

	return err
}

func (ur *userRepository) CreateNewAdmin(user models.User) error {

	if err := ur.Conn.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (ur *userRepository) FetchAllAdmins(offset int) ([]models.User, error) {

	var users []models.User

	if offset == 0 {
		if err := ur.Conn.Select([]string{"id", "name", "email", "phone_number"}).
			Limit(10).
			Find(&users).Error; err != nil {
			return []models.User{}, err
		}
	} else {
		if err := ur.Conn.Select([]string{"id", "name", "email", "phone_number"}).
			Offset(offset).
			Limit(10).
			Find(&users).Error; err != nil {
			return []models.User{}, err
		}
	}

	return users, nil

}

func (ur *userRepository) FetchAdminByID(ID string) (models.User, error) {

	var user models.User

	if err := ur.Conn.Select([]string{"id", "name", "email", "phone_number"}).
		Where("id = ?", ID).First(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (ur *userRepository) UpdateAdminBySuperuser(user models.User) (bool, error) {

	var u models.User

	if err := ur.Conn.Where("id = ?", user.ID).First(&u).Error; err != nil {
		return false, err
	}

	u.Fullname = user.Fullname
	u.Email = user.Email
	u.PhoneNumber = user.PhoneNumber

	if err := ur.Conn.Save(&u).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (ur *userRepository) DeleteAdminBySuperuser(ID string) (bool, error) {

	var user models.User

	if err := ur.Conn.Where("id = ?", ID).Delete(&user).Error; err != nil {
		return false, err
	}

	return true, nil
}

// GetById ..
func (ur *userRepository) GetById(ID string) (*models.User, error) {
	var (
		user models.User
	)
	if err := ur.Conn.Where("id = ?", ID).Find(&user).Error; err != nil {
		return &user, err
	}

	return &user, nil
}

func (ur *userRepository) IsUserVerified(userID string) bool {
	user, err := ur.GetById(userID)
	if err != nil {
		fmt.Println("error can't get userid")
		return false
	}
	if user.Status == "active" {
		return true
	}
	return false
}

func (ur *userRepository) ResetPassword(token string, userID string, pass string) error {
	var user = models.User{
		Password: pass,
		Status:   "active",
	}
	if err := ur.Conn.Model(&user).Where("id = ?", userID).
		Updates(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) GetPasswordUser(userID string) (string, error) {

	var user models.User

	if err := ur.Conn.Model(&models.User{}).Select("password").Where("id = ?", userID).First(&user).Error; err != nil {
		return "", err
	}

	return user.Password, nil
}
