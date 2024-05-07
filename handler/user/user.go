package user

type User struct {
}

func NewUser() *User {
	return &User{}
}

const (
	// 用于加密的盐值(自定义)
	pwdSalt    = "*#890"
	minPwdLen  = 1
	minNameLen = 1
)
