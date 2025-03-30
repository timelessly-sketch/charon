package cache

import "fmt"

const (
	service   = "charon"
	userToken = "%s:%s:Token"
	roleKey   = "%s:%s:Role"
	menuKey   = "%s:%s:Menu"
)

func BuildUserToken(name string) string {
	return fmt.Sprintf(userToken, service, name)
}

func BuildRole(name string) string {
	return fmt.Sprintf(roleKey, service, name)
}

func BuildMenu(name string) string {
	return fmt.Sprintf(menuKey, service, name)
}
