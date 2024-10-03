package models

type RegisterData struct { // for swagger
	FullName        string `json:"full_name"`
	Email           string `json:"email"`
	PhoneNumber     string `json:"phone_number"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type RegisterReq struct {
	FullName       string
	Email          string
	PhoneNumber    string
	HashedPassword string
}

type RegisterResp struct {
	ID string `json:"id"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type UserDetails struct {
	ID             string
	HashedPassword string
	Role           string
}

type ResetPassReq struct {
	Email       string `json:"email"`
	Code        string `json:"code"`
	NewPassword string `json:"new_password"`
}

type RefreshToken struct {
	RefreshToken string `json:"refresh_token"`
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
}
