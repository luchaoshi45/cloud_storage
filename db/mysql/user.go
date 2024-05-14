package mysql

import (
	"database/sql"
	"fmt"
	"reflect"
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

func (u *User) exists() bool {
	tableName := reflect.TypeOf(*u).Name()
	// 查询是否存在表格
	query := fmt.Sprintf("SHOW TABLES LIKE '%s'", tableName)
	rows, err := mySql.Query(query)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// 如果表格存在，则输出相应的消息
	if rows.Next() {
		return true
	} else {
		return false
	}
}

func (u *User) create() {
	createTableSQL := `
		CREATE TABLE User(
			id INT(11) NOT NULL AUTO_INCREMENT COMMENT '用户名ID',
			user_name VARCHAR(64) NOT NULL DEFAULT '' COMMENT '用户名',
			user_pwd VARCHAR(256) NOT NULL DEFAULT '' COMMENT '用户encoded密码',
			email VARCHAR(64) DEFAULT NULL COMMENT '邮箱',
			phone VARCHAR(128) DEFAULT NULL COMMENT '手机号',
			email_validated TINYINT(1) DEFAULT 0 COMMENT '邮箱是否已验证',
			phone_validated TINYINT(1) DEFAULT 0 COMMENT '手机号是否已验证',
			signup_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '注册日期',
			last_active DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后活跃时间',
			profile TEXT COMMENT'用户属性',
			status INT(11) NOT NULL DEFAULT 0 COMMENT '账户状态(启用/禁用/锁定/标记删除',
			PRIMARY KEY (id),
			UNIQUE KEY idx_user_name (user_name),
			UNIQUE KEY idx_phone (phone),
			KEY idx_status (status)
		)ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;
	`
	// 执行SQL语句
	_, err := mySql.Exec(createTableSQL)
	if err != nil {
		panic(err.Error())
	}
}

func (u *User) existsCreate() {
	exists := u.exists()
	if !exists {
		u.create()
	}
}

// SignUp : 用户注册
func (u *User) SignUp() bool {
	u.existsCreate()

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
