// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"charon/internal/dao/internal"
)

// apiDao is the data access object for the table tb_api.
// You can define custom methods on it to extend its functionality as needed.
type apiDao struct {
	*internal.ApiDao
}

var (
	// Api is a globally accessible object for table tb_api operations.
	Api = apiDao{internal.NewApiDao()}
)

// Add your custom methods and functionality below.
