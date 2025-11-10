# 🏋️ GOOD GYM - Deployment con Docker Compose

## 📋 Descripción

Sistema de gestión de gimnasio implementado con arquitectura de microservicios:

- **API_Usuarios** (Puerto 8082): Autenticación y gestión de usuarios con JWT
- **API_Actividades** (Puerto 8081): CRUD de actividades deportivas con RabbitMQ
- **API_Busquedas** (Puerto 8083): Búsqueda con Solr y caché (Memcached)
- **Frontend** (Puerto 3000): Interfaz React

## 🚀 Prerequisitos

- Docker Desktop instalado y en ejecución
- Puertos libres: 3000, 8081, 8082, 8083, 8983, 5672, 15672, 11211, 3308, 27017

## 📦 Instalación y Ejecución

### 1. Clonar el repositorio

```bash
git clone <repo-url>
cd arqsoft_proyecto
```

### 2. Configurar variables de entorno del frontend

```bash
cd frontend
cp .env.example .env
# Verificar que .env tenga:
# REACT_APP_API_USUARIOS_URL=http://localhost:8082
# REACT_APP_API_ACTIVIDADES_URL=http://localhost:8081
# REACT_APP_API_BUSQUEDAS_URL=http://localhost:8083
```

### 3. Levantar todos los servicios

Desde el directorio raíz del proyecto:

```bash
docker-compose up --build
```

Esto creará y levantará:
- ✅ MySQL para usuarios (puerto 3308)
- ✅ MongoDB para actividades (puerto 27017)
- ✅ RabbitMQ (puerto 5672, dashboard 15672)
- ✅ Solr (puerto 8983)
- ✅ Memcached (puerto 11211)
- ✅ API_Usuarios (puerto 8082)
- ✅ API_Actividades (puerto 8081)
- ✅ API_Busquedas (puerto 8083)
- ✅ Frontend React (puerto 3000)

### 4. Esperar a que todos los servicios estén listos

Los servicios tienen healthchecks configurados. Espera unos 30-60 segundos hasta que todos muestren `healthy`.

Puedes verificar el estado con:

```bash
docker-compose ps
```

### 5. Acceder a la aplicación

Abre tu navegador en: **http://localhost:3000**

## 👤 Flujo del Usuario

### Login → Búsqueda → Detalle → Acción → Confirmación

1. **Login**: Ingresa con un usuario existente o crea uno nuevo
2. **Home**: Busca actividades por nombre (ej: "Pilates", "MMA", "Zumba")
3. **Detalle**: Haz clic en una actividad para ver horarios y cupos
4. **Acción**: Haz clic en "Inscribirme" para verificar disponibilidad
5. **Confirmación**: Verás un mensaje de éxito o de error según la disponibilidad

## 🧪 Endpoints Disponibles

### API_Usuarios (8082)
- `POST /login` - Autenticación
- `GET /usuario/:id` - Obtener usuario
- `POST /usuario` - Crear usuario

### API_Actividades (8081)
- `GET /actividad/:id` - Obtener actividad
- `POST /actividad` - Crear actividad
- `PUT /actividad/:id` - Actualizar actividad
- `DELETE /actividad/:id` - Eliminar actividad
- `POST /actividad/:id/calcular-disponibilidad` - Verificar cupos

### API_Busquedas (8083)
- `GET /search/actividades?nombre=X` - Buscar actividades
- `GET /actividad/:id` - Obtener detalle (desde caché/Solr)
- `GET /health` - Estado del servicio
- `GET /stats` - Estadísticas de caché

## 🛠️ Herramientas de Administración

- **RabbitMQ Dashboard**: http://localhost:15672 (usuario: `guest`, contraseña: `guest`)
- **Solr Admin**: http://localhost:8983/solr

## 🐛 Troubleshooting

### Los servicios no se levantan

```bash
# Detener todos los contenedores
docker-compose down

# Limpiar volúmenes (⚠️ borra datos)
docker-compose down -v

# Volver a construir e iniciar
docker-compose up --build
```

### Error de conexión entre microservicios

Verifica que todos los servicios estén `healthy`:

```bash
docker-compose ps
```

Revisa logs de un servicio específico:

```bash
docker-compose logs -f api_usuarios
docker-compose logs -f api_actividades
docker-compose logs -f api_busquedas
```

### Frontend no se conecta al backend

1. Verifica que el archivo `frontend/.env` exista y tenga las URLs correctas
2. Reconstruye el frontend: `docker-compose up --build frontend`

### RabbitMQ no conecta

Espera 30 segundos más. RabbitMQ tarda en iniciarse completamente.

```bash
docker-compose logs -f rabbitmq
```

## 📊 Arquitectura Técnica

### Event-Driven con RabbitMQ

```
API_Actividades ──[Publish]──> RabbitMQ ──[Consume]──> API_Busquedas ──> Solr
```

Cuando se crea/actualiza/elimina una actividad:
1. API_Actividades publica evento en RabbitMQ
2. API_Busquedas consume el evento
3. Actualiza índice de Solr automáticamente

### Caché Multi-Nivel

```
Request ──> L1 Cache (Local) ──> L2 Cache (Memcached) ──> Solr/MongoDB
```

- **L1**: Caché en memoria local (1 minuto TTL)
- **L2**: Memcached compartido (5 minutos TTL)
- **Fuente**: Solr para búsquedas, MongoDB para CRUD

### Autenticación JWT

```
Frontend ──[JWT Token]──> API_Usuarios ──[Validate]──> API_Actividades
```

- Login genera JWT con claims (id, username, es_admin)
- Token se valida en cada petición protegida
- Middleware verifica ownership y permisos admin

## 📝 Notas de Desarrollo

### Concurrencia

`API_Actividades` implementa concurrencia con:
- **GoRoutines**: Para procesar múltiples horarios en paralelo
- **Channels**: Para comunicar resultados entre goroutines
- **WaitGroups**: Para sincronizar finalización de tareas

### Validación entre Microservicios

`API_Actividades` valida el `owner_id` contra `API_Usuarios` antes de crear/modificar actividades.

### Testing

Cada microservicio tiene tests unitarios:

```bash
# Ejecutar tests de un microservicio
cd backend/API_Usuarios
go test ./...
```

## 🔄 Detener los Servicios

```bash
# Detener sin borrar datos
docker-compose down

# Detener y borrar volúmenes (⚠️ borra BD)
docker-compose down -v
```

## 📧 Soporte

Para reportar issues o contribuir, crear un issue en el repositorio del proyecto.

---

**¡Listo para entrenar! 💪**
