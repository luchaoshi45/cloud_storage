package mysql

import "reflect"

// User : 用户表model
type User struct {
	Id           int64
	Username     string
	Email        string
	Phone        string
	SignupAt     string
	LastActiveAt string
	Status       int
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
