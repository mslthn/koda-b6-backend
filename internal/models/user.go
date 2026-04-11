package models

type User struct {
	ID             int    `json:"id" db:"id"`
	RoleID         string `json:"role_id" db:"role_id"`
	Fullname       string `json:"fullname" db:"fullname"`
	Email          string `json:"email" db:"email"`
	Password       string `json:"password" db:"password"`
	Address        string `json:"address" db:"address"`
	Phone          string `json:"phone" db:"phone"`
	ProfilePicture string `json:"profile_picture" db:"picture"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserResponse struct {
	Token string `json:"token"`
}

type CreateUserRequest struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserRequest struct {
	Fullname       string `json:"fullname"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	Address        string `json:"address"`
	Phone          string `json:"phone"`
	ProfilePicture string `json:"profile_picture"`
}
