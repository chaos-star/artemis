package UserController

type UserReq struct {
	Name string `json:"name" binding:"required"`
}
