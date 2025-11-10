# API_Busquedas - Microservicio de Búsqueda con Solr

Este microservicio implementa búsqueda paginada, filtrada y ordenada sobre actividades utilizando Apache Solr como motor de búsqueda, con doble capa de caché (local + Memcached) y sincronización en tiempo real mediante RabbitMQ.

## 🎯 Responsabilidades

- **Búsqueda paginada** con filtros y ordenamiento sobre actividades
- **Indexación en Solr** de la entidad principal (actividades)
- **Sincronización automática** mediante eventos de RabbitMQ
- **Doble capa de caché**: Local (en memoria) + Memcached (distribuida)
- **Garantía de consistencia** consultando API_Actividades por ID

## 🏗️ Arquitectura

```
┌─────────────────┐
│  API_Actividades│
│   (Puerto 8081) │
└────────┬────────┘
         │ HTTP GET /actividad/:id
         │ (validación de consistencia)
         │
    ┌────▼──────────────────────────────┐
    │      API_Busquedas (8083)          │
    │  ┌──────────────────────────────┐  │
    │  │  Controller (search)         │  │
    │  └───────┬─────────────┬────────┘  │
    │          │             │            │
    │  ┌───────▼──┐   ┌──────▼─────┐    │
    │  │  Cache   │   │   Solr     │    │
    │  │ (L1+L2)  │   │  Search    │    │
    │  └──────────┘   └────────────┘    │
    │                                     │
    │  ┌──────────────────────────────┐  │
    │  │  RabbitMQ Consumer           │  │
    │  │  (background goroutine)      │  │
    │  └───────┬──────────────────────┘  │
    └──────────┼──────────────────────────┘
               │
    ┌──────────▼──────────┐
    │  RabbitMQ (5672)    │
    │  Queue: actividades_│
    │         events      │
    └─────────────────────┘
               ▲
               │ Events: create/update/delete
               │
    ┌──────────┴────────┐
    │  API_Actividades  │
    │  (Producer)       │
    └───────────────────┘

┌──────────────────────┐
│  Solr (8983)         │
│  Core: actividades   │
└──────────────────────┘

┌──────────────────────┐
│  Memcached (11211)   │
│  Cache distribuida   │
└──────────────────────┘
```

## 📁 Estructura del Proyecto

```
API_Busquedas/
├── app/                      # Configuración del servidor Gin
│   ├── app.go               # Inicialización y CORS
│   └── url_mappings.go      # Rutas del API
├── controllers/search/      # Controladores de búsqueda
│   └── search_controller.go
├── cache/                   # Doble capa de caché
│   └── cache.go             # Local + Memcached
├── search/                  # Integración con Solr
│   └── solr.go
├── queue/                   # Consumidor de RabbitMQ
│   └── consumer.go
├── clients/                 # Clientes HTTP
│   └── actividades_client.go  # Consulta a API_Actividades
├── model/                   # Modelos de datos
│   └── actividad_search.go
├── solr/conf/               # Configuración de Solr
│   ├── managed-schema.xml   # Schema de actividades
│   └── solrconfig.xml       # Configuración de Solr
├── docker-compose.yml       # Orquestación de servicios
├── Dockerfile               # Imagen del microservicio
├── go.mod                   # Dependencias
└── README.md                # Esta documentación
```

## 🚀 Endpoints Disponibles

### 1. Búsqueda Paginada
```
GET /search/actividades
```

**Query Parameters:**
- `q` - Query de búsqueda (default: "*:*" - todos)
- `page` - Número de página (default: 1, min: 1)
- `size` - Tamaño de página (default: 10, min: 1, max: 100)
- `sort` - Campo para ordenar (ej: "nombre", "profesor")
- `order` - Orden: "asc" o "desc" (default: "asc")

**Ejemplo:**
```bash
curl "http://localhost:8083/search/actividades?q=yoga&page=1&size=10&sort=nombre&order=asc"
```

