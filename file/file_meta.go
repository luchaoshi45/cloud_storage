package file

type FileMeta struct {
	sha1     string
	name     string
	size     int64
	location string
	uploadAt string
}

func (fm *FileMeta) GetSha1() string {
	return fm.sha1
}

func (fm *FileMeta) SetSha1(sha1 string) {
	fm.sha1 = sha1
}

func (fm *FileMeta) GetName() string {
	return fm.name
}

func (fm *FileMeta) SetName(name string) {
	fm.name = name
}

func (fm *FileMeta) GetSize() int64 {
	return fm.size
}

func (fm *FileMeta) SetSize(size int64) {
	fm.size = size
}

func (fm *FileMeta) GetLocation() string {
	return fm.location
}

func (fm *FileMeta) SetLocation(location string) {
	fm.location = location
}

func (fm *FileMeta) GetUploadAt() string {
	return fm.uploadAt
}

func (fm *FileMeta) SetUploadAt(uploadAt string) {
	fm.uploadAt = uploadAt
}
