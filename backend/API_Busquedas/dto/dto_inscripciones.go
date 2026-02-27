package dto

type InscripcionDto struct {
	Id          int    `json:"id"`
	UsuarioId   int    `json:"usuario_id"`
	ActividadId string `json:"actividad_id"`
	HorarioId   string `json:"horario_id"`
	Nombre      string `json:"nombre,omitempty"`
	Descripcion string `json:"descripcion,omitempty"`
	Profesor    string `json:"profesor,omitempty"`
}

type InscripcionesDto []InscripcionDto
