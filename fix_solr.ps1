$ids = @(
    '69a1a0ca1b097a0339273275',
    '69a1a0ca1b097a0339273276',
    '69a1a0ca1b097a0339273277',
    '69a1a0ca1b097a0339273278',
    '69a1a0ca1b097a0339273279',
    '69a1a0ca1b097a033927327a',
    '69a1a0ca1b097a033927327b',
    '69a1a0ca1b097a033927327c',
    '69a1a0ca1b097a033927327d',
    '69a1a0ca1b097a033927327e'
)

$ok = 0
foreach ($id in $ids) {
    try {
        $act = Invoke-RestMethod "http://localhost:8081/actividad/$id"

        # Convertir horarios a JSON string (como espera Solr en este schema)
        $horariosJson = $null
        if ($act.horario -ne $null) {
            $horariosJson = ($act.horario | ConvertTo-Json -Compress)
        }

        $doc = @{
            add = @{
                doc = @{
                    id          = $act.id
                    nombre      = $act.nombre
                    descripcion = $act.descripcion
                    profesor    = $act.profesor
                    imagen_url  = $act.imagen_url
                    tags        = $act.tags
                    horarios    = $horariosJson
                }
            }
        } | ConvertTo-Json -Depth 10

        Invoke-RestMethod -Uri 'http://localhost:8983/solr/actividades/update?commit=true' `
            -Method Post -Body $doc -ContentType 'application/json' | Out-Null

        Write-Host "[OK] $($act.nombre)" -ForegroundColor Green
        $ok++
    } catch {
        Write-Host "[ERR] id=$id : $_" -ForegroundColor Red
    }
}

Write-Host "`n$ok/10 actividades indexadas en Solr!" -ForegroundColor Cyan
