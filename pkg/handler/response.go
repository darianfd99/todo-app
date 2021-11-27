package handler

import (
	todo "github.com/darianfd99/todo-app/pkg"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})

}

type StatusResponse struct {
	Status string `json:"status"`
}

type getAllListResponse struct {
	Data []todo.TodoList `json:"data"`
}

type getAllItemResponse struct {
	Data []todo.TodoItem `json:"data"`
}
