package dto

type InscripcionDto struct {
	Id          int    `json:"id"`
	UsuarioId   int    `json:"usuario_id"`
	ActividadId string `json:"actividad_id"`
	HorarioId   string `json:"horario_id"`
}

type InscripcionesDto []InscripcionDto
