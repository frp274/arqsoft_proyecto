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

	token, err := utils.GenerateJWT(usuario.Id, usuario.EsAdmin)
	if err != nil {
		log.Errorf("Failed to generate JWT for user %s: %v", username, err)
		return 0, "", false, fmt.Errorf("error generating token: %w", err)
	}

	log.Infof("Login successful for user: %s (ID: %d)", username, usuario.Id)
	return usuario.Id, token, usuario.EsAdmin, nil
}

func GetUsuarioById(id int) (dto.UsuarioDto, error) {
	usuario, err := usuarioClient.GetUsuarioById(id)
	if err != nil {
		log.Warnf("User not found with ID: %d", id)
		return dto.UsuarioDto{}, fmt.Errorf("user not found")
	}

	// Mapear modelo a DTO (sin devolver el hash de contraseña)
	usuarioDto := dto.UsuarioDto{
		Id:       usuario.Id,
		Username: usuario.Username,
		Email:    usuario.Email,
		Nombre:   usuario.Nombre,
		Apellido: usuario.Apellido,
		EsAdmin:  usuario.EsAdmin,
	}

	log.Infof("User retrieved: %s (ID: %d)", usuario.Username, usuario.Id)
	return usuarioDto, nil
}

func CreateUsuario(usuarioDto dto.CreateUsuarioRequest) (dto.UsuarioDto, error) {
	// Validaciones
	if usuarioDto.Username == "" {
		return dto.UsuarioDto{}, fmt.Errorf("username is required")
	}
	if usuarioDto.Email == "" {
		return dto.UsuarioDto{}, fmt.Errorf("email is required")
	}
	if usuarioDto.Password == "" {
		return dto.UsuarioDto{}, fmt.Errorf("password is required")
	}
	if usuarioDto.Nombre == "" {
		return dto.UsuarioDto{}, fmt.Errorf("nombre is required")
	}
	if usuarioDto.Apellido == "" {
		return dto.UsuarioDto{}, fmt.Errorf("apellido is required")
	}

	// Hashear la contraseña
	passwordHash := utils.HashSHA256(usuarioDto.Password)

	// Crear modelo
	usuario := model.Usuario{
		Username:     usuarioDto.Username,
		Email:        usuarioDto.Email,
		Nombre:       usuarioDto.Nombre,
		Apellido:     usuarioDto.Apellido,
		PasswordHash: passwordHash,
		EsAdmin:      usuarioDto.EsAdmin,
	}

	// Guardar en la base de datos
	usuarioCreado, err := usuarioClient.CreateUsuario(usuario)
	if err != nil {
		log.Errorf("Error creating user: %v", err)
		return dto.UsuarioDto{}, fmt.Errorf("error creating user: %w", err)
	}

	// Mapear a DTO
	responseDto := dto.UsuarioDto{
		Id:       usuarioCreado.Id,
		Username: usuarioCreado.Username,
		Email:    usuarioCreado.Email,
		Nombre:   usuarioCreado.Nombre,
		Apellido: usuarioCreado.Apellido,
		EsAdmin:  usuarioCreado.EsAdmin,
	}

	log.Infof("User created: %s (ID: %d)", usuarioCreado.Username, usuarioCreado.Id)
	return responseDto, nil
}
