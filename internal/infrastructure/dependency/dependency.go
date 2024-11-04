package dependency

import (
	"estudai-api/internal/infrastructure/repository"
	"estudai-api/internal/service"
	"gorm.io/gorm"
)

type Dependencies struct {
	ContentService service.ContentService
}

func InitDependencies(db *gorm.DB) *Dependencies {
	// Inicializar repositórios
	//userRepo := repository.NewUserRepository()
	contentRepository := repository.NewContentRepository(db)

	// Inicializar serviços com os repositórios
	//userService := service.NewUserService(userRepo)
	productService := service.NewContentService(contentRepository)

	return &Dependencies{
		//UserService:    userService,
		ContentService: productService,
	}
}
