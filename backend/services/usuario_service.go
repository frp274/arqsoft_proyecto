package services

import (
	usuarioClient "arqsoft_proyecto/clients/usuarios"
	"arqsoft_proyecto/utils"
	"fmt"
	"log"
)


func Login(username string, password string) (int, string, error){
	usuario, err := usuarioClient.GetUsuarioByUsername(username)

	if err !=  nil{
		log.Print("service if 1")
		return 0, "", fmt.Errorf("error getting user: %w", err)
	}

	if utils.HashSHA256(password) != usuario.PasswordHash{
		log.Print("service if 2")
		return 0, "", fmt.Errorf("invalid Password")
	}
	token, err := utils.GenerateJWT(usuario.Id, usuario.Es_admin)
	if err != nil {
		log.Print("service if 3")
		return 0, "", fmt.Errorf("error generating token: %w", err)
	}
	return usuario.Id, token, nil
}