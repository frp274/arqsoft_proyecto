# Script para poblar las bases de datos con datos de prueba
Write-Host "🚀 Poblando bases de datos..." -ForegroundColor Green

# URL de las APIs
$API_USUARIOS = "http://localhost:8082"
$API_ACTIVIDADES = "http://localhost:8081"

# 1. Login como admin para obtener el token
Write-Host "`n1. Obteniendo token de admin..." -ForegroundColor Yellow
$adminLogin = @{
    username = "admin"
    password = "admin"
} | ConvertTo-Json

try {
    $response = Invoke-RestMethod -Uri "$API_USUARIOS/login" -Method Post -Body $adminLogin -ContentType "application/json"
    $ADMIN_TOKEN = $response.token
    $ADMIN_ID = $response.id
    Write-Host "✅ Admin ID: $ADMIN_ID" -ForegroundColor Green
    Write-Host "✅ Token obtenido" -ForegroundColor Green
} catch {
    Write-Host "❌ Error al obtener token de admin: $_" -ForegroundColor Red
    exit 1
}

# 2. Crear usuarios normales
Write-Host "`n2. Creando usuarios normales..." -ForegroundColor Yellow

$usuarios = @(
    @{username="juan_gomez"; email="juan.gomez@email.com"; nombre="Juan"; apellido="Gomez"; password="password123"; es_admin=$false},
    @{username="maria_lopez"; email="maria.lopez@email.com"; nombre="Maria"; apellido="Lopez"; password="password123"; es_admin=$false},
    @{username="carlos_diaz"; email="carlos.diaz@email.com"; nombre="Carlos"; apellido="Diaz"; password="password123"; es_admin=$false}
)

foreach ($usuario in $usuarios) {
    try {
        $body = $usuario | ConvertTo-Json
        Invoke-RestMethod -Uri "$API_USUARIOS/usuario" -Method Post -Body $body -ContentType "application/json" | Out-Null
        Write-Host "✅ Usuario: $($usuario.username)" -ForegroundColor Green
    } catch {
        Write-Host "⚠️  Usuario $($usuario.username) ya existe o error: $($_.Exception.Message)" -ForegroundColor Yellow
    }
}

# 3. Crear actividades variadas
Write-Host "`n3. Creando actividades..." -ForegroundColor Yellow

