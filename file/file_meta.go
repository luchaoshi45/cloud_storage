package file

import "path"

type FileMeta struct {
	Sha1     string `json:"sha1"`
	Name     string `json:"name"`
	Dir      string `json:"dir"`
	Size     int64  `json:"size"`
	UploadAt string `json:"uploadAt"`
}

func (fm *FileMeta) GetSha1() string {
	return fm.Sha1
}

func (fm *FileMeta) SetSha1(sha1 string) {
	fm.Sha1 = sha1
}

func (fm *FileMeta) GetName() string {
	return fm.Name
}

func (fm *FileMeta) SetName(name string) {
	fm.Name = name
}

func (fm *FileMeta) GetDir() string {
	return fm.Dir
}

func (fm *FileMeta) SetDir(dir string) {
	fm.Dir = dir
}

func (fm *FileMeta) GetSize() int64 {
	return fm.Size
}

func (fm *FileMeta) SetSize(size int64) {
	fm.Size = size
}

func (fm *FileMeta) GetLocation() string {
	return path.Join(fm.GetDir(), fm.GetName())
}

func (fm *FileMeta) SetLocation(location string) {
	fm.SetDir(path.Dir(location))
	fm.SetName(path.Base(location))
}

func (fm *FileMeta) GetUploadAt() string {
	return fm.UploadAt
}

func (fm *FileMeta) SetUploadAt(uploadAt string) {
	fm.UploadAt = uploadAt
}
