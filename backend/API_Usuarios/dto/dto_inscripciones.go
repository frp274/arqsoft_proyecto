package dto

type InscripcionDto struct {
	Id          int `json:"id"`
	UsuarioId   int `json:"usuario_id"`
	ActividadId int `json:"actividad_id"`
}

type InscripcionesDto []InscripcionDto
