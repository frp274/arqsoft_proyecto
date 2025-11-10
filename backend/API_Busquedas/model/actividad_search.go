package model

// ActividadSolr representa una actividad indexada en Solr
type ActividadSolr struct {
	ID          string   `json:"id"`
	Nombre      string   `json:"nombre"`
	Descripcion string   `json:"descripcion,omitempty"`
	Profesor    string   `json:"profesor"`
	Tags        []string `json:"tags,omitempty"`
	Horarios    string   `json:"horarios"` // JSON string de los horarios
	CreatedAt   string   `json:"created_at,omitempty"`
	UpdatedAt   string   `json:"updated_at,omitempty"`
}

// ActividadSearchResult representa el resultado de una búsqueda
type ActividadSearchResult struct {
	Actividades []ActividadSolr `json:"actividades"`
	Total       int64           `json:"total"`
	Page        int             `json:"page"`
	PageSize    int             `json:"page_size"`
	TotalPages  int             `json:"total_pages"`
}

// SearchParams representa los parámetros de búsqueda
type SearchParams struct {
	Query    string            `form:"q"`
	Page     int               `form:"page" binding:"min=1"`
	PageSize int               `form:"size" binding:"min=1,max=100"`
	Sort     string            `form:"sort"`
	Order    string            `form:"order"` // asc o desc
	Filters  map[string]string `form:"filters"`
}
