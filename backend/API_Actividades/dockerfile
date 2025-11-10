# ----------------------------------------------------
# STAGE 1: BUILDER (Compilación)
# ----------------------------------------------------
# Usa una imagen base con Go (versión 1.24.1-alpine para ser ligera)
FROM golang:1.24.1-alpine AS builder

# Establece el directorio de trabajo
WORKDIR /app

# Copia los archivos de módulos para descargar dependencias primero y aprovechar el cache
# Asegúrate de que go.mod y go.sum existan en la raíz de tu proyecto
COPY go.mod go.sum ./
RUN go mod download

# Copia el resto del código fuente
COPY . .

# Compilación Cruzada:
# - CGO_ENABLED=0: Crea un binario estático, crucial para Alpine.
# - GOOS=linux: Especifica Linux como sistema operativo de destino.
# - -o api_actividades: Define el nombre del ejecutable.
# - -ldflags '-w -s': Opciones para reducir el tamaño del binario.
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix netgo -ldflags '-w -s' -o api_actividades ./main.go

# ----------------------------------------------------
# STAGE 2: FINAL (Ejecución - Imagen ligera)
# ----------------------------------------------------
# Usa una imagen base mínima y segura (Alpine)
FROM alpine:latest

# Instala certificados SSL (necesario para HTTPS y algunas conexiones a DB)
RUN apk --no-cache add ca-certificates

# Establece la zona horaria (opcional)
ENV TZ=America/Argentina/Buenos_Aires

# Establece el directorio de trabajo
WORKDIR /root/

# Copia el ejecutable COMPILADO desde la etapa 'builder'
COPY --from=builder /app/api_actividades .

# Otorga permisos de ejecución al binario (soluciona 'permission denied')
RUN chmod +x api_actividades

# Expone el puerto que usa tu API de Gin
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./api_actividades"]