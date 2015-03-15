package models

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	ID         string    `json:"id" gorm:"column:id;primary_key" xml:"id"`
	Email      string    `json:"email" gorm:"column:email" xml:"email"`
	Password   string    `json:"password" gorm:"column:password" xml:"password"`
	Name       string    `json:"name" gorm:"column:name" xml:"name"`
	IsVerified bool      `json:"is_verified" gorm:"column:is_verified" xml:"is_verified"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at" xml:"created_at"`
	Subscribed bool      `json:"subscribed" gorm:"column:subscribed" xml:"subscribed"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"column:updated_at" xml:"updated_at"`
	AuthToken  string    `json:"auth_token" gorm:"column:auth_token" xml:"auth_token"`
}

func (u User) TableName() string {
	return "Users"
}

func FindUser(uuid string, orm gorm.DB) *User {
	users := []User{}
	orm.Where("id = ?", uuid).Find(&users)

	if len(users) > 0 {
		return &users[0]
	}

	return nil
}

func FindUsers(orm gorm.DB) []User {

	users := []User{}
	orm.Find(&users)

	if len(users) > 0 {
		return users
	}

	return nil
}

func CreateUser(name, email, password string, orm gorm.DB) *User {
	user := User{
		ID:        getUUID(),
		Name:      name,
		Email:     email,
		Password:  superHashingSecureFunction(password),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now()}

	orm.Create(&user)
	return &user
}

// utility function to generate uuid string
func getUUID() string {
	f, _ := os.Open("/dev/urandom")
	b := make([]byte, 16)
	f.Read(b)
	f.Close()
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid
}

// not really secure
func superHashingSecureFunction(input string) string {
	h := sha1.New()
	h.Write([]byte(input))
	return hex.EncodeToString(h.Sum(nil))
}
