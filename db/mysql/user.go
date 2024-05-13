package mysql

import (
	"database/sql"
	"fmt"
)

// User : 用户表model
type User struct {
	ID             int    `json:"id"`
	UserName       string `json:"user_name"`
	UserPwd        string `json:"user_pwd"`
	Email          string `json:"email"`
	Phone          string `json:"phone"`
	EmailValidated bool   `json:"email_validated"`
	PhoneValidated bool   `json:"phone_validated"`
	SignupAt       string `json:"signup_at"`
	LastActive     string `json:"last_active"`
	Profile        string `json:"profile"`
	Status         int    `json:"status"`
}

func NewUser(UserName, UserPwd string) *User {
	return &User{UserName: UserName, UserPwd: UserPwd}
}

// SignUp : 用户注册
func (u *User) SignUp() bool {
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

// SignIn : Check if the provided username and encrypted password match
func (u *User) SignIn() bool {
	// Prepare the SQL statement
	stmt, err := mySql.Prepare("SELECT user_pwd FROM User WHERE user_name=? LIMIT 1")
	if err != nil {
		fmt.Println("Error preparing SQL statement:", err.Error())
		return false
	}
	defer stmt.Close()

	// Query the database for the user's password
	row := stmt.QueryRow(u.UserName)

	var dbPwd string
	// Scan the password from the row
	err = row.Scan(&dbPwd)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Username not found:", u.UserName)
			return false
		}
		fmt.Println("Error retrieving password from database:", err.Error())
		return false
	}

	// Compare the encrypted passwords
	if dbPwd == u.UserPwd {
		return true
	}
	return false
}

// GetUserInfo : 查询用户信息
func (u *User) GetUserInfo(username string) (User, error) {
	user := User{}

	stmt, err := mySql.Prepare(
		"select user_name,signup_at from User where user_name=? limit 1")
	if err != nil {
		fmt.Println(err.Error())
		return user, err
	}
	defer stmt.Close()

	// 执行查询的操作
	err = stmt.QueryRow(username).Scan(&user.UserName, &user.SignupAt)
	if err != nil {
		return user, err
	}
	return user, nil
}
