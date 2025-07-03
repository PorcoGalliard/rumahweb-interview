package handler

import (
	"fmt"
	"net/http"

	"github.com/PorcoGalliard/rumahweb-interview/cmd/services"
	"github.com/PorcoGalliard/rumahweb-interview/models"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService services.UserServices
}

func NewUserHandler(userService *services.UserServices) *UserHandler {
	return &UserHandler{
		UserService: *userService,
	}
}

func (h *UserHandler) Register(c *gin.Context) {
	var param models.RegisterParameter
	if err := c.ShouldBindJSON(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": err.Error(),
		})
		return
	}

	if len(param.Password) < 8 ||
	len(param.ConfirmPassword) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Password must more than 8 characters",
		})
		return
	}

	if param.Password != param.ConfirmPassword {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Password and Confirm Password Not Match",
		})
		return
	}

	user, err := h.UserService.GetUserByEmail(c.Request.Context() , param.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error_message": err.Error(),
		})
		return
	}

	if user != nil && user.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Email already exist",
		})
		return
	}

	err = h.UserService.RegisterUser(c.Request.Context(), &models.User{
		Name: param.Name,
		Email: param.Email,
		Password: param.Password,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error_message": err.Error(),
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User successfully registered",
	})
}

func (h *UserHandler) Login(c *gin.Context) {
	var params *models.LoginParameter
	if err := c.ShouldBindJSON(&params); err != nil {
		return
	}

	if len(params.Password) < 8 {
		return
	}

	token, err := h.UserService.LoginUser(c.Request.Context(), params)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (h *UserHandler) GetAllUser(c *gin.Context) {
	users, err := h.UserService.GetAllUser(c.Request.Context())
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	var param *models.User
	if err := c.ShouldBindJSON(&param); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Invalid input",
		})
		return
	}

	if param.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
		"error_message": "Invalid request",			
		})
	return
	}

	user, err := h.UserService.UpdateUser(c.Request.Context(), param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
				"error_message": err,
			})
			return
	}

	c.JSON(http.StatusOK, gin.H{
		"message" : "success updating user",
		"user": user,
	})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	var param *models.User
	if err := c.ShouldBindJSON(&param); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Invalid input",
		})
		return
	}

	if param.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Invalid request",
		})
		return
	}

	if err := h.UserService.DeleteUser(c.Request.Context(), param.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
		"error_message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Successfully deleted product %s", param.Name),
	})
	return
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	var param *models.User
	if err := c.ShouldBindJSON(&param); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Invalid input",
		})
		return
	}

	if param.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Invalid request",
		})
		return
	}

	user, err := h.UserService.GetUserByID(c.Request.Context(), param.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error_message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}