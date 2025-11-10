package usuarios

import (
	"api_usuarios/model"
	"fmt"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var Db *gorm.DB

func GetUsuarioByUsername(username string) (model.Usuario, error) {
	var usuario model.Usuario

	log.Debugf("Searching for user: %s", username)
	txn := Db.First(&usuario, "Username = ?", username)

	if txn.Error != nil {
		if txn.Error == gorm.ErrRecordNotFound {
			return model.Usuario{}, fmt.Errorf("user not found: %s", username)
		}
		return model.Usuario{}, fmt.Errorf("error getting user: %w", txn.Error)
	}

	log.Debugf("User found: %s (ID: %d)", username, usuario.Id)
	return usuario, nil
}
