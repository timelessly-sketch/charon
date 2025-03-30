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
	*/
	CodeUserNotFound      = gcode.New(10211, "user not found", nil)
	CodeUserExists        = gcode.New(10212, "user exists", nil)
	CodeUserNotMultiLogin = gcode.New(10213, "user not allow multiLogin", nil)
	CodePasswordInvalid   = gcode.New(10214, "invalid password", nil)
	CodeResetPassError    = gcode.New(10005, "reset password error", nil)
	CodeMenuExists        = gcode.New(10221, "menu exists", nil)
	CodeApiExists         = gcode.New(10231, "api exists", nil)
)
