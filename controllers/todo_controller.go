package controllers

import (
	"strconv"

	"github.com/chanmineom/todo-api-with-gin/middleware"
	"github.com/chanmineom/todo-api-with-gin/models"
	"github.com/chanmineom/todo-api-with-gin/utils"
	"github.com/gin-gonic/gin"
)

// Register 用户注册
func UserRegister(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.BadRequestResponse(c)
		return
	}

	// 创建用户
	result := utils.DB.Create(&user)
	if result.Error != nil {
		utils.ErrorResponse(c, 500, "failed to create user: " + result.Error.Error())
		return
	}

	utils.SuccessResponse(c, gin.H{"user_id": user.ID, "username": user.Username})
}

func UserLogin(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		utils.BadRequestResponse(c)
		return
	}

	// 查询用户
	var existingUser models.User
	result := utils.DB.Where("username = ?", user.Username).First(&existingUser)
	if result.Error != nil || existingUser.Password != user.Password {
		utils.ErrorResponse(c, 401, "invalid username or password")
		return
	}

	// 生成令牌
	token, err := middleware.GenerateToken(existingUser.ID)
	if err != nil {
		utils.ErrorResponse(c, 500, "failed to generate token")
		return
	}

	utils.SuccessResponse(c, gin.H{"token": token, "user_id": existingUser.ID})
}

// CreateTodo 创建待办事项
func CreatedTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		utils.BadRequestResponse(c)
		return
	}

	// 获取当前登录用户ID
	userID, _ := c.Get("user_id")
	todo.UserID = userID.(uint)

	result := utils.DB.Create(&todo)
	if result.Error != nil {
		utils.ErrorResponse(c, 500, "failed to create todo: " + result.Error.Error())
		return
	}

	utils.SuccessResponse(c, todo)
}

// GetTodos 获取当前用户的所有待办事项
func GetTodos(c *gin.Context) {
	userID, _ := c.Get("user_id")
	var todos []models.Todo
	utils.DB.Where("user_id = ?", userID).Find(&todos)
	utils.SuccessResponse(c, todos)
}

// GetTodo 获取单个待办事项
func GetTodo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequestResponse(c)
		return
	}

	userID, _ := c.Get("user_id")
	var todo models.Todo
	result := utils.DB.Where("id = ? AND user_id = ?", id, userID).First(&todo)
	if result.Error != nil {
		utils.NotFoundResponse(c)
		return
	}

	utils.SuccessResponse(c, todo)
}

// UpdateTodo 更新待办事项
func UpdateTodo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequestResponse(c)
		return
	}

	userID, _ := c.Get("user_id")
	var todo models.Todo
	result := utils.DB.Where("id = ? AND user_id = ?", id, userID).First(&todo)
	if result.Error != nil {
		utils.NotFoundResponse(c)
		return
	}

	var updatedTodo models.Todo
	if err = c.ShouldBindJSON(&updatedTodo); err != nil {
		utils.BadRequestResponse(c)
		return
	}

	todo.Title = updatedTodo.Title
	todo.Description = updatedTodo.Description
	todo.IsCompleted = updatedTodo.IsCompleted
	utils.DB.Save(&todo)

	utils.SuccessResponse(c, todo)
}

// DeleteTodo 删除待办事项
func DeleteTodo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		utils.BadRequestResponse(c)
		return
	}

	userID, _ := c.Get("user_id")
	result := utils.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Todo{})
	if result.RowsAffected == 0 {
		utils.NotFoundResponse(c)
		return
	}

	utils.SuccessResponse(c, gin.H{"message": "todo deleted successfully"})
}

