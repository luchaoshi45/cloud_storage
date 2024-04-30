package file

import (
	"errors"
	"os"
	"sync"
)

type AbstractFileMeta interface {
	GetSha1() string
	SetSha1(sha1 string)
	GetName() string
	SetName(name string)
	GetDir() string
	SetDir(dir string)
	GetSize() int64
	SetSize(size int64)
	GetLocation() string
	SetLocation(location string)
	GetUploadAt() string
	SetUploadAt(uploadAt string)
}

// NewFileMeta New FileMeta 结构体
func NewFileMeta() AbstractFileMeta {
	return new(FileMeta)
}

// fileMetaDict 单例
var fileMetaDict map[string]AbstractFileMeta
var once sync.Once

// GetFileMetaDict 得到/初始化 GetfileMetaDict
func GetFileMetaDict() map[string]AbstractFileMeta {
	once.Do(func() {
		fileMetaDict = make(map[string]AbstractFileMeta)
	})
	return fileMetaDict
}

// UpdateFileMetaDict 更新/新增 FileMeta
func UpdateFileMetaDict(fileMate AbstractFileMeta) {
	fileMetaDict[fileMate.GetSha1()] = fileMate
}

func checkKeyExist(sha1 string) error {
	// 检查键是否存在
	_, exists := fileMetaDict[sha1]
	if !exists {
		return errors.New("键不存在")
	}
	return nil
}

func GetFileMeta(sha1 string) (AbstractFileMeta, error) {
	// checkKeyExist
	err := checkKeyExist(sha1)
	if err != nil {
		return nil, err
	}
	return fileMetaDict[sha1], nil
}

func RemoveFileMeta(sha1 string) error {
	// checkKeyExist
	err := checkKeyExist(sha1)
	if err != nil {
		return err
	}
	// 删除键值对
	delete(fileMetaDict, sha1)
	return nil
}

func SafeRename(oldpath, newpath string) error {
	err := os.Rename(oldpath, newpath)
	return err
}

func SafeRemove(path string) error {
	err := os.Remove(path)
	return err
}
