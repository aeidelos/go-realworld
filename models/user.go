package models

import "database/sql"

type UserRegisterForm struct {
    Username string `json:"username"`
    Email    string `json:"email"`
    Password string `json:"password"`
}

type UserLogin struct {
    Username string `json:"username"`
    Password string `json:"password"`
    Token    string `json:"token"`
}

type UserResponse struct {
    UserId   string         `json:"user_id" db:"user_id"`
    Email    string         `json:"email" db:"email"`
    Token    string         `json:"token" db:"token"`
    Username string         `json:"username" db:"username"`
    Bio      string         `json:"bio" db:"bio"`
    Image    sql.NullString `json:"image" db:"image"`
}
