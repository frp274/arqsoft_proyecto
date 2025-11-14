#!/bin/bash

# Script para poblar las bases de datos con datos de prueba
echo "🚀 Poblando bases de datos..."

# URL de las APIs
API_USUARIOS="http://localhost:8082"
API_ACTIVIDADES="http://localhost:8081"

# 1. Login como admin para obtener el token
echo "1. Obteniendo token de admin..."
ADMIN_RESPONSE=$(curl -s -X POST "$API_USUARIOS/login" \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"admin"}')

ADMIN_TOKEN=$(echo $ADMIN_RESPONSE | grep -o '"token":"[^"]*' | sed 's/"token":"//')
ADMIN_ID=$(echo $ADMIN_RESPONSE | grep -o '"id":[0-9]*' | sed 's/"id"://')

echo "✅ Admin ID: $ADMIN_ID"
echo "✅ Token obtenido"

# 2. Crear usuarios normales
echo ""
echo "2. Creando usuarios normales..."

curl -s -X POST "$API_USUARIOS/usuario" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "juan_gomez",
    "email": "juan.gomez@email.com",
    "nombre": "Juan",
    "apellido": "Gomez",
    "password": "password123",
    "es_admin": false
  }' > /dev/null
echo "✅ Usuario: juan_gomez"

curl -s -X POST "$API_USUARIOS/usuario" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "maria_lopez",
    "email": "maria.lopez@email.com",
    "nombre": "Maria",
    "apellido": "Lopez",
    "password": "password123",
    "es_admin": false
  }' > /dev/null
echo "✅ Usuario: maria_lopez"

curl -s -X POST "$API_USUARIOS/usuario" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "carlos_diaz",
    "email": "carlos.diaz@email.com",
    "nombre": "Carlos",
    "apellido": "Diaz",
    "password": "password123",
    "es_admin": false
  }' > /dev/null
echo "✅ Usuario: carlos_diaz"

# 3. Crear actividades variadas
echo ""
echo "3. Creando actividades..."

