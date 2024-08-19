package errs

// User 相关
const (
	// UserInternalServerError 统一的用户模块的系统错误
	UserInternalServerError = 401000
	// UserInvalidInput 统一的用户模块的输入错误
	UserInvalidInput = 401001
	// UserInvalidOrPassword 用户名错误或者密码不对
	UserInvalidOrPassword = 401002
	// UserDuplicateEmail 用户邮箱冲突
	UserDuplicateEmail = 401003
)
