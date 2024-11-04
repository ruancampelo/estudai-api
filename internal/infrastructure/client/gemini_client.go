package client

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/generative-ai-go/genai"
	_ "github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
	_ "google.golang.org/api/option"
	"log"
	"mime/multipart"
	"os"
)

type Opcao struct {
	Texto   string `json:"texto"`
	Correta bool   `json:"correta"`
}

// Estrutura para o conteúdo (pergunta)
type Conteudo struct {
	ID        int     `json:"id"`
	Enunciado string  `json:"enunciado"`
	Opcoes    []Opcao `json:"opcoes"`
}

// Estrutura para o tema
type Tema struct {
	Tema     string     `json:"tema"`
	Conteudo []Conteudo `json:"conteudo"`
}

func GetQuestionsByFile(file *multipart.FileHeader) Tema {
	ctx, client, err, img := UploadFile(file)

	// Choose a Gemini model.
	resp := GenerateContentResponse(client, img, err, ctx)

	result := FormatResponse(resp)

	return result
}

func FormatResponse(resp *genai.GenerateContentResponse) Tema {
	var tema Tema

	for _, part := range resp.Candidates[0].Content.Parts {
		if txt, ok := part.(genai.Text); ok {
			var result []string
			if err := json.Unmarshal([]byte(txt), &result); err != nil {
				log.Fatal(err)
			}

			// Deserializando cada string JSON individualmente

			for _, conteudo := range result {
				if err := json.Unmarshal([]byte(conteudo), &tema); err != nil {
					log.Fatal(err)
				}
			}
		}
	}
	return tema
}

func GenerateContentResponse(client *genai.Client, img *genai.File, err error, ctx context.Context) *genai.GenerateContentResponse {
	model := client.GenerativeModel("gemini-1.5-flash-002")
	model.ResponseMIMEType = "application/json"
	model.ResponseSchema = &genai.Schema{
		Type:  genai.TypeArray,
		Items: &genai.Schema{Type: genai.TypeString},
	}
	// Create a prompt using text and the URI reference for the uploaded file.
	prompt := []genai.Part{
		genai.FileData{URI: img.URI},
		genai.Text(`Dado este arquivo:
					faça um quiz sobre o conteúdo do arquivo
					o quiz sera em formato json, onde cada questao tem um id numerico, 
					e cada opcão de resposta tem  um id, e ainda a quatidade exata de quiz e 5.
					Seguido o modelo do json como exemplo:
					[
						{
							"tema": "area do canhecimento",
							"conteudo": [
								{
									"id": 1,
									"enunciado": "Exemplo de enunciado",
									"opcoes": [
										{
											"texto": "opcao 1",
											"correta": false
										},
										{
											"texto": "opcao 2",
											"correta": false
										},
										{
											"texto": "opcao 3",
											"correta": true
										},
										{
											"texto": "opcao 4",
											"correta": true
										}
									]
								}
							]
						}
					]`),
	}

	// Generate content using the prompt.
	resp, err := model.GenerateContent(ctx, prompt...)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

func UploadFile(file *multipart.FileHeader) (context.Context, *genai.Client, error, *genai.File) {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatal(err)
	}

	f, err := file.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	opts := genai.UploadFileOptions{DisplayName: file.Filename}
	img, err := client.UploadFile(ctx, "", f, &opts)
	if err != nil {
		log.Fatal(err)
	}

	// View the response.
	fmt.Printf("Uploaded file %s as: %q\n", img.DisplayName, img.URI)
	return ctx, client, err, img
}
