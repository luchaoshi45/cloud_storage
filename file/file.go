package file

import "sync"

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

func GetFileMeta(sha1 string) AbstractFileMeta {
	return fileMetaDict[sha1]
}
