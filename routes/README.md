# Microservicio de Rutas - Chetapp

Este microservicio gestiona rutas con origen, destino y un array de coordenadas (lat, lon) para la aplicación Chetapp. Permite cargar y procesar archivos GPX para crear rutas automáticamente.

## Características

- CRUD completo de rutas (Crear, Leer, Actualizar, Eliminar)
- Carga y procesamiento de archivos GPX
- Almacenamiento en base de datos PostgreSQL
- Autenticación mediante JWT

## Estructura de datos

Una ruta contiene:

- ID (UUID)
- Nombre
- Origen (coordenadas)
- Destino (coordenadas)
- Array de coordenadas (puntos intermedios)
- Distancia (en metros)
- Duración (en segundos)
- Metadatos (creación, actualización)

## API Endpoints

| Método | Endpoint | Descripción |
|--------|----------|-------------|
| GET | `/routes` | Obtiene todas las rutas |
| GET | `/routes/:id` | Obtiene una ruta por su ID |
| POST | `/routes` | Crea una nueva ruta |
| PUT | `/routes/:id` | Actualiza una ruta existente |
| DELETE | `/routes/:id` | Elimina una ruta |
| POST | `/routes/upload-gpx` | Carga y procesa un archivo GPX |

## Ejemplo de carga de archivo GPX

```bash
curl -X POST \
  http://localhost:8080/routes/upload-gpx \
  -H 'Authorization: Bearer YOUR_TOKEN' \
  -F 'gpx_file=@/ruta/al/archivo.gpx' \
  -F 'name=Mi Ruta GPX'
```

## Ejemplo de JSON de una ruta

```json
{
  "id": "f47ac10b-58cc-4372-a567-0e02b2c3d479",
  "name": "Ruta de ejemplo",
  "origin": "4.570868,-74.297333",
  "destination": "4.641654,-74.078683",
  "coordinates": [
    {"lat": 4.570868, "lon": -74.297333},
    {"lat": 4.590209, "lon": -74.240103},
    {"lat": 4.629113, "lon": -74.091798},
    {"lat": 4.641654, "lon": -74.078683}
  ],
  "distance": 24500.34,
  "duration": 1764,
  "created_at": "2023-05-10T15:00:00Z",
  "updated_at": "2023-05-10T15:00:00Z",
  "created_by": "a47ac10b-58cc-4372-a567-0e02b2c3d479",
  "updated_by": "a47ac10b-58cc-4372-a567-0e02b2c3d479"
}
```

## Configuración

El microservicio utiliza las siguientes variables de entorno:

```
SERVER_PORT=8082
POSTGRES_HOST=localhost
POSTGRES_PORT=5435
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_DB=chetapp-routes-db
POSTGRES_SSLMODE=disable
JWT_SECRET=your-secret-key
JWT_EXPIRATION_HOURS=24
```

## Ejecución

### Desarrollo local

```bash
cd routes
go mod download
go run cmd/main.go
```

### Docker Compose

```bash
docker-compose up -d routes
```

## Tecnologías utilizadas

- Go 1.21+
- Gin Web Framework
- GORM (ORM para Go)
- PostgreSQL
- JWT para autenticación
- GPXGo para procesamiento de archivos GPX
