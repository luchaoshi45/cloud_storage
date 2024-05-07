package mysql

import "fmt"

// UserToken : 用户 Token
type UserToken struct {
	ID       int    `json:"id"`
	UserName string `json:"user_name"`
	Token    string `json:"token"`
}

func NewUserToken(UserName, Token string) *UserToken {
	return &UserToken{UserName: UserName, Token: Token}
}

func (ut *UserToken) UpdateToken() bool {
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
