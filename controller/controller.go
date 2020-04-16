package controller

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// Controller example
type Controller struct {
}

// NewController example
func NewController() *Controller {
	return &Controller{}
}

// Message example
type Message struct {
	Message string `json:"message" example:"message"`
}

func GinHandler(r *gin.Engine) *gin.Engine {
	c := NewController()

	v1 := r.Group("/api/v1")
	{
		users := v1.Group("/users")
		{
			users.GET(":id", c.ShowUser)
			users.GET("", c.ListUsers)
			users.POST("", c.AddUser)
			users.DELETE(":id", c.DeleteUser)
			users.PATCH(":id", c.UpdateUser)
			users.POST(":id/images", c.UploadUserImage)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r;
}
