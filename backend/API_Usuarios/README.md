# API_Usuarios - Microservicio de Autenticación y Gestión de Usuarios

Este microservicio es responsable de la autenticación y gestión de usuarios en la arquitectura de microservicios.

## 🎯 Responsabilidades

- Autenticación de usuarios (login)
- Generación y validación de tokens JWT
- Gestión de usuarios (CRUD - en desarrollo)
- Hash de contraseñas con SHA256

## 🏗️ Arquitectura

- **Puerto**: 8082
- **Base de Datos**: MySQL (puerto 3308)
- **Nombre BD**: usuarios_db
- **ORM**: GORM

## 📁 Estructura del Proyecto

```
API_Usuarios/
├── app/                    # Configuración del servidor Gin
│   ├── app.go             # Inicialización y CORS
│   └── url_mappings.go    # Definición de rutas
├── controllers/           
│   └── usuario/           # Controladores de usuario
├── services/              # Lógica de negocio
├── clients/               # Capa de acceso a datos
│   └── usuarios/          
├── db/                    # Configuración de base de datos
│   └── init/              # Scripts SQL de inicialización
├── model/                 # Modelos de datos
├── dto/                   # Data Transfer Objects
├── utils/                 # Utilidades (JWT, hash)
├── docker-compose.yml     # Configuración de contenedores
├── Dockerfile             # Imagen de la aplicación
└── go.mod                 # Dependencias del proyecto
```

## 🚀 Endpoints Disponibles

### Autenticación

#### POST /login
Autentica un usuario y devuelve un token JWT.

**Request:**
```json
{
  "Username": "genacanas",
  "Password": "genaro123"
}
```

**Response (200 OK):**
```json
{
  "id": 1,
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "es_admin": true
}
```

**Response (401 Unauthorized):**
```json
{
  "Error": "invalid password"
}
```

### Usuarios

#### GET /usuario/:id
Obtiene información de un usuario por su ID. **Requiere autenticación JWT y ser el dueño del recurso o admin**.

**Headers:**
```
Authorization: Bearer <token_jwt>
```

**Request:**
```bash
curl http://localhost:8082/usuario/1 \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
```

**Response (200 OK):**
```json
{
  "id": 1,
  "nombre_apellido": "Genaro Cañas",
  "username": "genacanas",
  "es_admin": true
}
```

**Response (401 Unauthorized):** Sin token o token inválido
```json
{
  "error": "Authorization header required"
}
```

**Response (403 Forbidden):** Usuario no es admin ni dueño
```json
{
  "error": "You can only modify your own resources"
}
```

**Response (404 Not Found):** Usuario no existe
```json
{
  "error": "User not found"
}
```

#### POST /usuario
Crea un nuevo usuario. **Requiere autenticación JWT y permisos de admin**.

**Headers:**
```
Authorization: Bearer <token_jwt_admin>
```

**Request:**
```bash
curl -X POST http://localhost:8082/usuario \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
  -H "Content-Type: application/json" \
  -d '{
    "nombre_apellido": "Juan Pérez",
    "username": "juanperez",
    "password": "mipassword123",
    "es_admin": false
  }'
```

**Response (201 Created):**
```json
{
  "id": 4,
  "nombre_apellido": "Juan Pérez",
  "username": "juanperez",
  "es_admin": false
}
```

**Response (401 Unauthorized):** Sin token o token inválido
```json
{
  "error": "Invalid or expired token"
}
```

**Response (403 Forbidden):** Usuario no es admin
```json
{
  "error": "Admin privileges required"
}
```

**Response (400 Bad Request):** Datos inválidos
```json
{
  "error": "Key: 'CreateUsuarioRequest.UserName' Error:Field validation for 'UserName' failed on the 'required' tag"
}
```

### Endpoints Futuros (Planificados)
- `PUT /usuario/:id` - Actualizar usuario (owner o admin)
- `DELETE /usuario/:id` - Eliminar usuario (solo admin)

