package mysql

import (
	"fmt"
	"time"
)

// User : 用户表model
type User struct {
	ID             int       `json:"id"`
	UserName       string    `json:"user_name"`
	UserPwd        string    `json:"user_pwd"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	EmailValidated bool      `json:"email_validated"`
	PhoneValidated bool      `json:"phone_validated"`
	SignupAt       time.Time `json:"signup_at"`
	LastActive     time.Time `json:"last_active"`
	Profile        string    `json:"profile"`
	Status         int       `json:"status"`
}

func NewUser(UserName, UserPwd string) *User {
	return &User{UserName: UserName, UserPwd: UserPwd}
}

// Signup : 用户登录
func (u *User) Signup() bool {
	stmt, err := mySql.Prepare("insert ignore into User (`user_name`,`user_pwd`) values (?,?)")
	if err != nil {
		fmt.Println("Failed to insert, err:" + err.Error())
		return false
	}
	defer stmt.Close()

	ret, err := stmt.Exec(u.UserName, u.UserPwd)
	if err != nil {
		fmt.Println("Failed to insert, err:" + err.Error())
		return false
	}

	if rowsAffected, err := ret.RowsAffected(); err == nil && rowsAffected > 0 {
		return true
	}

	return false
}
