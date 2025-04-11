# Gestión de Mascotas

Este proyecto es un microservicio para la gestión de mascotas, desarrollado en Go. Proporciona funcionalidades para agregar, gestionar y consultar información sobre mascotas.

## Estructura del Proyecto

El proyecto está organizado de la siguiente manera:

```
gestion-de-mascotas
├── cmd
│   └── main.go               # Punto de entrada de la aplicación
├── internal
│   ├── pets
│   │   ├── handler.go        # Manejador de solicitudes HTTP para mascotas
│   │   ├── service.go        # Lógica de negocio para la gestión de mascotas
│   │   └── repository.go     # Interacción con la base de datos
│   └── db
│       └── connection.go     # Conexión a la base de datos
├── migrations
│   └── 001_create_tables.sql # Instrucciones SQL para crear tablas
├── go.mod                    # Configuración del módulo Go
├── go.sum                    # Sumas de verificación de dependencias
└── README.md                 # Documentación del proyecto
```

## Requisitos

- Go 1.16 o superior
- Una base de datos compatible (por ejemplo, PostgreSQL, MySQL)

## Instalación

1. Clona el repositorio:
   ```
   git clone <URL_DEL_REPOSITORIO>
   cd gestion-de-mascotas
   ```

2. Instala las dependencias:
   ```
   go mod tidy
   ```

3. Configura la base de datos y ejecuta las migraciones:
   - Asegúrate de que tu base de datos esté en funcionamiento.
   - Ejecuta el archivo de migración `001_create_tables.sql` para crear las tablas necesarias.

## Ejecución

Para ejecutar el microservicio, utiliza el siguiente comando:

```
go run cmd/main.go
```

El servidor se iniciará y estará disponible para recibir solicitudes relacionadas con la gestión de mascotas.

## Contribuciones

Las contribuciones son bienvenidas. Si deseas contribuir, por favor abre un issue o envía un pull request.

## Licencia

Este proyecto está bajo la Licencia MIT. Consulta el archivo LICENSE para más detalles.