## 🐳 Ejecución con Docker

### Iniciar el servicio

```bash
cd backend/API_Usuarios
docker-compose up -d
```

### Ver logs

```bash
docker-compose logs -f api_usuarios
```

### Detener el servicio

```bash
docker-compose down
```

### Detener y eliminar volúmenes (reiniciar BD)

```bash
docker-compose down -v
```

## 🛠️ Desarrollo Local

### Prerequisitos

- Go 1.24.1 o superior
- MySQL 8.0
- Make (opcional)

### Instalación de dependencias

```bash
go mod download
```

### Compilar

```bash
go build -o api_usuarios
```

### Ejecutar

```bash
./api_usuarios
```

O directamente:

```bash
go run main.go
```

## 🔑 Usuarios de Prueba

| Username    | Password   | Admin | Hash SHA256                                                      |
|-------------|------------|-------|------------------------------------------------------------------|
| genacanas   | genaro123  | ✅    | 3bd517332b9d96f9fbc0de89b613dc07b3101292fd54fd7cb52da0d8846303e2 |
| facubuffaz  | facu123    | ❌    | 293bb6d0e7e4c2ee8761e60be2169d09d42156f4167fa58f3e2a0e39e78773d4 |
| testuser    | test123    | ❌    | ecd71870d1963316a97e3ac3408c9835ad8cf0f3c1bc703527c30265534f75ae |

## 🧪 Pruebas con curl

### 1. Login exitoso (obtener token)

```bash
curl -X POST http://localhost:8082/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "genacanas",
    "password": "genaro123"
  }'
```

**Response:**
```json
{
  "id": 1,
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJiYWNrZW5kIiwic3ViIjoiYXV0aCIsImV4cCI6MTY5OTk5OTk5OSwibmJmIjoxNjk5OTEzNTk5LCJpYXQiOjE2OTk5MTM1OTksImp0aSI6IjEiLCJlc19hZG1pbiI6dHJ1ZX0...",
  "es_admin": true
}
```

### 2. Login fallido (contraseña incorrecta)

```bash
curl -X POST http://localhost:8082/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "genacanas",
    "password": "wrongpassword"
  }'
```

**Response:**
```json
{
  "error": "invalid password"
}
```

### 3. Obtener usuario por ID (siendo el owner)

```bash
# Primero hacer login y guardar el token
TOKEN=$(curl -X POST http://localhost:8082/login \
  -H "Content-Type: application/json" \
  -d '{"username":"genacanas","password":"genaro123"}' \
  -s | jq -r '.token')

# Usar el token para obtener el usuario
curl http://localhost:8082/usuario/1 \
  -H "Authorization: Bearer $TOKEN"
```

**Response:**
```json
{
  "id": 1,
  "nombre_apellido": "Genaro Cañas",
  "username": "genacanas",
  "es_admin": true
}
```

### 4. Intentar acceder a otro usuario (sin ser admin)

```bash
# Login como usuario normal
TOKEN=$(curl -X POST http://localhost:8082/login \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"test123"}' \
  -s | jq -r '.token')

# Intentar acceder a usuario con ID 1
curl http://localhost:8082/usuario/1 \
  -H "Authorization: Bearer $TOKEN"
```

**Response (403 Forbidden):**
```json
{
  "error": "You can only modify your own resources"
}
```

### 5. Crear nuevo usuario (como admin)

```bash
# Login como admin
ADMIN_TOKEN=$(curl -X POST http://localhost:8082/login \
  -H "Content-Type: application/json" \
  -d '{"username":"genacanas","password":"genaro123"}' \
  -s | jq -r '.token')

# Crear usuario
curl -X POST http://localhost:8082/usuario \
  -H "Authorization: Bearer $ADMIN_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "nombre_apellido": "Maria González",
    "username": "mariagonzalez",
    "password": "maria123",
    "es_admin": false
  }'
```

