package common

// ErrorCode int32
type ErrorCode int32

const (
	_ int32 = iota + 9999
	// StatusOK : 正常
	StatusOK
	// StatusParamInvalid : 请求参数无效
	StatusParamInvalid
	// StatusServerError : 服务出错
	StatusServerError
	// StatusMysqlDBError ： mysql数据库出错
	StatusMysqlDBError
	// StatusMysqlDBOK : mysql查询成功
	StatusMysqlDBOK
	// StatusMysqlDBNoChange : 数据库无变动
	StatusMysqlDBNoChange
	// StatusRegisterFailed : 注册失败
	StatusRegisterFailed
	// StatusUserSigninFailed	: 登录失败
	StatusUserSigninFailed
	// StatusTokenInvalid : token无效
	StatusTokenInvalid
	// StatusGetTokenFailed: 获取token失败
	StatusGetTokenFailed
	// StatusSessionExpiration: 会话过期(被抢占)
	StatusSessionExpiration
)
