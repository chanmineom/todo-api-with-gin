package main

import (
	"github.com/gin-gonic/gin"
	"github.com/chanmineom/todo-api-with-gin/controllers"
	"github.com/chanmineom/todo-api-with-gin/middleware"
	"github.com/chanmineom/todo-api-with-gin/models"
	"github.com/chanmineom/todo-api-with-gin/utils"
)

func main() {
	r := gin.Default()
	r.Use(utils.LoggerMiddleware())

	if err := utils.InitDB(); err != nil {
		panic("failed to connect database: " + err.Error())
	}

	utils.DB.AutoMigrate(&models.User{}, &models.Todo{})

	public := r.Group("/api")
	{
		public.POST("/register", controllers.UserRegister)
		public.POST("/login", controllers.UserLogin)
	}

	protected := r.Group("/api/todos")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("", controllers.CreatedTodo)
		protected.GET("", controllers.GetTodos)
		protected.GET("/:id", controllers.GetTodo)
		protected.PUT("/:id", controllers.UpdateTodo)
		protected.DELETE("/:id", controllers.DeleteTodo)
	}

	r.Run(":8080")
}