**Response:**
```json
{
  "actividades": [
    {
      "id": "507f1f77bcf86cd799439011",
      "nombre": "Yoga",
      "descripcion": "Clase de yoga para todos los niveles",
      "profesor": "María González",
      "tags": ["relax", "flexibility"],
      "horarios": "[{\"dia\":\"Lunes\",\"horaInicio\":\"10:00\",\"horaFin\":\"11:00\",\"cupo\":20}]"
    }
  ],
  "total": 1,
  "page": 1,
  "page_size": 10,
  "total_pages": 1
}
```

### 2. Obtener Actividad por ID (con caché)
```
GET /actividad/:id
```

**Ejemplo:**
```bash
curl http://localhost:8083/actividad/507f1f77bcf86cd799439011
```

### 3. Health Check
```
GET /health
```

**Response:**
```json
{
  "status": "healthy",
  "solr": "connected",
  "cache": "active"
}
```

### 4. Estadísticas de Caché
```
GET /stats
```

**Response:**
```json
{
  "local_items": 42,
  "memcached": {...},
  "memcached_enabled": true
}
```

## 🔄 Sincronización con RabbitMQ

El servicio consume eventos de la cola `actividades_events`:

### Formato de Eventos

```json
{
  "operation": "create|update|delete",
  "actividad_id": "507f1f77bcf86cd799439011",
  "timestamp": "2025-11-10T10:30:00Z"
}
```

### Flujo de Sincronización

1. **Event Received**: Consumidor recibe evento de RabbitMQ
2. **Validation**: Consulta a API_Actividades para obtener datos actualizados
3. **Indexing**: Actualiza el índice en Solr
4. **Cache Invalidation**: Elimina entrada de ambas capas de caché
5. **Confirmation**: Commit en Solr

## 💾 Sistema de Caché (Doble Capa)

### Capa 1: Caché Local (en memoria)
- Implementación similar a CCache
- TTL: 5 minutos
- Sincronización por goroutine
- Sin dependencias externas

### Capa 2: Memcached (distribuida)
- Compartida entre instancias
- TTL: 5 minutos
- Fallback si no está disponible

### Estrategia de Caché

```
GET request → L1 (local) → L2 (Memcached) → Solr/API → Store in L1 & L2
```

## 🐳 Ejecución con Docker

### Iniciar todos los servicios

```bash
cd backend/API_Busquedas
docker-compose up -d
```

Esto inicia:
- ✅ Solr (puerto 8983)
- ✅ RabbitMQ (5672 + Management UI en 15672)
- ✅ Memcached (11211)
- ✅ API_Busquedas (8083)

### Ver logs

```bash
# Ver todos los logs
docker-compose logs -f

# Ver logs específicos
docker-compose logs -f api_busquedas
docker-compose logs -f solr
docker-compose logs -f rabbitmq
```

### Acceder a las UIs

- **Solr Admin**: http://localhost:8983/solr/
- **RabbitMQ Management**: http://localhost:15672/ (guest/guest)
- **API Health**: http://localhost:8083/health

### Detener servicios

```bash
docker-compose down
```

### Reiniciar desde cero (eliminar datos)

```bash
docker-compose down -v
docker-compose up -d
```

## 🔧 Variables de Entorno

| Variable | Descripción | Default |
|----------|-------------|---------|
| **Solr** | | |
| `SOLR_HOST` | Host de Solr | solr |
| `SOLR_PORT` | Puerto de Solr | 8983 |
| **RabbitMQ** | | |
| `RABBITMQ_URL` | URL de conexión | amqp://guest:guest@rabbitmq:5672/ |
| `QUEUE_NAME` | Nombre de la cola | actividades_events |
| **Memcached** | | |
| `MEMCACHED_HOST` | Host de Memcached | memcached |
| `MEMCACHED_PORT` | Puerto de Memcached | 11211 |
| **API_Actividades** | | |
| `API_ACTIVIDADES_HOST` | Host de API_Actividades | host.docker.internal |
| `API_ACTIVIDADES_PORT` | Puerto de API_Actividades | 8081 |

