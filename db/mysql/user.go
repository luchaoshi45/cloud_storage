package mysql

import (
	"fmt"
	"reflect"
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

func (u *User) SetAttrs(attrs map[string]interface{}) error {
	v := reflect.ValueOf(u).Elem()

	for attr, value := range attrs {
		field := v.FieldByName(attr)

		// 检查字段是否存在
		if !field.IsValid() {
			return &FieldNotFoundError{Field: attr}
		}

		// 检查字段是否可设置
		if !field.CanSet() {
			return &FieldNotSettableError{Field: attr}
		}

		// 将传入的值转换为字段的类型并设置
		val := reflect.ValueOf(value)
		if val.Type().ConvertibleTo(field.Type()) {
			field.Set(val.Convert(field.Type()))
		} else {
			return &FieldTypeMismatchError{
				Field:    attr,
				Expected: field.Type().String(),
				Actual:   val.Type().String(),
			}
		}
	}

	return nil
}

func (u *User) GetAttr(attr string) (interface{}, error) {
	v := reflect.ValueOf(u).Elem()
	field := v.FieldByName(attr)

	// 检查字段是否存在
	if !field.IsValid() {
		return nil, &FieldNotFoundError{Field: attr}
	}

	// 检查字段是否可取值
	if !field.CanInterface() {
		return nil, &FieldNotGettableError{Field: attr}
	}

	return field.Interface(), nil
}
