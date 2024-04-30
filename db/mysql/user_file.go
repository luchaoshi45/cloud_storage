package mysql

import "reflect"

type UserFile struct {
	Sha1     string `json:"sha1"`
	Name     string `json:"name"`
	Dir      string `json:"dir"`
	Size     int64  `json:"size"`
	UploadAt string `json:"uploadAt"`
	UserId   int64  `json:"userId"`
}

func NewUserFile() *UserFile {
	return &UserFile{}
}

// Insert : 更新用户文件表
func (uf *UserFile) Insert() bool {
	stmt, err := mySql.Prepare("insert into UserFile (`user_id`, `Sha1`, `Name`, `Dir`,`Size`, `upload_at`) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return false
	}
	defer stmt.Close()

	_, err = stmt.Exec(uf.UserId, uf.Sha1, uf.Name, uf.Dir, uf.Size, uf.UploadAt)
	if err != nil {
		return false
	}
	return true
}

func (uf *UserFile) SetAttrs(attrs map[string]interface{}) error {
	v := reflect.ValueOf(uf).Elem()

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

func (uf *UserFile) GetAttr(attr string) (interface{}, error) {
	v := reflect.ValueOf(uf).Elem()
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
