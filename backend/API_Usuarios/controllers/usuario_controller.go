package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"api_usuarios/dto"
	"api_usuarios/model"

	"github.com/gin-gonic/gin"
)

// Mantenemos la interface como está
type Service interface {
	GetAll() ([]model.Usuario, error)
	GetByID(id int) (model.Usuario, error)
	Create(user model.Usuario) (int, error)
	Update(user model.Usuario) error
	Delete(id int) error
	Login(username string, password string) (dto.LoginResponse, error)
}

type Controller struct {
	service Service
}

func NewController(service Service) Controller {
	return Controller{
		service: service,
	}
}

func (controller Controller) GetAll(c *gin.Context) {
	usuarios, err := controller.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("error getting all users: %s", err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, usuarios)
}

func (controller Controller) GetByID(c *gin.Context) {
	userID := c.Param("id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	usuario, err := controller.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("usuario no encontrado: %s", err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, usuario)
}

func (controller Controller) Create(c *gin.Context) {
	var usuario model.Usuario
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("datos inválidos: %s", err.Error()),
		})
		return
	}

	id, err := controller.service.Create(usuario)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("error creating user: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}

func (controller Controller) Update(c *gin.Context) {
	userID := c.Param("id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	var usuario model.Usuario
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("datos inválidos: %s", err.Error()),
		})
		return
	}

	usuario.Id = id
	if err := controller.service.Update(usuario); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("error actualizando usuario: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, usuario)
}

func (controller Controller) Delete(c *gin.Context) {
	userID := c.Param("id")
	id, err := strconv.Atoi(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID inválido",
		})
		return
	}

	if err := controller.service.Delete(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("error eliminando usuario: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"mensaje": "Usuario eliminado correctamente",
	})
}

func (controller Controller) Login(c *gin.Context) {
	var loginRequest dto.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("datos inválidos: %s", err.Error()),
		})
		return
	}

	response, err := controller.service.Login(loginRequest.Username, loginRequest.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": fmt.Sprintf("error en login: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}
