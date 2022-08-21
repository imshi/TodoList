package serializer

type User struct {
	ID       uint   `json:"id" form:"id" example:"1"`
	UserName string `json:"userName" form:"userName" example:"FanOne"`
	CreateAt int64  `json:"createAt" form:"createAt"`
}