## 🧪 Pruebas

### 1. Probar búsqueda básica

```bash
curl "http://localhost:8083/search/actividades?q=*:*&page=1&size=10"
```

### 2. Buscar por nombre

```bash
curl "http://localhost:8083/search/actividades?q=yoga"
```

### 3. Búsqueda con ordenamiento

```bash
curl "http://localhost:8083/search/actividades?q=*:*&sort=nombre&order=desc"
```

### 4. Verificar caché

```bash
# Primera llamada (sin caché)
time curl "http://localhost:8083/search/actividades?q=yoga"

# Segunda llamada (con caché, debe ser más rápida)
time curl "http://localhost:8083/search/actividades?q=yoga"
```

### 5. Ver estadísticas

```bash
curl http://localhost:8083/stats
```

### 6. Publicar evento de prueba en RabbitMQ

```bash
# Instalar amqp-tools
# Ubuntu: sudo apt-get install amqp-tools
# Mac: brew install rabbitmq-c

# Publicar mensaje
docker exec -it rabbitmq_busquedas rabbitmqadmin publish \
  routing_key=actividades_events \
  payload='{"operation":"create","actividad_id":"507f1f77bcf86cd799439011","timestamp":"2025-11-10T10:30:00Z"}'
```

## 📊 Monitoreo

### Solr Admin UI
```
http://localhost:8983/solr/#/actividades
```

Permite:
- Ver documentos indexados
- Ejecutar queries manuales
- Monitorear estadísticas del core

### RabbitMQ Management
```
http://localhost:15672/
Usuario: guest
Password: guest
```

Permite:
- Ver colas y mensajes
- Monitorear consumers
- Ver throughput

## 🔒 Consideraciones de Producción

### Escalabilidad
- ✅ Stateless (puede escalar horizontalmente)
- ✅ Memcached compartida entre instancias
- ✅ Solr puede configurarse en modo Cloud (SolrCloud)

### Alta Disponibilidad
- Configurar Solr en cluster
- RabbitMQ en modo cluster
- Load balancer para múltiples instancias

### Seguridad
- Autenticación en RabbitMQ
- Solr con autenticación básica
- Rate limiting en endpoints
- Validación de inputs

### Observabilidad
- Métricas de Solr
- Logs estructurados con Logrus
- Métricas de caché (hit rate)
- Tracing distribuido (Jaeger/Zipkin)

## 📝 Notas de Implementación

### Garantía de Consistencia
Para cada evento de create/update, el servicio:
1. Consulta API_Actividades por ID (fuente de verdad)
2. Indexa los datos obtenidos en Solr
3. Invalida caché para forzar refresh

Esto garantiza que Solr siempre tiene datos consistentes con la base de datos principal.

### Gestión de Errores
- Reintentos automáticos en conexión a RabbitMQ
- Fallback a caché local si Memcached no disponible
- Logging detallado de errores

### Rendimiento
- Búsquedas cacheadas: < 5ms
- Búsquedas en Solr: 10-50ms
- Indexación asíncrona (no bloquea)

## 🤝 Integración con Otros Microservicios

### API_Actividades
- **Produce eventos** a RabbitMQ en operaciones CUD
- **Responde a consultas** por ID para validación

### Frontend
- **Consume** endpoint de búsqueda paginada
- **Muestra resultados** con filtros y ordenamiento

## 🚧 TODOs

- [ ] Implementar búsqueda fuzzy (tolerancia a errores)
- [ ] Agregar faceted search (agrupación por campos)
- [ ] Implementar autocomplete
- [ ] Métricas con Prometheus
- [ ] Tests unitarios y de integración
- [ ] CI/CD pipeline
- [ ] Documentación OpenAPI/Swagger
- [ ] Rate limiting por usuario
- [ ] Búsqueda por geolocalización (si aplica)

## 📞 Contacto y Soporte

Para reportar problemas o sugerencias, crear un issue en el repositorio.
