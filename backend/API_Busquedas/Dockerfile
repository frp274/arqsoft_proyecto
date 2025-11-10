FROM golang:1.24.1-alpine AS builder

WORKDIR /app

# Copiar archivos de dependencias
COPY go.mod go.sum* ./
RUN go mod download

# Copiar código fuente
COPY . .

# Compilar la aplicación
RUN go build -o main .

# Imagen final
FROM alpine:latest

WORKDIR /app

# Copiar el binario compilado
COPY --from=builder /app/main .

# Exponer puerto
EXPOSE 8083

# Ejecutar la aplicación
CMD ["./main"]
