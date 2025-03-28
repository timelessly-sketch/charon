// =================================================================================
// This file is auto-generated by the GoFrame CLI tool. You may modify it as needed.
// =================================================================================

package dao

import (
	"charon/internal/dao/internal"
)

// menuDao is the data access object for the table tb_menu.
// You can define custom methods on it to extend its functionality as needed.
type menuDao struct {
	*internal.MenuDao
}

var (
	// Menu is a globally accessible object for table tb_menu operations.
	Menu = menuDao{internal.NewMenuDao()}
)

// Add your custom methods and functionality below.
