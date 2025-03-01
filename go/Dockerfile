# Usa una imagen base de Go
FROM golang:1.22 AS builder

# Setea el directorio de trabajo
WORKDIR /app

# Copia el go.mod y go.sum y descarga las dependencias
COPY go/go.mod go/go.sum ./
RUN go mod download

# Copia el resto del código fuente
COPY go/ ./

# Compila la aplicación
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./main.go

# Usa una imagen base ligera para el contenedor final
FROM alpine:latest

# Setea el directorio de trabajo
WORKDIR /app

# Copia el binario de la etapa de construcción
COPY --from=builder /app/main .

# Copia los archivos HTML
COPY web /app/web

# Expone el puerto
EXPOSE 8080

RUN chmod +x /app/main
RUN apk add --no-cache libc6-compat

# Comando para ejecutar la aplicación
CMD ["./main"]
