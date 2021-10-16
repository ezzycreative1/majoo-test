package requests

// RegisterRequest ..
type RegisterRequest struct {
	Fullname    string `form:"fullname" json:"fullname" validate:"required"`
	Email       string `form:"email" json:"email" validate:"required,email"`
	Password    string `form:"password" json:"password" validate:"required"`
	PhoneNumber string `form:"phone_number" json:"phone_number" validate:"required"`
}

type EditRequest struct {
	ID          int64  `json:"id"`
	Fullname    string `json:"fullname"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
}

// LoginRequest ..
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// CreateUser ..
type CreateUser struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	NoHp   string `json:"noHp"`
	RoleId int    `json:"roleId"`
}

// UpdateUser ..
type UpdateUser struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	NoHp   string `json:"noHp"`
	RoleId int    `json:"roleId"`
}

// ChangePhoto ..
type ChangePhoto struct {
	Id    string `json:"id"`
	Photo string `json:"photo"`
}

// ChangeName ..
type ChangeName struct {
	Id   string `json:"Id"`
	Name string `json:"name"`
}
type ForgotPassword struct {
	Email string `json:"email"`
}

// ChangePassword ..
type ChangePasswordRequest struct {
	OldPassword     string `json:"old_password"`
	NewPassword     string `json:"new_password"`
	ConfirmPassword string `json:"confirm_password"`
}

type ChangePasswordParam struct {
	UserID          string `json:"user_id"`
	OldPassword     string `json:"oldPassword"`
	NewPassword     string `json:"newPassword"`
	ConfirmPassword string `json:"confirmPassword"`
}

type ResetPassword struct {
	Token           string `json:"token"`
	NewPassword     string `json:"newPassword"`
	ConfirmPassword string `json:"confirmPassword"`
}
