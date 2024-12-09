package auth

import (
	"VOU-Server/cmd/gateway/config"

	"github.com/gin-gonic/gin"
)

func (authSvc *AuthServiceClient) InitRoutes(r *gin.Engine, config config.Config) {
	routes := r.Group("/auth")
	routes.POST("/register", authSvc.Register)

	r.Run(config.HTTPServerAddress)
}
