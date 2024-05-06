package mysql

import (
	"reflect"
)

type UserFile struct {
	Sha1     string `json:"sha1"`
	Name     string `json:"name"`
	Dir      string `json:"dir"`
	Size     int64  `json:"size"`
	UploadAt string `json:"uploadAt"`
	UserID   int64  `json:"userId"`
}

func NewUserFile() *UserFile {
	return &UserFile{}
}

// Insert : 插入用户文件表
func (uf *UserFile) Insert() bool {
	stmt, err := mySql.Prepare("insert into UserFile (`user_id`, `sha1`, `name`, `dir`,`size`, `upload_at`) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return false
	}
	defer stmt.Close()

	_, err = stmt.Exec(uf.UserID, uf.Sha1, uf.Name, uf.Dir, uf.Size, uf.UploadAt)
	if err != nil {
		return false
	}
	return true
}

// Query : 查询用户文件表
func (uf *UserFile) Query(sha1 string) (*UserFile, error) {
	err := mySql.QueryRow("SELECT `sha1`, `name`, `dir`, `size`, `upload_at`, `user_id` FROM UserFile WHERE `sha1` = ?",
		sha1).Scan(&uf.Sha1, &uf.Name, &uf.Dir, &uf.Size, &uf.UploadAt, &uf.UserID)
	if err != nil {
		return nil, err
	}
	return uf, nil
}

// Update : 更新用户文件表
func (uf *UserFile) Update(sha1 string, newName string) (*UserFile, error) {
	uf.Name = newName
	// 准备 SQL 更新语句，更新指定 sha1 的记录的 Name 字段为 newName
	stmtUpdate, err := mySql.Prepare("UPDATE UserFile SET `name` = ? WHERE `sha1` = ?")
	if err != nil {
		return nil, err
	}
	defer stmtUpdate.Close()

	// 执行更新操作
	_, err = stmtUpdate.Exec(newName, sha1)
	if err != nil {
		return nil, err
	}
	return uf, nil
}

// Delete : 删除用户文件表
func (uf *UserFile) Delete(sha1 string) (*UserFile, error) {
	// 准备 SQL 更新语句，更新指定 sha1 的记录的 Name 字段为 newName
	stmtUpdate, err := mySql.Prepare("DELETE FROM UserFile WHERE `sha1` = ?")
	if err != nil {
		return nil, err
	}
	defer stmtUpdate.Close()

	// 执行更新操作
	_, err = stmtUpdate.Exec(sha1)
	if err != nil {
		return nil, err
	}
	return uf, nil
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

func (uf *UserFile) GetLocation() string {
	return uf.Dir + uf.Name
}
