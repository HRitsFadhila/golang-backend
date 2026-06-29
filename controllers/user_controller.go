package controllers

import (
	"net/http"

	"github.com/HRitsFadhila/golang-backend/database"
	"github.com/HRitsFadhila/golang-backend/helpers"
	"github.com/HRitsFadhila/golang-backend/models"
	"github.com/HRitsFadhila/golang-backend/structs"
	"github.com/gin-gonic/gin"
)

// func FindUsers(c *gin.Context){
// 	var users []models.User

// 	database.DB.Find(&users)

// 	c.JSON(http.StatusOK, structs.SuccessResponse{
// 		Success: true,
// 		Message: "Lists Data Users",
// 		Data: users,
// 	})
// }

func FindUsers(c *gin.Context){
	var users []models.User
	var response []structs.UserResponse

	database.DB.Find(&users)

	for _, user := range users{
		response = append(response, structs.UserResponse{
			Id: user.Id,
			Name: user.Name,
			Username: user.Username,
			Email: user.Email,
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "List Data Users",
		Data: response,
	})
}

func CreateUser(c *gin.Context){
	var req = structs.UserCreateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil{
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Error",
			Errors: helpers.TranslateErrorMessage(err),
		})
		return
	}

	user := models.User{
		Name: req.Name,
		Username: req.Username,
		Email: req.Email,
		Password: helpers.HashPassword(req.Password),
	}

	if err := database.DB.Create(&user).Error; err != nil{
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to create user",
			Errors: helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusCreated, structs.SuccessResponse{
		Success: true,
		Message: "User created successfully",
		Data: structs.UserResponse{
			Id: user.Id,
			Name: user.Name,
			Username: user.Username,
			Email: user.Email,
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}

func FinsUserById(c *gin.Context){
	id := c.Param("id")

	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil{
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "User not found",
			Errors: helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "User Found",
		Data: structs.UserResponse{
			Id: user.Id,
			Name: user.Name,
			Username: user.Username,
			Email: user.Email,
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}

func UpdateUser(c *gin.Context){
	id := c.Param("id")

	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil{
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "User not found",
			Errors: helpers.TranslateErrorMessage(err),
		})
		return
	}

	var req = structs.UserUpdateRequest{}

	if err := c.ShouldBindJSON(&req); err != nil{
		c.JSON(http.StatusUnprocessableEntity, structs.ErrorResponse{
			Success: false,
			Message: "Validation Errors",
			Errors: helpers.TranslateErrorMessage(err),
		})
		return
	}
	
	if req.Name != nil{
		user.Name = *req.Name
	}
	if req.Username != nil{
		user.Username = *req.Username
	}
	if req.Email != nil{
		user.Email = *req.Email
	}
	if req.Password != nil{
		user.Password = helpers.HashPassword(*req.Password)
	}

	if err := database.DB.Save(&user).Error; err != nil{
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to update user",
			Errors: helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "User updated successfully",
		Data: structs.UserResponse{
			Id: user.Id,
			Name: user.Name,
			Username: user.Username,
			Email: user.Email,
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
		},
	})
}

func DeleteUser(c *gin.Context){
	id := c.Param("id")

	var user models.User

	if err := database.DB.First(&user, id).Error; err != nil{
		c.JSON(http.StatusNotFound, structs.ErrorResponse{
			Success: false,
			Message: "User not found",
			Errors: helpers.TranslateErrorMessage(err),
		})
		return
	}

	if err := database.DB.Delete(&user).Error; err != nil{
		c.JSON(http.StatusInternalServerError, structs.ErrorResponse{
			Success: false,
			Message: "Failed to delete user",
			Errors: helpers.TranslateErrorMessage(err),
		})
		return
	}

	c.JSON(http.StatusOK, structs.SuccessResponse{
		Success: true,
		Message: "User deleted successfully",
	})
}