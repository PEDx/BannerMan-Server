package errno

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "内部错误"}
	ErrBind             = &Errno{Code: 10002, Message: "参数不合法"}

	ErrValidation = &Errno{Code: 20001, Message: "Validation failed."}
	ErrDatabase   = &Errno{Code: 20002, Message: "Database error."}
	ErrToken      = &Errno{Code: 20003, Message: "Error occurred while signing the JSON web token."}

	// user errors
	ErrEncrypt            = &Errno{Code: 20101, Message: "Error occurred while encrypting the user password."}
	ErrUserNotFound       = &Errno{Code: 20102, Message: "未找到用户"}
	ErrUserAlreadyExisted = &Errno{Code: 20103, Message: "用户名已存在"}
	ErrTokenInvalid       = &Errno{Code: 20104, Message: "token 不合法"}
	ErrPasswordIncorrect  = &Errno{Code: 20105, Message: "密码错误"}
	// page errors
	ErrPageNotFound  = &Errno{Code: 20201, Message: "页面未找到"}
	ErrGetWidget     = &Errno{Code: 20202, Message: "组件获取出错"}
	ErrGetWidgetLink = &Errno{Code: 20203, Message: "组件获取链接出错"}
	ErrGetWidgetData = &Errno{Code: 20204, Message: "组件数据出错"}
	ErrBuildNetwork  = &Errno{Code: 20205, Message: "构建系统网络错误"}
	ErrBuildResponse = &Errno{Code: 20206, Message: "构建系统数据错误"}
)
