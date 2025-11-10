package services

import (
	usuarioClient "api_usuarios/clients/usuarios"
	"api_usuarios/utils"
	"fmt"

	log "github.com/sirupsen/logrus"
)

func Login(username string, password string) (int, string, bool, error) {
	usuario, err := usuarioClient.GetUsuarioByUsername(username)

	if err != nil {
		log.Warnf("User not found: %s", username)
		return 0, "", false, fmt.Errorf("error getting user: %w", err)
	}

	if utils.HashSHA256(password) != usuario.PasswordHash {
		log.Warnf("Invalid password attempt for user: %s", username)
		return 0, "", false, fmt.Errorf("invalid password")
	}

	token, err := utils.GenerateJWT(usuario.Id, usuario.Es_admin)
	if err != nil {
		log.Errorf("Failed to generate JWT for user %s: %v", username, err)
		return 0, "", false, fmt.Errorf("error generating token: %w", err)
	}

	log.Infof("Login successful for user: %s (ID: %d)", username, usuario.Id)
	return usuario.Id, token, usuario.Es_admin, nil
}
