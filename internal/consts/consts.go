package consts

import "fmt"

var (
	pathMethod = "%s:%s"
)

func BuildPathMethod(path, method string) string {
	return fmt.Sprintf(pathMethod, path, method)
}
