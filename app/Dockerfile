# Estágio 1: Build
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copia os arquivos de dentro da pasta local /app para o WORKDIR atual
COPY go.mod go.sum ./
RUN go mod download

# Copia todo o conteúdo da pasta local /app para o container
COPY . . 

# Compila o binário estático
RUN go build -o main .

# Estágio 2: Execução
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .

# ADICIONE ESTA LINHA:
COPY .env . 

EXPOSE 3000
CMD ["./main"]
