package dto

type ActividadDto struct {
	Id          string       `json:"id"`
	Nombre      string       `json:"nombre"`
	Descripcion string       `json:"descripcion"`
	Profesor    string       `json:"profesor"`
	OwnerId     int          `json:"owner_id" binding:"required"`
	Horario     []HorarioDto `json:"horarios"`
}

type ActividadesDto []ActividadDto

type HorarioDto struct {
	Dia        string `json:"dia"`
	HoraInicio string `json:"horarioInicio"`
	HoraFin    string `json:"horarioFinal"`
	Cupo       int    `json:"cupo"`
}

// DisponibilidadResult representa el resultado del cálculo de disponibilidad
type DisponibilidadResult struct {
	Dia                string `json:"dia"`
	Cupo               int    `json:"cupo"`
	Disponibles        int    `json:"disponibles"`
	PorcentajeOcupado  float64 `json:"porcentaje_ocupado"`
}

// DisponibilidadResponse es la respuesta del endpoint de disponibilidad
type DisponibilidadResponse struct {
	ActividadId string                 `json:"actividad_id"`
	Nombre      string                 `json:"nombre"`
	Horarios    []DisponibilidadResult `json:"horarios"`
	TotalCupo   int                    `json:"total_cupo"`
	TotalOcupado int                   `json:"total_ocupado"`
	Tiempo      string                 `json:"tiempo_calculo"`
}
