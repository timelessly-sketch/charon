package consts

import "github.com/gogf/gf/v2/errors/gcode"

// defining user error codes
var (
	CodeUserNotFound      = gcode.New(1001, "user not found", nil)
	CodePasswordInvalid   = gcode.New(1002, "invalid password", nil)
	CodeUserNotMultiLogin = gcode.New(1003, "user not allow multiLogin", nil)
	CodeUserDisabled      = gcode.New(1004, "user disabled", nil)
	CodeUserExists        = gcode.New(1005, "user exists", nil)
	CodeGenerateToken     = gcode.New(1006, "generate token failed", nil)
	CodeTokenInvalid      = gcode.New(1007, "token invalid", nil)
	CodeMenuExists        = gcode.New(1008, "menu exists", nil)

	CodeRedisError = gcode.New(2001, "redis error", nil)
)
