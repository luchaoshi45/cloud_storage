package mysql

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

// SetAttrs 方法用于根据字典设置 UserFile 结构体的属性
func (uf *UserFile) SetAttrs(attrs map[string]interface{}) {
	for key, value := range attrs {
		switch key {
		case "Sha1":
			uf.Sha1 = value.(string)
		case "Name":
			uf.Name = value.(string)
		case "Dir":
			uf.Dir = value.(string)
		case "Size":
			uf.Size = value.(int64)
		case "UploadAt":
			uf.UploadAt = value.(string)
		case "UserID":
			uf.UserID = value.(int64)
		}
	}
}

func (uf *UserFile) GetLocation() string {
	return uf.Dir + uf.Name
}
