package controller

import (
	"net/http"
	"strconv"

	"github.com/Nina-99/TripSpotter/backend/data/request"
	"github.com/Nina-99/TripSpotter/backend/data/response"
	"github.com/Nina-99/TripSpotter/backend/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.UserService
}

func NewUserController(service service.UserService) *UserController {
	return &UserController{userService: service}
}

// TODO: Function of Login
// TODO: Función de Inicio de Sesión
func (controller *UserController) Login(ctx *gin.Context) {
	var req request.LoginUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := controller.userService.Login(req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

// TODO: Function of user Register
// TODO: Función de Registro de usuario
func (controller *UserController) Register(ctx *gin.Context) {
	var req request.RegisterUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := controller.userService.Register(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (controller *UserController) FindAll(ctx *gin.Context) {
	users, err := controller.userService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var res []response.UserResponse
	for _, u := range users {
		res = append(res, response.UserResponse{Id: u.Id, Username: u.Username, Email: u.Email})
	}
	ctx.JSON(http.StatusOK, res)
}

func (controller *UserController) Update(ctx *gin.Context) {
	idParam := ctx.Param("id")
	userId, _ := strconv.Atoi(idParam)
	var req request.RegisterUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := controller.userService.Update(uint(userId), req); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func (controller *UserController) Delete(ctx *gin.Context) {
	idParam := ctx.Param("id")
	userId, _ := strconv.Atoi(idParam)
	if err := controller.userService.Delete(uint(userId)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "deleted"})
}

func (controller *UserController) GetUserByEmail(ctx *gin.Context) {
	Email := ctx.Param("email")
	user, err := controller.userService.FindEmail(Email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}
