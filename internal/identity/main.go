package identity

import "github.com/gin-gonic/gin"

func Init(r *gin.Engine) {
	r.POST("/identity/provider", ProviderLogin)
	r.POST("/auth/refresh", Refresh)
}