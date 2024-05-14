package mysql

import (
	"fmt"
	"reflect"
)

// UserToken : 用户 Token
type UserToken struct {
	ID       int    `json:"id"`
	UserName string `json:"user_name"`
	Token    string `json:"token"`
}

func NewUserToken(UserName, Token string) *UserToken {
	return &UserToken{UserName: UserName, Token: Token}
}

func (ut *UserToken) exists() bool {
	tableName := reflect.TypeOf(*ut).Name()
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

func (ut *UserToken) create() {
	createTableSQL := `
		CREATE TABLE UserToken(
    		id INT(11) NOT NULL AUTO_INCREMENT,
     		user_name VARCHAR(64) NOT NULL DEFAULT '' COMMENT '用户名',
    	 	token CHAR(40) NOT NULL DEFAULT '' COMMENT '用户登录token',
     		PRIMARY KEY (id),
    		UNIQUE KEY idx_user_name (user_name)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
	`
	// 执行SQL语句
	_, err := mySql.Exec(createTableSQL)
	if err != nil {
		panic(err.Error())
	}
}

func (ut *UserToken) existsCreate() {
	exists := ut.exists()
	if !exists {
		ut.create()
	}
}
func (ut *UserToken) UpdateToken() bool {
	ut.existsCreate()
	stmt, err := mySql.Prepare("replace into UserToken (`user_name`,`token`) values (?,?)")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer stmt.Close()

	_, err = stmt.Exec(ut.UserName, ut.Token)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}
