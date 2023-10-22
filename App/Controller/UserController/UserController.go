package UserController

import (
	"Artemis/Common/Response"
	"Artemis/Service/UserService"
	"github.com/gin-gonic/gin"
)

type Controller struct {
}

var Instance = new(Controller)

func (ctrl *Controller) Example(c *gin.Context) {
	var params UserService.ExampleParam
	err := c.ShouldBindJSON(&params)
	if err != nil {
		Response.FailValidate(err, c)
		return
	}
	srv := UserService.New()
	if ex := srv.Example(params); ex.IsError() {
		Response.FailWithException(ex, c)
	} else {
		Response.Success(c)
	}
}
