package public

import "github.com/gogf/gf/v2/frame/g"

type VerifyUnique struct {
	Id    int
	Where g.Map
}

type Pagination struct {
	Page int `json:"page" d:"1"`
	Size int `json:"size" d:"10"`
}
