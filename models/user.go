package models

type UserRegisterForm struct {
    Username string `json:"username"`
    Email string `json:"email"`
    Password string `json:"password"`
}

type UserResponse struct {
    UserId   string `json:"user_id"`
    Email string `json:"email"`
    Token string `json:"token"`
    Username string `json:"username"`
    Bio string `json:"bio"`
    Image string `json:"image, omitempty"`
}
