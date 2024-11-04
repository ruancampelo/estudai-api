package api

import (
	"estudai-api/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterContentRoutes(router *gin.RouterGroup, contentService service.ContentService) {

	router.GET("/all", func(c *gin.Context) {
		user, err := contentService.GetAllContent()
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
			return
		}
		c.JSON(http.StatusOK, user)
	})

	router.POST("/upload", func(c *gin.Context) {
		// Obtendo o arquivo do formulário
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Arquivo não enviado"})
			return
		}

		err = contentService.CreateContent(file)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}
		// Retornando uma resposta de sucesso
		c.JSON(http.StatusNoContent, gin.H{})
	})

}
