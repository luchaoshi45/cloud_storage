package mysql

import (
	"fmt"
	"reflect"
)

type File struct {
	Sha1     string `json:"sha1"`
	Name     string `json:"name"`
	Dir      string `json:"dir"`
	Size     int64  `json:"size"`
	UploadAt string `json:"uploadAt"`
}

func NewFile() *File {
	return &File{}
}

func (f *File) exists() bool {
	tableName := reflect.TypeOf(*f).Name()
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

func (f *File) create() {
	createTableSQL := `
		CREATE TABLE File (
			sha1 VARCHAR(255) PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			dir VARCHAR(255) NOT NULL,
			size BIGINT NOT NULL,
			upload_at DATETIME NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
		);
	`
	// 执行SQL语句
	_, err := mySql.Exec(createTableSQL)
	if err != nil {
		panic(err.Error())
	}
}

func (f *File) existsCreate() {
	exists := f.exists()
	if !exists {
		f.create()
	}
}

// Insert : 插入用户文件表
func (f *File) Insert() bool {
	f.existsCreate()
	stmt, err := mySql.Prepare("insert into File (`sha1`, `name`, `dir`,`size`, `upload_at`) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return false
	}
	defer stmt.Close()

	_, err = stmt.Exec(f.Sha1, f.Name, f.Dir, f.Size, f.UploadAt)
	if err != nil {
		return false
	}
	return true
}

// Query : 查询用户文件表
func (f *File) Query(sha1 string) (*File, error) {
	err := mySql.QueryRow("SELECT `sha1`, `name`, `dir`, `size`, `upload_at` FROM File WHERE `sha1` = ?",
		sha1).Scan(&f.Sha1, &f.Name, &f.Dir, &f.Size, &f.UploadAt)
	if err != nil {
		return nil, err
	}
	return f, nil
}

// Update : 更新用户文件表
func (f *File) Update(sha1 string, newName string) (*File, error) {
	f.Name = newName
	// 准备 SQL 更新语句，更新指定 sha1 的记录的 Name 字段为 newName
	stmtUpdate, err := mySql.Prepare("UPDATE File SET `name` = ? WHERE `sha1` = ?")
	if err != nil {
		return nil, err
	}
	defer stmtUpdate.Close()

	// 执行更新操作
	_, err = stmtUpdate.Exec(newName, sha1)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (f *File) UpdateDir(sha1 string, dir string) (*File, error) {
	// 准备 SQL 更新语句，更新指定 sha1 的记录的 Name 字段为 newName
	stmtUpdate, err := mySql.Prepare("UPDATE File SET `dir` = ? WHERE `sha1` = ?")
	if err != nil {
		return nil, err
	}
	defer stmtUpdate.Close()

	// 执行更新操作
	_, err = stmtUpdate.Exec(dir, sha1)
	if err != nil {
		return nil, err
	}
	return f, nil
}

// Delete : 删除用户文件表
func (f *File) Delete(sha1 string) (*File, error) {
	// 准备 SQL 更新语句，更新指定 sha1 的记录的 Name 字段为 newName
	stmtUpdate, err := mySql.Prepare("DELETE FROM File WHERE `sha1` = ?")
	if err != nil {
		return nil, err
	}
	defer stmtUpdate.Close()

	// 执行更新操作
	_, err = stmtUpdate.Exec(sha1)
	if err != nil {
		return nil, err
	}
	return f, nil
}

// SetAttrs 方法用于根据字典设置 File 结构体的属性
func (f *File) SetAttrs(attrs map[string]interface{}) {
	for key, value := range attrs {
		switch key {
		case "Sha1":
			f.Sha1 = value.(string)
		case "Name":
			f.Name = value.(string)
		case "Dir":
			f.Dir = value.(string)
		case "Size":
			f.Size = value.(int64)
		case "UploadAt":
			f.UploadAt = value.(string)
		}
	}
}

func (f *File) GetLocation() string {
	return f.Dir + f.Name
}
