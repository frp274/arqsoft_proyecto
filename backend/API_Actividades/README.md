# API Actividades - Microservicio de Gestión de Actividades

## Descripción

Microservicio encargado de la gestión CRUD de actividades deportivas. Utiliza MongoDB como base de datos principal y publica eventos en RabbitMQ para sincronización con otros microservicios (especialmente API_Busquedas).

## Arquitectura

```
┌─────────────────────────────────────────────────────────────┐
│                     API_Actividades                         │
├─────────────────────────────────────────────────────────────┤
│                                                             │
│  ┌──────────────┐      ┌──────────────┐                   │
│  │ Controllers  │◄─────┤   Services   │                   │
│  └──────────────┘      └──────┬───────┘                   │
│                               │                            │
│                     ┌─────────┴─────────┐                 │
│                     │                   │                 │
│              ┌──────▼──────┐    ┌──────▼──────┐          │
│              │ Repositories│    │ RabbitMQ    │          │
│              │  (MongoDB)  │    │  Producer   │          │
│              └──────┬──────┘    └──────┬──────┘          │
│                     │                   │                 │
└─────────────────────┼───────────────────┼─────────────────┘
                      │                   │
                      │                   │
              ┌───────▼────────┐  ┌───────▼────────┐
              │    MongoDB     │  │   RabbitMQ     │
              │   actividades  │  │  Exchange:     │
              │      _db       │  │  actividades_  │
              │                │  │   exchange     │
              │  Collections:  │  │                │
              │  - actividades │  │  Queue:        │
              │  - horarios    │  │  actividades_  │
              │                │  │    events      │
              └────────────────┘  └────────────────┘
```

## Características

### Base de Datos (MongoDB)
- **Gestión de Actividades**: CRUD completo de actividades deportivas
- **Caché Local**: Implementación de caché en memoria usando CCache para mejorar performance
- **Colecciones**:
  - `actividades`: Información de las actividades (nombre, descripción, profesor, horarios)
  - Embedded: `horarios` dentro de cada actividad

### Sistema de Eventos (RabbitMQ)
- **Productor de Eventos**: Publica eventos cuando ocurren cambios en las actividades
- **Tipos de Eventos**:
  - `create`: Cuando se crea una nueva actividad
  - `update`: Cuando se actualiza una actividad existente
  - `delete`: Cuando se elimina una actividad
- **Exchange**: `actividades_exchange` (tipo fanout)
- **Cola**: `actividades_events`

### Endpoints REST

#### GET /actividad/:id
Obtiene una actividad por su ID.

**Request:**
```bash
curl http://localhost:8081/actividad/507f1f77bcf86cd799439011
```

**Response:**
```json
{
  "id": "507f1f77bcf86cd799439011",
  "nombre": "Yoga",
  "descripcion": "Clase de yoga para principiantes",
  "profesor": "Juan Pérez",
  "horario": [
    {
      "dia": "Lunes",
      "hora_inicio": "09:00",
      "hora_fin": "10:00",
      "cupo": 20
    }
  ]
}
```

#### POST /actividad
Crea una nueva actividad.

**Request:**
```bash
curl -X POST http://localhost:8081/actividad \
  -H "Content-Type: application/json" \
  -d '{
    "nombre": "Pilates",
    "descripcion": "Clase de pilates avanzado",
    "profesor": "María García",
    "horario": [
      {
        "dia": "Martes",
        "hora_inicio": "18:00",
        "hora_fin": "19:00",
        "cupo": 15
      }
    ]
  }'
```

**Response:**
```json
{
  "id": "507f1f77bcf86cd799439012",
  "nombre": "Pilates",
  "descripcion": "Clase de pilates avanzado",
  "profesor": "María García",
  "horario": [
    {
      "dia": "Martes",
      "hora_inicio": "18:00",
      "hora_fin": "19:00",
      "cupo": 15
    }
  ]
}
```

**Evento Publicado:**
```json
{
  "type": "create",
  "actividad_id": "507f1f77bcf86cd799439012",
  "timestamp": "2025-11-10T10:30:00Z"
}
```

#### PUT /actividad
Actualiza una actividad existente.

**Request:**
```bash
curl -X PUT http://localhost:8081/actividad \
  -H "Content-Type: application/json" \
  -d '{
    "id": "507f1f77bcf86cd799439012",
    "nombre": "Pilates Avanzado",
    "descripcion": "Clase de pilates nivel avanzado",
    "profesor": "María García",
    "horario": [
      {
        "dia": "Martes",
        "hora_inicio": "18:00",
        "hora_fin": "19:30",
        "cupo": 12
      }
    ]
  }'
```

**Response:**
```json
{
  "id": "507f1f77bcf86cd799439012",
  "nombre": "Pilates Avanzado",
  "descripcion": "Clase de pilates nivel avanzado",
  "profesor": "María García",
  "horario": [
    {
      "dia": "Martes",
      "hora_inicio": "18:00",
      "hora_fin": "19:30",
      "cupo": 12
    }
  ]
}
```

**Evento Publicado:**
```json
{
  "type": "update",
  "actividad_id": "507f1f77bcf86cd799439012",
  "timestamp": "2025-11-10T10:35:00Z"
}
```

#### DELETE /actividad/:id
Elimina una actividad.

**Request:**
```bash
curl -X DELETE http://localhost:8081/actividad/507f1f77bcf86cd799439012
```

**Response:**
```
204 No Content
```

**Evento Publicado:**
```json
{
  "type": "delete",
  "actividad_id": "507f1f77bcf86cd799439012",
  "timestamp": "2025-11-10T10:40:00Z"
}
```

## Configuración

### Variables de Entorno