$actividades = @(
    @{
        nombre = "Yoga Integral"
        descripcion = "Clases de yoga para todos los niveles. Mejora tu flexibilidad, fuerza y equilibrio mental."
        profesor = "Laura Martinez"
        owner_id = $ADMIN_ID
        horarios = @(
            @{dia="Lunes"; horarioInicio="08:00"; horarioFinal="09:00"; cupo=15},
            @{dia="Miércoles"; horarioInicio="08:00"; horarioFinal="09:00"; cupo=15},
            @{dia="Viernes"; horarioInicio="18:00"; horarioFinal="19:00"; cupo=15}
        )
    },
    @{
        nombre = "CrossFit Avanzado"
        descripcion = "Entrenamiento de alta intensidad combinando levantamiento de pesas, gimnasia y cardio."
        profesor = "Roberto Sanchez"
        owner_id = $ADMIN_ID
        horario = @(
            @{dia="Lunes"; horaInicio="19:00"; horaFin="20:00"; cupo=20},
            @{dia="Miércoles"; horaInicio="19:00"; horaFin="20:00"; cupo=20},
            @{dia="Viernes"; horaInicio="19:00"; horaFin="20:00"; cupo=20}
        )
    },
    @{
        nombre = "Spinning Indoor"
        descripcion = "Entrenamiento cardiovascular sobre bicicleta estática con música motivadora."
        profesor = "Ana Rodriguez"
        owner_id = $ADMIN_ID
        horario = @(
            @{dia="Martes"; horaInicio="07:00"; horaFin="08:00"; cupo=25},
            @{dia="Jueves"; horaInicio="07:00"; horaFin="08:00"; cupo=25},
            @{dia="Sábado"; horaInicio="09:00"; horaFin="10:00"; cupo=25}
        )
    },
    @{
        nombre = "Pilates Mat"
        descripcion = "Método de ejercicio que enfatiza el equilibrio, la postura y la respiración para fortalecer el core."
        profesor = "Sofia Fernandez"
        owner_id = $ADMIN_ID
        horario = @(
            @{dia="Martes"; horaInicio="10:00"; horaFin="11:00"; cupo=12},
            @{dia="Jueves"; horaInicio="10:00"; horaFin="11:00"; cupo=12}
        )
    },
    @{
        nombre = "Entrenamiento Funcional"
        descripcion = "Ejercicios que imitan movimientos cotidianos para mejorar la funcionalidad y prevenir lesiones."
        profesor = "Diego Torres"
        owner_id = $ADMIN_ID
        horario = @(
            @{dia="Lunes"; horaInicio="18:00"; horaFin="19:00"; cupo=18},
            @{dia="Miércoles"; horaInicio="18:00"; horaFin="19:00"; cupo=18},
            @{dia="Viernes"; horaInicio="07:00"; horaFin="08:00"; cupo=18}
        )
    },
    @{
        nombre = "Zumba Fitness"
        descripcion = "Baile fitness con ritmos latinos que combina cardio y tonificación muscular."
        profesor = "Valentina Castro"
        owner_id = $ADMIN_ID
        horario = @(
            @{dia="Martes"; horaInicio="19:00"; horaFin="20:00"; cupo=30},
            @{dia="Jueves"; horaInicio="19:00"; horaFin="20:00"; cupo=30},
            @{dia="Sábado"; horaInicio="10:00"; horaFin="11:00"; cupo=30}
        )
    },
    @{
        nombre = "Boxeo Fitness"
        descripcion = "Entrenamiento de boxeo para mejorar resistencia cardiovascular, coordinación y fuerza."
        profesor = "Martin Ruiz"
        owner_id = $ADMIN_ID
        horario = @(
            @{dia="Lunes"; horaInicio="20:00"; horaFin="21:00"; cupo=16},
            @{dia="Miércoles"; horaInicio="20:00"; horaFin="21:00"; cupo=16},
            @{dia="Viernes"; horaInicio="20:00"; horaFin="21:00"; cupo=16}
        )
    },
    @{
        nombre = "Natación Intermedia"
        descripcion = "Clases de natación para mejorar técnica y resistencia en los diferentes estilos."
        profesor = "Paula Medina"
        owner_id = $ADMIN_ID
        horario = @(
            @{dia="Martes"; horaInicio="06:00"; horaFin="07:00"; cupo=10},
            @{dia="Jueves"; horaInicio="06:00"; horaFin="07:00"; cupo=10},
            @{dia="Sábado"; horaInicio="08:00"; horaFin="09:00"; cupo=10}
        )
    },
    @{
        nombre = "Stretching y Flexibilidad"
        descripcion = "Clases enfocadas en mejorar la flexibilidad, reducir tensión muscular y prevenir lesiones."
        profesor = "Lucia Vargas"
        owner_id = $ADMIN_ID
        horario = @(
            @{dia="Lunes"; horaInicio="12:00"; horaFin="13:00"; cupo=20},
            @{dia="Miércoles"; horaInicio="12:00"; horaFin="13:00"; cupo=20},
            @{dia="Viernes"; horaInicio="12:00"; horaFin="13:00"; cupo=20}
        )
    },
    @{
        nombre = "GAP (Glúteos, Abdomen, Piernas)"
        descripcion = "Entrenamiento localizado para tonificar y fortalecer el tren inferior y abdomen."
        profesor = "Carolina Paz"
        owner_id = $ADMIN_ID
        horario = @(
            @{dia="Martes"; horaInicio="18:00"; horaFin="19:00"; cupo=25},
            @{dia="Jueves"; horaInicio="18:00"; horaFin="19:00"; cupo=25}
        )
    }
)

$headers = @{
    "Authorization" = "Bearer $ADMIN_TOKEN"
    "Content-Type" = "application/json"
}

foreach ($actividad in $actividades) {
    try {
        $body = $actividad | ConvertTo-Json -Depth 10
        Invoke-RestMethod -Uri "$API_ACTIVIDADES/actividad" -Method Post -Body $body -Headers $headers | Out-Null
        Write-Host "✅ Actividad: $($actividad.nombre)" -ForegroundColor Green
        Start-Sleep -Milliseconds 500  # Pequeña pausa para que RabbitMQ procese
    } catch {
        Write-Host "⚠️  Error creando $($actividad.nombre): $($_.Exception.Message)" -ForegroundColor Yellow
    }
}

Write-Host "`n✨ ¡Base de datos poblada exitosamente!" -ForegroundColor Green
Write-Host "📊 Resumen:" -ForegroundColor Cyan
Write-Host "   - 1 Admin (admin/admin)" -ForegroundColor White
Write-Host "   - 3 Usuarios normales" -ForegroundColor White
Write-Host "   - 10 Actividades con múltiples horarios" -ForegroundColor White
Write-Host "`n🔍 Las actividades deberían estar indexadas en Solr automáticamente" -ForegroundColor Cyan
Write-Host "💾 Los datos están en MongoDB y PostgreSQL" -ForegroundColor Cyan
