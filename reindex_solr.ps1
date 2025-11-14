# Script para re-indexar manualmente todas las actividades en Solr
Write-Host "Re-indexando actividades en Solr..." -ForegroundColor Green

# Obtener token de admin
$response = Invoke-RestMethod -Uri "http://localhost:8082/login" -Method Post `
    -Body '{"username":"admin","password":"admin"}' -ContentType "application/json"
$token = $response.token

# Obtener todas las actividades desde API_Actividades
Write-Host "Obteniendo actividades desde MongoDB..." -ForegroundColor Yellow
$actividadesIds = docker exec mongodb_actividades mongosh -u mongouser -p mongopass --quiet `
    --eval "db.getSiblingDB('actividades_db').actividades.find({}, {_id: 1}).toArray().map(doc => doc._id.toString())" | ConvertFrom-Json

Write-Host "Se encontraron $($actividadesIds.Count) actividades" -ForegroundColor Cyan

# Indexar cada actividad en Solr a través de API_Busquedas
$indexadas = 0
foreach ($id in $actividadesIds) {
    try {
        $actividad = Invoke-RestMethod -Uri "http://localhost:8081/actividad/$id" `
            -Headers @{"Authorization"="Bearer $token"}
        
        # Indexar directamente en Solr
        $solrDoc = @{
            id = $actividad.id
            nombre = $actividad.nombre
            descripcion = $actividad.descripcion
            profesor = $actividad.profesor
        }
        
        $solrBody = @{
            add = @{
                doc = $solrDoc
            }
        } | ConvertTo-Json -Depth 10
        
        Invoke-RestMethod -Uri "http://localhost:8983/solr/actividades/update?commit=true" `
            -Method Post -Body $solrBody -ContentType "application/json" | Out-Null
        
        Write-Host "[OK] Indexada: $($actividad.nombre)" -ForegroundColor Green
        $indexadas++
    } catch {
        Write-Host "[ERROR] No se pudo indexar actividad $id" -ForegroundColor Red
    }
}

Write-Host "`n$indexadas actividades indexadas correctamente en Solr!" -ForegroundColor Green
Write-Host "Ahora recarga localhost:3000 y deberias ver todas las actividades!" -ForegroundColor Cyan
