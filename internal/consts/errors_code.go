package consts

import "github.com/gogf/gf/v2/errors/gcode"

// defining user error codes
var (
	CodeUserNotFound      = gcode.New(1001, "user not found", nil)
	CodePasswordInvalid   = gcode.New(1002, "invalid password", nil)
	CodeUserNotMultiLogin = gcode.New(1003, "user not allow multiLogin", nil)
	CodeUserExists        = gcode.New(1004, "user exists", nil)

	CodeRoelExists = gcode.New(1011, "roel exists", nil)

	CodeMenuExists = gcode.New(1021, "menu exists", nil)

	CodeApiExists = gcode.New(1031, "api exists", nil)

	CodeTokenInvalid  = gcode.New(1041, "token invalid", nil)
	CodeTokenMiss     = gcode.New(1042, "miss token", nil)
	CodeGenerateToken = gcode.New(1043, "generate token failed", nil)
	CodePathInvalid   = gcode.New(1044, "api path invalid", nil)

	CodeRedisError = gcode.New(2001, "redis error", nil)
)