```bash
# MongoDB
MONGO_URI=mongodb://mongouser:mongopass@mongodb:27017
MONGO_DB_NAME=actividades_db

# RabbitMQ
RABBITMQ_URL=amqp://guest:guest@rabbitmq:5672/
RABBITMQ_QUEUE=actividades_events
RABBITMQ_EXCHANGE=actividades_exchange
```

### Docker Compose

El archivo `docker-compose.yml` incluye:
- **MongoDB**: Base de datos principal (puerto 27017)
- **RabbitMQ**: Message broker con interfaz de administración (puertos 5672 y 15672)
- **API_Actividades**: El microservicio (puerto 8081)

## Instalación y Ejecución

### Con Docker Compose (Recomendado)

```bash
# Iniciar todos los servicios
docker-compose up -d

# Ver logs
docker-compose logs -f api_actividades

# Detener servicios
docker-compose down

# Detener y eliminar volúmenes (datos)
docker-compose down -v
```

### Desarrollo Local

```bash
# Instalar dependencias
go mod download

# Compilar
go build -o api_actividades

# Ejecutar (requiere MongoDB y RabbitMQ corriendo)
export MONGO_URI="mongodb://mongouser:mongopass@localhost:27017"
export MONGO_DB_NAME="actividades_db"
export RABBITMQ_URL="amqp://guest:guest@localhost:5672/"
export RABBITMQ_QUEUE="actividades_events"
export RABBITMQ_EXCHANGE="actividades_exchange"
./api_actividades
```

## Dependencias

```go
require (
    github.com/gin-gonic/gin v1.11.0
    github.com/karlseguin/ccache v2.0.3+incompatible
    github.com/rabbitmq/amqp091-go v1.10.0
    github.com/sirupsen/logrus v1.9.3
    go.mongodb.org/mongo-driver v1.17.4
)
```

## Integración con Otros Microservicios

### API_Busquedas
Este microservicio consume los eventos publicados por API_Actividades para mantener sincronizado el índice de Solr:

1. **Evento CREATE**: API_Busquedas indexa la nueva actividad en Solr
2. **Evento UPDATE**: API_Busquedas actualiza el documento en Solr
3. **Evento DELETE**: API_Busquedas elimina el documento de Solr

### Flujo de Sincronización

```
API_Actividades                  RabbitMQ              API_Busquedas
     │                              │                       │
     │ 1. POST /actividad           │                       │
     ├─────────────────────────────►│                       │
     │                              │                       │
     │ 2. Save to MongoDB           │                       │
     ├──────────────┐               │                       │
     │              │               │                       │
     │◄─────────────┘               │                       │
     │                              │                       │
     │ 3. Publish Event             │                       │
     ├─────────────────────────────►│                       │
     │                              │                       │
     │                              │ 4. Consume Event      │
     │                              ├──────────────────────►│
     │                              │                       │
     │                              │ 5. Fetch from API     │
     │◄────────────────────────────────────────────────────┤
     │                              │                       │
     │ 6. Return Activity Data      │                       │
     ├─────────────────────────────────────────────────────►│
     │                              │                       │
     │                              │ 7. Index in Solr      │
     │                              │       ┌───────────────┤
     │                              │       │               │
     │                              │       └──────────────►│
```

## Monitoring

### RabbitMQ Management UI
Acceder a http://localhost:15672
- **Usuario**: guest
- **Password**: guest

Desde aquí puedes:
- Ver el estado de las colas
- Monitorear mensajes publicados y consumidos
- Ver conexiones activas
- Depurar problemas de mensajería

### Logs
Los logs se escriben en formato JSON usando Logrus:

```bash
# Ver logs en tiempo real
docker-compose logs -f api_actividades

# Logs de RabbitMQ
docker-compose logs -f rabbitmq

# Logs de MongoDB
docker-compose logs -f mongodb
```

## Testing

### Probar la Publicación de Eventos

1. **Crear una actividad y verificar el evento**:
```bash
# Terminal 1: Monitorear la cola en RabbitMQ UI
open http://localhost:15672

# Terminal 2: Crear actividad
curl -X POST http://localhost:8081/actividad \
  -H "Content-Type: application/json" \
  -d '{
    "nombre": "Test Activity",
    "descripcion": "Testing events",
    "profesor": "Test Professor",
    "horario": [{"dia": "Lunes", "hora_inicio": "10:00", "hora_fin": "11:00", "cupo": 10}]
  }'
```

2. **Verificar en RabbitMQ UI**:
   - Ve a "Queues" → `actividades_events`
   - Deberías ver 1 mensaje en la cola
   - Click en "Get messages" para ver el contenido

## Troubleshooting

### Error: "Failed to connect to MongoDB"
```bash
# Verificar que MongoDB está corriendo
docker-compose ps mongodb

# Ver logs de MongoDB
docker-compose logs mongodb

# Reiniciar MongoDB
docker-compose restart mongodb
```

### Error: "Failed to initialize RabbitMQ producer"
```bash
# Verificar que RabbitMQ está corriendo
docker-compose ps rabbitmq

# Ver logs de RabbitMQ
docker-compose logs rabbitmq

# Reiniciar RabbitMQ
docker-compose restart rabbitmq
```

### Los eventos no se publican
```bash
# Verificar logs del servicio
docker-compose logs api_actividades | grep -i rabbit

# Verificar conexiones en RabbitMQ UI
# http://localhost:15672/#/connections
```

## Próximos Pasos

- [ ] Agregar tests unitarios para el productor de RabbitMQ
- [ ] Implementar health check endpoint
- [ ] Agregar métricas de Prometheus
- [ ] Implementar circuit breaker para RabbitMQ
- [ ] Agregar retry logic para eventos fallidos
- [ ] Implementar dead letter queue para mensajes no procesables

## Autor

Proyecto de Arquitectura de Software - 2025
