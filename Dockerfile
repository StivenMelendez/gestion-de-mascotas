# Usar una imagen base de Go
FROM golang:1.20

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar los archivos del proyecto al contenedor
COPY . .

# Descargar las dependencias del proyecto
RUN go mod tidy

# Compilar la aplicación
RUN go build -o main .

# Exponer el puerto en el que se ejecutará la aplicación
EXPOSE 8080

# Comando para ejecutar la aplicación
CMD ["./main"]