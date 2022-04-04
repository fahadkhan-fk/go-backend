package schema

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

//User struct stores user
type User struct {
	ID        uuid.UUID  `json:"id"`
	UserName  string     `json:"user_name"`
	Password  string     `json:"password"`
	Email     string     `json:"email"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func AddNewUser(client *gorm.DB, user *User) (*User, *gorm.DB) {
	if err := client.Create(&user); err.Error != nil {
		return nil, err
	} else {
		return user, nil
	}
}

func FetchAllUsers(client *gorm.DB, users *[]User) (*[]User, *gorm.DB) {
	if err := client.Find(&users); err.Error != nil {
		return nil, err
	} else {
		return users, nil
	}
}

func GetUserByID(client *gorm.DB, user *User, userID string) (*User, *gorm.DB) {
	if err := client.Where("id = ?", userID).First(&user); err.Error != nil {
		return nil, err
	} else {
		return user, nil
	}
}

func UpdateUser(client *gorm.DB, user *User) {
	client.Save(&user)
}

func DeleteUser(client *gorm.DB, user *User) {
	client.Delete(&user)
}
