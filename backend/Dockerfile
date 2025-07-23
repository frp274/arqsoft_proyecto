FROM golang:1.24

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o server

EXPOSE 8080

# CMD ["./server"]
CMD ["sh", "-c", "echo INICIO && ./server || echo 🚨 FALLÓ"]