# Actividad 1: Yoga
curl -s -X POST "$API_ACTIVIDADES/actividad" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $ADMIN_TOKEN" \
  -d "{
    \"nombre\": \"Yoga Integral\",
    \"descripcion\": \"Clases de yoga para todos los niveles. Mejora tu flexibilidad, fuerza y equilibrio mental.\",
    \"profesor\": \"Laura Martinez\",
    \"owner_id\": $ADMIN_ID,
    \"horario\": [
      {\"dia\": \"Lunes\", \"horaInicio\": \"08:00\", \"horaFin\": \"09:00\", \"cupo\": 15},
      {\"dia\": \"Miércoles\", \"horaInicio\": \"08:00\", \"horaFin\": \"09:00\", \"cupo\": 15},
      {\"dia\": \"Viernes\", \"horaInicio\": \"18:00\", \"horaFin\": \"19:00\", \"cupo\": 15}
    ]
  }" > /dev/null
echo "✅ Actividad: Yoga Integral"

# Actividad 2: CrossFit
curl -s -X POST "$API_ACTIVIDADES/actividad" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $ADMIN_TOKEN" \
  -d "{
    \"nombre\": \"CrossFit Avanzado\",
    \"descripcion\": \"Entrenamiento de alta intensidad combinando levantamiento de pesas, gimnasia y cardio.\",
    \"profesor\": \"Roberto Sanchez\",
    \"owner_id\": $ADMIN_ID,
    \"horario\": [
      {\"dia\": \"Lunes\", \"horaInicio\": \"19:00\", \"horaFin\": \"20:00\", \"cupo\": 20},
      {\"dia\": \"Miércoles\", \"horaInicio\": \"19:00\", \"horaFin\": \"20:00\", \"cupo\": 20},
      {\"dia\": \"Viernes\", \"horaInicio\": \"19:00\", \"horaFin\": \"20:00\", \"cupo\": 20}
    ]
  }" > /dev/null
echo "✅ Actividad: CrossFit Avanzado"

# Actividad 3: Spinning
curl -s -X POST "$API_ACTIVIDADES/actividad" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $ADMIN_TOKEN" \
  -d "{
    \"nombre\": \"Spinning Indoor\",
    \"descripcion\": \"Entrenamiento cardiovascular sobre bicicleta estática con música motivadora.\",
    \"profesor\": \"Ana Rodriguez\",
    \"owner_id\": $ADMIN_ID,
    \"horario\": [
      {\"dia\": \"Martes\", \"horaInicio\": \"07:00\", \"horaFin\": \"08:00\", \"cupo\": 25},
      {\"dia\": \"Jueves\", \"horaInicio\": \"07:00\", \"horaFin\": \"08:00\", \"cupo\": 25},
      {\"dia\": \"Sábado\", \"horaInicio\": \"09:00\", \"horaFin\": \"10:00\", \"cupo\": 25}
    ]
  }" > /dev/null
echo "✅ Actividad: Spinning Indoor"

# Actividad 4: Pilates
curl -s -X POST "$API_ACTIVIDADES/actividad" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $ADMIN_TOKEN" \
  -d "{
    \"nombre\": \"Pilates Mat\",
    \"descripcion\": \"Método de ejercicio que enfatiza el equilibrio, la postura y la respiración para fortalecer el core.\",
    \"profesor\": \"Sofia Fernandez\",
    \"owner_id\": $ADMIN_ID,
    \"horario\": [
      {\"dia\": \"Martes\", \"horaInicio\": \"10:00\", \"horaFin\": \"11:00\", \"cupo\": 12},
      {\"dia\": \"Jueves\", \"horaInicio\": \"10:00\", \"horaFin\": \"11:00\", \"cupo\": 12}
    ]
  }" > /dev/null
echo "✅ Actividad: Pilates Mat"

# Actividad 5: Funcional
curl -s -X POST "$API_ACTIVIDADES/actividad" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $ADMIN_TOKEN" \
  -d "{
    \"nombre\": \"Entrenamiento Funcional\",
    \"descripcion\": \"Ejercicios que imitan movimientos cotidianos para mejorar la funcionalidad y prevenir lesiones.\",
    \"profesor\": \"Diego Torres\",
    \"owner_id\": $ADMIN_ID,
    \"horario\": [
      {\"dia\": \"Lunes\", \"horaInicio\": \"18:00\", \"horaFin\": \"19:00\", \"cupo\": 18},
      {\"dia\": \"Miércoles\", \"horaInicio\": \"18:00\", \"horaFin\": \"19:00\", \"cupo\": 18},
      {\"dia\": \"Viernes\", \"horaInicio\": \"07:00\", \"horaFin\": \"08:00\", \"cupo\": 18}
    ]
  }" > /dev/null
echo "✅ Actividad: Entrenamiento Funcional"

# Actividad 6: Zumba
curl -s -X POST "$API_ACTIVIDADES/actividad" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $ADMIN_TOKEN" \
  -d "{
    \"nombre\": \"Zumba Fitness\",
    \"descripcion\": \"Baile fitness con ritmos latinos que combina cardio y tonificación muscular.\",
    \"profesor\": \"Valentina Castro\",
    \"owner_id\": $ADMIN_ID,
    \"horario\": [
      {\"dia\": \"Martes\", \"horaInicio\": \"19:00\", \"horaFin\": \"20:00\", \"cupo\": 30},
      {\"dia\": \"Jueves\", \"horaInicio\": \"19:00\", \"horaFin\": \"20:00\", \"cupo\": 30},
      {\"dia\": \"Sábado\", \"horaInicio\": \"10:00\", \"horaFin\": \"11:00\", \"cupo\": 30}
    ]
  }" > /dev/null
echo "✅ Actividad: Zumba Fitness"

# Actividad 7: Boxing
curl -s -X POST "$API_ACTIVIDADES/actividad" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $ADMIN_TOKEN" \
  -d "{
    \"nombre\": \"Boxeo Fitness\",
    \"descripcion\": \"Entrenamiento de boxeo para mejorar resistencia cardiovascular, coordinación y fuerza.\",
    \"profesor\": \"Martin Ruiz\",
    \"owner_id\": $ADMIN_ID,
    \"horario\": [
      {\"dia\": \"Lunes\", \"horaInicio\": \"20:00\", \"horaFin\": \"21:00\", \"cupo\": 16},
      {\"dia\": \"Miércoles\", \"horaInicio\": \"20:00\", \"horaFin\": \"21:00\", \"cupo\": 16},
      {\"dia\": \"Viernes\", \"horaInicio\": \"20:00\", \"horaFin\": \"21:00\", \"cupo\": 16}
    ]
  }" > /dev/null
echo "✅ Actividad: Boxeo Fitness"

# Actividad 8: Natación
curl -s -X POST "$API_ACTIVIDADES/actividad" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $ADMIN_TOKEN" \
  -d "{
    \"nombre\": \"Natación Intermedia\",
    \"descripcion\": \"Clases de natación para mejorar técnica y resistencia en los diferentes estilos.\",
    \"profesor\": \"Paula Medina\",
    \"owner_id\": $ADMIN_ID,
    \"horario\": [
      {\"dia\": \"Martes\", \"horaInicio\": \"06:00\", \"horaFin\": \"07:00\", \"cupo\": 10},
      {\"dia\": \"Jueves\", \"horaInicio\": \"06:00\", \"horaFin\": \"07:00\", \"cupo\": 10},
      {\"dia\": \"Sábado\", \"horaInicio\": \"08:00\", \"horaFin\": \"09:00\", \"cupo\": 10}
    ]
  }" > /dev/null
echo "✅ Actividad: Natación Intermedia"

# Actividad 9: Stretching
curl -s -X POST "$API_ACTIVIDADES/actividad" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $ADMIN_TOKEN" \
  -d "{
    \"nombre\": \"Stretching y Flexibilidad\",
    \"descripcion\": \"Clases enfocadas en mejorar la flexibilidad, reducir tensión muscular y prevenir lesiones.\",
    \"profesor\": \"Lucia Vargas\",
    \"owner_id\": $ADMIN_ID,
    \"horario\": [
      {\"dia\": \"Lunes\", \"horaInicio\": \"12:00\", \"horaFin\": \"13:00\", \"cupo\": 20},
      {\"dia\": \"Miércoles\", \"horaInicio\": \"12:00\", \"horaFin\": \"13:00\", \"cupo\": 20},
      {\"dia\": \"Viernes\", \"horaInicio\": \"12:00\", \"horaFin\": \"13:00\", \"cupo\": 20}
    ]
  }" > /dev/null
echo "✅ Actividad: Stretching y Flexibilidad"

# Actividad 10: GAP
curl -s -X POST "$API_ACTIVIDADES/actividad" \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $ADMIN_TOKEN" \
  -d "{
    \"nombre\": \"GAP (Glúteos, Abdomen, Piernas)\",
    \"descripcion\": \"Entrenamiento localizado para tonificar y fortalecer el tren inferior y abdomen.\",
    \"profesor\": \"Carolina Paz\",
    \"owner_id\": $ADMIN_ID,
    \"horario\": [
      {\"dia\": \"Martes\", \"horaInicio\": \"18:00\", \"horaFin\": \"19:00\", \"cupo\": 25},
      {\"dia\": \"Jueves\", \"horaInicio\": \"18:00\", \"horaFin\": \"19:00\", \"cupo\": 25}
    ]
  }" > /dev/null
echo "✅ Actividad: GAP"

echo ""
echo "✨ ¡Base de datos poblada exitosamente!"
echo "📊 Resumen:"
echo "   - 1 Admin (admin/admin)"
echo "   - 3 Usuarios normales"
echo "   - 10 Actividades con múltiples horarios"
echo ""
echo "🔍 Las actividades deberían estar indexadas en Solr automáticamente"
echo "💾 Los datos están en MongoDB y PostgreSQL"
