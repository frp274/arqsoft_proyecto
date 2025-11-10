FROM golang:1.24.1-alpine

WORKDIR /app

# Copiar go.mod y go.sum
COPY go.mod go.sum* ./

# Descargar dependencias
RUN go mod download

# Copiar todo el código fuente
COPY . .

# Compilar la aplicación
RUN go build -o main .

# Exponer el puerto
EXPOSE 8082

# Ejecutar la aplicación
CMD ["./main"]
