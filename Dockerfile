# Utiliza una imagen base de Go
FROM golang:latest

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos Go necesarios al contenedor
COPY proto/ /app/proto
COPY server/server.go .
COPY go.mod .
COPY go.sum .
COPY README.md .

# Compila tu servidor Go
RUN go build server.go

# Expone el puerto en el que escucha tu servidor
EXPOSE 50051

# Ejecuta tu servidor cuando se inicie el contenedor
CMD ["./server"]
