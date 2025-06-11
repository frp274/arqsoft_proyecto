package dto

type ActividadDto struct {
	Id          int    		`json:"id"`
	Nombre        string 	`json:"nombre"`
	Descripcion string 		`json:"descripcion"`
	Profesor   string 		`json:"profesor"`
	Cupo int 				`json:"cupo"`
	Horario []HorarioDto	`json:"horarios"`
}

type ActividadesDto []ActividadDto	

type HorarioDto struct{
	Dia        string	`json:"dia"`
	HoraInicio string	`json:"horarioInicio"`
	HoraFin    string	`json:"horarioFinal"`
}