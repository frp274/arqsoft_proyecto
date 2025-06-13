package usuarios

import (
	"arqsoft_proyecto/model"
	"fmt"

	"gorm.io/gorm"
)

var Db *gorm.DB

func GetUsuarioByUsername(username string) (model.Usuario, error) {
	var usuario model.Usuario
	txn := Db.First(&usuario, "Username = ?", username)
	if txn.Error != nil {
		return model.Usuario{}, fmt.Errorf("error getting user: %w", txn.Error)
	}
	return usuario, nil
}


