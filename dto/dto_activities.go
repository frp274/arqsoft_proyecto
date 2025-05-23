package dto

type ActivityDto struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Capacity    int    `json:"capacity"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Professor   string `json:"professor"`
}

type ActivitiesDto []ActivityDto	