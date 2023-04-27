package routes

import (
	"github.com/gin-gonic/gin"
	"go1/app/http/controllers/api/v1/auth"
	"net/http"
)

func RegisterAPIRoutes(r *gin.Engine) {

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"hello": "1",
			})
		})

		authGroup := v1.Group("/auth")
		{
			sc := new(auth.SignupController)

			authGroup.POST("signup/phone/exist", sc.IsPhoneExist)
		}

	}

}
