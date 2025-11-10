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

func GetUsuarioById(id int) (model.Usuario, error) {
	var usuario model.Usuario

	log.Debugf("Searching for user with ID: %d", id)
	txn := Db.First(&usuario, id)

	if txn.Error != nil {
		if txn.Error == gorm.ErrRecordNotFound {
			return model.Usuario{}, fmt.Errorf("user not found with ID: %d", id)
		}
		return model.Usuario{}, fmt.Errorf("error getting user: %w", txn.Error)
	}

	log.Debugf("User found: %s (ID: %d)", usuario.UserName, usuario.Id)
	return usuario, nil
}

func CreateUsuario(usuario model.Usuario) (model.Usuario, error) {
	log.Debugf("Creating new user: %s", usuario.UserName)
	
	txn := Db.Create(&usuario)
	if txn.Error != nil {
		log.Errorf("Error creating user %s: %v", usuario.UserName, txn.Error)
		return model.Usuario{}, fmt.Errorf("error creating user: %w", txn.Error)
	}

	log.Infof("User created successfully: %s (ID: %d)", usuario.UserName, usuario.Id)
	return usuario, nil
}
