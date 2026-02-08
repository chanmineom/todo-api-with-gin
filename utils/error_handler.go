package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 统一响应结构体
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Success 成功响应
func SuccessResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "success",
		Data:    data,
	})
}

// Error 错误响应
func ErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}

// NotFound 资源不存在
func NotFoundResponse(c *gin.Context) {
	ErrorResponse(c, 404, "resource not found")
}

// BadRequest 参数错误
func BadRequestResponse(c *gin.Context) {
	ErrorResponse(c, 400, "invalid request parameters")
}

// Unauthorized 未认证
func UnauthorizedResponse(c *gin.Context) {
	ErrorResponse(c, 401, "authentication required")
}