**Response (201 Created):**
```json
{
  "id": 4,
  "nombre_apellido": "Maria González",
  "username": "mariagonzalez",
  "es_admin": false
}
```

### 6. Intentar crear usuario sin ser admin

```bash
# Login como usuario normal
USER_TOKEN=$(curl -X POST http://localhost:8082/login \
  -H "Content-Type: application/json" \
  -d '{"username":"testuser","password":"test123"}' \
  -s | jq -r '.token')

# Intentar crear usuario
curl -X POST http://localhost:8082/usuario \
  -H "Authorization: Bearer $USER_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "nombre_apellido": "Test User",
    "username": "newuser",
    "password": "pass123",
    "es_admin": false
  }'
```

**Response (403 Forbidden):**
```json
{
  "error": "Admin privileges required"
}
```

### 7. Intentar acceder sin token

```bash
curl http://localhost:8082/usuario/1
```

**Response (401 Unauthorized):**
```json
{
  "error": "Authorization header required"
}
```

## 🔒 Seguridad

- **Hashing de contraseñas**: SHA256 - ninguna contraseña se almacena en texto plano
- **Autenticación JWT**: Tokens con expiración de 24 horas
- **Autorización por roles**: 2 tipos de usuarios soportados:
  - **Normal**: Acceso a sus propios recursos
  - **Admin**: Acceso completo y operaciones de escritura
- **Middlewares de protección**:
  - `AuthMiddleware`: Valida token JWT
  - `RequireAdmin`: Verifica permisos de administrador
  - `RequireOwnerOrAdmin`: Permite acceso al owner o admin
- **CORS**: Configurado para permitir todos los orígenes (ajustar en producción)

## 🔧 Variables de Entorno

| Variable    | Descripción              | Valor por Defecto |
|-------------|--------------------------|-------------------|
| DB_HOST     | Host de MySQL            | mysql_usuarios    |
| DB_PORT     | Puerto de MySQL          | 3306              |
| DB_USER     | Usuario de la BD         | root              |
| DB_PASSWORD | Contraseña de la BD      | genagena1         |
| DB_NAME     | Nombre de la base de datos | usuarios_db       |

## 📝 Notas de Migración

Este microservicio fue extraído del backend monolítico (`backend_viejo`). Los cambios principales:

1. ✅ Eliminada toda lógica de actividades e inscripciones
2. ✅ Puerto cambiado a 8082 (evitar conflicto con otros servicios)
3. ✅ Base de datos independiente: `usuarios_db`
4. ✅ Módulo Go renombrado a `api_usuarios`
5. ✅ Logging mejorado con Logrus
6. ✅ Variables de entorno para configuración
7. ✅ Docker compose independiente

## 🚧 TODOs

- [x] Implementar endpoint GET /usuario/:id
- [x] Implementar endpoint POST /usuario
- [x] Agregar validación de tokens JWT en middleware
- [x] Middleware de autorización (admin, owner)
- [ ] Implementar endpoint PUT /usuario/:id
- [ ] Implementar endpoint DELETE /usuario/:id
- [ ] Implementar refresh tokens
- [ ] Agregar tests unitarios
- [ ] Mejorar manejo de errores con códigos específicos
- [ ] Documentación con Swagger/OpenAPI
- [ ] Rate limiting para login (prevenir brute force)
- [ ] Implementar roles y permisos granulares (RBAC)
- [ ] Logging de auditoría para operaciones sensibles
- [ ] Validación de complejidad de contraseñas

## 🤝 Comunicación con Otros Microservicios

Este microservicio NO hace llamadas HTTP a otros servicios. Otros servicios pueden:

- Validar tokens JWT generados por este servicio
- Consultar información de usuarios (cuando se implementen endpoints GET)

## 📞 Contacto y Soporte

Para reportar problemas o sugerencias, crear un issue en el repositorio.
