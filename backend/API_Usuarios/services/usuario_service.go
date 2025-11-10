package services

import (
	usuarioClient "api_usuarios/clients/usuarios"
	"api_usuarios/dto"
	"api_usuarios/model"
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

func GetUsuarioById(id int) (dto.UsuarioDto, error) {
	usuario, err := usuarioClient.GetUsuarioById(id)
	if err != nil {
		log.Warnf("User not found with ID: %d", id)
		return dto.UsuarioDto{}, fmt.Errorf("user not found")
	}

	// Mapear modelo a DTO (sin devolver el hash de contraseña)
	usuarioDto := dto.UsuarioDto{
		Id:              usuario.Id,
		Nombre_apellido: usuario.Nombre_apellido,
		UserName:        usuario.UserName,
		Es_admin:        usuario.Es_admin,
	}

	log.Infof("User retrieved: %s (ID: %d)", usuario.UserName, usuario.Id)
	return usuarioDto, nil
}

func CreateUsuario(usuarioDto dto.CreateUsuarioRequest) (dto.UsuarioDto, error) {
	// Validaciones
	if usuarioDto.UserName == "" {
		return dto.UsuarioDto{}, fmt.Errorf("username is required")
	}
	if usuarioDto.Password == "" {
		return dto.UsuarioDto{}, fmt.Errorf("password is required")
	}
	if usuarioDto.Nombre_apellido == "" {
		return dto.UsuarioDto{}, fmt.Errorf("nombre_apellido is required")
	}

	// Hashear la contraseña
	passwordHash := utils.HashSHA256(usuarioDto.Password)

	// Crear modelo
	usuario := model.Usuario{
		Nombre_apellido: usuarioDto.Nombre_apellido,
		UserName:        usuarioDto.UserName,
		Es_admin:        usuarioDto.Es_admin,
		PasswordHash:    passwordHash,
	}

	// Guardar en la base de datos
	usuarioCreado, err := usuarioClient.CreateUsuario(usuario)
	if err != nil {
		log.Errorf("Error creating user: %v", err)
		return dto.UsuarioDto{}, fmt.Errorf("error creating user: %w", err)
	}

	// Mapear a DTO
	responseDto := dto.UsuarioDto{
		Id:              usuarioCreado.Id,
		Nombre_apellido: usuarioCreado.Nombre_apellido,
		UserName:        usuarioCreado.UserName,
		Es_admin:        usuarioCreado.Es_admin,
	}

	log.Infof("User created: %s (ID: %d)", usuarioCreado.UserName, usuarioCreado.Id)
	return responseDto, nil
}
