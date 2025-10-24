package dto

type InscripcionDto struct {
	Id         		int   	`json:"id"`
	UsuarioId     	int    	`json:"usuario_id"`
	ActividadId 	int    	`json:"actividad_id"`
	HorarioId 		int		`json:"horario_id"`
	// UsuarioNombre   string `json:"usuario_nombre"`
	// ActividadNombre string `json:"actividad_nombre"`
}

type InscripcionesDto []InscripcionDto