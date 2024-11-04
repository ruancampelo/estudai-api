# Etapa de construção
FROM golang:1.23-alpine AS builder

# Instala ferramentas auxiliares e define o diretório de trabalho
RUN apk add --no-cache git
WORKDIR /app

# Copia o diretório estudai-api para o container
COPY . ./

# Instala as dependências Go
RUN go mod download

# Compila o programa Go a partir do diretório cmd
RUN go build -o main ./cmd

COPY ca-bundle.crt /etc/ssl/certs/ca-bundle.crt
COPY ca-bundle.trust.crt /etc/ssl/certs/ca-bundle.trust.crt

# Etapa final para execução da aplicação
FROM scratch
COPY --from=builder /app/main /main
EXPOSE 5112
ENTRYPOINT ["/main"]
