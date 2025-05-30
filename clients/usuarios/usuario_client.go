package usuarios

import (
	"arqsoft_proyecto/model"
	"fmt"

	"gorm.io/gorm"
)

var Db *gorm.DB

func GetUsuarioByUsername(username string) (model.Usuario, error) {
	var usuario model.Usuario
	txn := Db.First(&usuario, "usuario = ?", username)
	if txn.Error != nil {
		return model.Usuario{}, fmt.Errorf("Error getting user: %w", txn.Error)
	}
	return usuario, nil
}
