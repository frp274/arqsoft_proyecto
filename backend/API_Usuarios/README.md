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

### Usuarios (En desarrollo)

Los siguientes endpoints están planificados:
- `GET /usuario/:id` - Obtener usuario por ID
- `POST /usuario` - Crear nuevo usuario
- `PUT /usuario/:id` - Actualizar usuario
- `DELETE /usuario/:id` - Eliminar usuario

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

### Login exitoso

```bash
curl -X POST http://localhost:8082/login \
  -H "Content-Type: application/json" \
  -d '{
    "Username": "genacanas",
    "Password": "genaro123"
  }'
```

### Login fallido

```bash
curl -X POST http://localhost:8082/login \
  -H "Content-Type: application/json" \
  -d '{
    "Username": "genacanas",
    "Password": "wrongpassword"
  }'
```

## 🔒 Seguridad

- Las contraseñas se almacenan hasheadas con SHA256
- Los tokens JWT incluyen información del usuario e información de admin
- CORS configurado para permitir todos los orígenes (ajustar en producción)

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

- [ ] Implementar endpoints CRUD de usuarios
- [ ] Agregar validación de tokens JWT en middleware
- [ ] Implementar refresh tokens
- [ ] Agregar tests unitarios
- [ ] Mejorar manejo de errores
- [ ] Documentación con Swagger/OpenAPI
- [ ] Rate limiting para login
- [ ] Implementar roles y permisos granulares

## 🤝 Comunicación con Otros Microservicios

Este microservicio NO hace llamadas HTTP a otros servicios. Otros servicios pueden:

- Validar tokens JWT generados por este servicio
- Consultar información de usuarios (cuando se implementen endpoints GET)

## 📞 Contacto y Soporte

Para reportar problemas o sugerencias, crear un issue en el repositorio.
