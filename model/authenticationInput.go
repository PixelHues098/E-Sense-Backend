package model

type AuthenticationInput struct {
    FirstName string `json:"firstName"`
    LastName string `json:"lastName"`
    Username string `json:"username"`
    Email string `json:"email" binding:"required"`
    Password string `json:"password" binding:"required"`
}
