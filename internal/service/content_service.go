package service

import (
	. "estudai-api/internal/infrastructure/client"
	. "estudai-api/internal/infrastructure/repository"
	. "estudai-api/internal/model"
	"github.com/goccy/go-json"
	"log"
	"mime/multipart"
)

type ContentService interface {
	CreateContent(file *multipart.FileHeader) error
	GetContentByID(id uint) (*Content, error)
	GetAllContent() ([]Content, error)
}

type ContentServiceImpl struct {
	repository IRepository[Content]
}

func NewContentService(repo IRepository[Content]) ContentService {
	return &ContentServiceImpl{repository: repo}
}

func (s *ContentServiceImpl) CreateContent(file *multipart.FileHeader) error {
	var result = GetQuestionsByFile(file)

	quetionsJSONBytes, err := json.Marshal(result.Conteudo)

	if err != nil {
		log.Fatalf("Erro ao converter para JSON: %v", err)
	}
	var contentItem = Content{Title: result.Tema}

	// Convertendo []byte para string
	var item string
	perguntaJSON := string(quetionsJSONBytes)
	item = perguntaJSON
	contentItem.JsonContent = item

	err = s.repository.Create(&contentItem)
	if err == nil {
		return err
	}
	return nil

}

func (s *ContentServiceImpl) GetContentByID(id uint) (*Content, error) {
	return s.repository.FindByID(id)
}

func (s *ContentServiceImpl) GetAllContent() ([]Content, error) {
	return s.repository.FindAll()
}

func (s *ContentServiceImpl) processContent() ([]Content, error) {

	return s.repository.FindAll()
}
