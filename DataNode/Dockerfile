# Utiliza una imagen base de Go
FROM golang:latest

# Establece el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copia los archivos Go necesarios al contenedor

COPY ../proto/ /app/proto
COPY ../DataNode /app/
COPY ../go.mod /app/
COPY ../go.sum /app/
COPY ../README.md /app/

# Compila tu servidor Go
RUN go build datanode.go

# Expone el puerto en el que escucha tu servidor
EXPOSE 50055

# Ejecuta tu servidor cuando se inicie el contenedor
CMD ["./datanode"]
