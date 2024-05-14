package mysql

import (
	"fmt"
	"reflect"
	"time"
)

// UserFile : 用户文件表
type UserFile struct {
	ID         int    `json:"id"`
	UserName   string `json:"user_name"`
	FileSHA1   string `json:"file_sha1"`
	FileSize   int    `json:"file_size"`
	FileName   string `json:"file_name"`
	UploadAt   string `json:"upload_at"`
	LastUpdate string `json:"last_update"`
	Status     int    `json:"status"`
}

func NewUserFile() *UserFile {
	return &UserFile{}
}

func (uf *UserFile) exists() bool {
	tableName := reflect.TypeOf(*uf).Name()
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

func (uf *UserFile) create() {
	createTableSQL := `
		CREATE TABLE UserFile (
			id INT(11) NOT NULL PRIMARY KEY AUTO_INCREMENT,
			user_name VARCHAR(64) NOT NULL COMMENT '用户名',
			file_sha1 VARCHAR(64) NOT NULL DEFAULT '' COMMENT '文件哈希',
			file_size INT(11) DEFAULT 0 COMMENT '文件大小',
			file_name VARCHAR(256) NOT NULL DEFAULT '' COMMENT '文件名',
			upload_at DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '上传时间',
			last_update DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
			status INT(11) NOT NULL DEFAULT 0 COMMENT '文件状态(0正常1已删除2禁用)',
			UNIQUE KEY idx_user_file (user_name, file_sha1),
			KEY idx_status (status),
			KEY idx_user_id (user_name)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
	`
	// 执行SQL语句
	_, err := mySql.Exec(createTableSQL)
	if err != nil {
		panic(err.Error())
	}
}

func (uf *UserFile) existsCreate() {
	exists := uf.exists()
	if !exists {
		uf.create()
	}
}

func (uf *UserFile) Insert(username, filehash, filename string, filesize int64) bool {
	uf.existsCreate()
	stmt, err := mySql.Prepare(
		"insert ignore into UserFile (`user_name`,`file_sha1`,`file_name`, `file_size`,`upload_at`) values (?,?,?,?,?)")
	if err != nil {
		return false
	}
	defer stmt.Close()

	_, err = stmt.Exec(username, filehash, filename, filesize, time.Now())
	if err != nil {
		return false
	}
	return true
}
