package dto

type InscriptionDto struct {
	Id         int    `json:"id"`
	Fecha      string `json:"fecha"`
	UserId     int    `json:"user_id"`
	ActivityId int    `json:"activity_id"`
	UserName   string `json:"user_name"`
	ActivityName string `json:"activity_name"`
}
type InscriptionsDto []InscriptionDto