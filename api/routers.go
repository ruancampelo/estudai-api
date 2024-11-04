package api

import (
	. "estudai-api/internal/infrastructure/dependency"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, deps *Dependencies) {
	// Rotas de usuário
	//userRoutes := router.Group("/users")
	//RegisterUserRoutes(userRoutes, deps.UserService)

	// Rotas de produto
	contentRoutes := router.Group("/contents")
	RegisterContentRoutes(contentRoutes, deps.ContentService)

}
