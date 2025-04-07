package consts

import (
	"github.com/gogf/gf/v2/errors/gcode"
)

var (
	/*
		系统错误码 db、redis等 10000 ~ 10100
	*/
	CodeDbOperationError    = gcode.New(10001, "Database Operation Error", nil)
	CodeRedisOperationError = gcode.New(10002, "Redis Operation Error", nil)
	CodeGenerateTokenError  = gcode.New(10003, "Generate Token Operation Error", nil)
	CodeCleanupTokenError   = gcode.New(10004, "Cleanup Token Operation Error", nil)
	CodeTokenMissError      = gcode.New(10005, "Token Miss Error", nil)
	CodeTokenInvalidError   = gcode.New(10006, "Token Invalid", nil)

	/*
		业务系统错误码 10200 ~ 10300
		用户10210
		菜单10220
		接口10230
		配置10240
		集群10250
	*/
	CodeUserNotFound      = gcode.New(10211, "用户不存在", nil)
	CodeUserNotMultiLogin = gcode.New(10213, "user not allow multiLogin", nil)
	CodePasswordInvalid   = gcode.New(10214, "invalid password", nil)
	CodeResetPassError    = gcode.New(10005, "重置用户密码失败", nil)
	CodeMenuExists        = gcode.New(10221, "菜单已经存在", nil)
	CodeClusterNotFound   = gcode.New(10250, "集群不存在", nil)
	CodeClusterInitError  = gcode.New(10251, "集群测试失败,请检查参数", nil)
	CodeServiceExists     = gcode.New(10252, "该环境服务已存在", nil)
)
