package dto

type InscripcionDto struct {
	Id         int    `json:"id"`
	UsuarioId     int    `json:"usuario_id"`
	ActividadId int    `json:"actividad_id"`
	HorarioInscripcion HorarioDto `json:"horario_inscripcion"`
	// UsuarioNombre   string `json:"usuario_nombre"`
	// ActividadNombre string `json:"actividad_nombre"`
}

type InscripcionesDto []InscripcionDto