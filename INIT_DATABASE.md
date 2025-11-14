# 🎯 Guía de Inicialización de Base de Datos

Este proyecto incluye **población automática** de las bases de datos cuando se ejecuta por primera vez en una máquina nueva.

## 📋 ¿Qué se carga automáticamente?

### 🔐 MySQL (API_Usuarios)
**Usuarios pre-cargados:**
- **Admin**: 
  - Username: `admin`
  - Password: `admin`
  - Email: `admin@gym.com`
  
- **Usuarios de prueba** (password: `password123`):
  - `juan_gomez` - juan.gomez@email.com
  - `maria_lopez` - maria.lopez@email.com
  - `carlos_diaz` - carlos.diaz@email.com

### 🏃 MongoDB (API_Actividades)
**10 Actividades pre-cargadas:**
1. Yoga Integral
2. CrossFit Avanzado
3. Spinning Indoor
4. Pilates Mat
5. Entrenamiento Funcional
6. Zumba Fitness
7. Boxeo Fitness
8. Natación Intermedia
9. Stretching y Flexibilidad
10. GAP (Glúteos-Abdomen-Piernas)

Cada actividad incluye múltiples horarios con días, horas y cupos.

## 🚀 Primer Inicio (Máquina Nueva)

### Opción 1: Inicio Limpio (Recomendado)

```powershell
# 1. Limpiar volúmenes existentes (si los hay)
docker-compose down -v

# 2. Iniciar servicios
docker-compose up -d

# 3. Esperar a que los servicios estén listos (30-60 segundos)
Start-Sleep -Seconds 60

# 4. Verificar que todo está funcionando
docker-compose ps
```

### Opción 2: Reiniciar Solo las Bases de Datos

```powershell
# Eliminar solo los volúmenes de bases de datos
docker volume rm arqsoft_proyecto_mysql_usuarios_data
docker volume rm arqsoft_proyecto_mongodb_actividades_data

# Reiniciar servicios
docker-compose up -d mysql_usuarios mongodb_actividades
```

## ✅ Verificación

### 1. Verificar MySQL
```powershell
docker exec mysql_usuarios mysql -uroot -proot -e "SELECT username, email FROM usuarios_db.usuario;"
```

**Resultado esperado:**
```
username        email
admin           admin@gym.com
juan_gomez      juan.gomez@email.com
maria_lopez     maria.lopez@email.com
carlos_diaz     carlos.diaz@email.com
```

### 2. Verificar MongoDB
```powershell
docker exec mongodb_actividades mongosh -u mongouser -p mongopass --authenticationDatabase admin --eval "db.getSiblingDB('actividades_db').actividades.countDocuments()"
```

**Resultado esperado:**
```
10
```

### 3. Verificar Solr (debe sincronizarse automáticamente)
```powershell
Invoke-RestMethod -Uri "http://localhost:8983/solr/actividades/select?q=*:*&rows=0" | Select-Object -ExpandProperty response | Select-Object numFound
```

**Resultado esperado:**
```
numFound
--------
      10
```

## 🔧 Problemas Comunes

### ❌ "Las actividades no aparecen en el frontend"

**Causa:** Solr no se ha sincronizado desde MongoDB.

**Solución:**
1. Verificar que RabbitMQ está corriendo:
   ```powershell
   docker logs rabbitmq
   ```

2. Verificar que API_Busquedas está consumiendo eventos:
   ```powershell
   docker logs api_busquedas
   ```

3. Forzar re-indexación ejecutando el script de población manual:
   ```powershell
   .\populate_db_fixed.ps1
   ```

### ❌ "Error de autenticación con admin"

**Causa:** La base de datos se creó antes de añadir el script de inicialización.

**Solución:**
```powershell
# Recrear volumen de MySQL
docker-compose down
docker volume rm arqsoft_proyecto_mysql_usuarios_data
docker-compose up -d
```

### ❌ "MongoDB no tiene actividades"

**Causa:** El script de inicialización no se ejecutó.

**Solución:**
```powershell
# Ejecutar script manualmente
docker exec mongodb_actividades mongosh -u mongouser -p mongopass --authenticationDatabase admin /docker-entrypoint-initdb.d/init-mongo.js

# O recrear el volumen
docker-compose down
docker volume rm arqsoft_proyecto_mongodb_actividades_data
docker-compose up -d
```

## 📁 Archivos de Inicialización

- **MySQL**: `backend/API_Usuarios/db/init/01-seed-usuarios.sql`
- **MongoDB**: `backend/API_Actividades/db/init-mongo.js`

Estos archivos se ejecutan automáticamente cuando Docker crea los volúmenes por primera vez.

## 🔄 Repoblar Bases de Datos

Si necesitas **repoblar las bases de datos con datos frescos**:

```powershell
# Opción A: Usando el script de PowerShell (más control)
.\populate_db_fixed.ps1

# Opción B: Recreando todo desde cero
docker-compose down -v
docker-compose up -d
```

## 📊 Datos Incluidos

### Usuarios (4 total)
- 1 admin
- 3 usuarios normales

### Actividades (10 total)
- 26 horarios totales distribuidos de Lunes a Sábado
- Cupos variados (10-30 personas)
- Horarios desde 06:00 hasta 21:00

### Inscripciones
- Inicialmente vacías (se crean mediante el frontend)

---

**🎉 Con esta configuración, cualquier persona que clone el repositorio y ejecute `docker-compose up -d` tendrá un sistema completamente funcional con datos de prueba.**
