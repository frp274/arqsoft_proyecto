package dto

type ActividadDto struct {
	Id          int    		`json:"id"`
	Nombre        string 	`json:"nombre"`
	Descripcion string 		`json:"descripcion"`
	Profesor   string 		`json:"profesor"`
	Cupo int 				`json:"cupo"`
}

type ActividadesDto []ActividadDto	