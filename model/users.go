package model

import (
    "esense/database"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
    "html"
    "strings"
)

type User struct {
    gorm.Model
	FirstName string `gorm:"size:255;not null" json:"firstName"`
	LastName  string `gorm:"size:255;not null" json:"lastName"`
    Username string `gorm:"size:255;not null" json:"username"`
	Email     string `gorm:"size:255;not null" json:"email"`
    Password string `gorm:"size:255;not null;" json:"-"`
}

func (user *User) Save() (*User, error) {
    err := database.Database.Create(&user).Error
    if err != nil {
        return &User{}, err
    }
    return user, nil
}

func (user *User) BeforeSave(*gorm.DB) error {
    passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.Password = string(passwordHash)
    user.Username = html.EscapeString(strings.TrimSpace(user.Username))
    return nil
}

func (user *User) ValidatePassword(password string) error {
    return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func FindUserByEmail(email string) (User, error) {
    var user User
    err := database.Database.Where("email=?", email).Find(&user).Error
    if err != nil {
        return User{}, err
    }
    return user, nil
}

func FindUserById(id uint) (User, error) {
    var user User
    err := database.Database.Preload("Entries").Where("ID=?", id).Find(&user).Error
    if err != nil {
        return User{}, err
    }
    return user, nil
}

