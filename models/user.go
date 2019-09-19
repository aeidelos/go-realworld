package models

type User struct {
    UserId   string `json:"user_id"`
    email    string `json:"email"`
    token    string `json:"token"`
    username string `json:"username"`
    bio      string `json:"bio"`
    image    string `json:"image"`
}
