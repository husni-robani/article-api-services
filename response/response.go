package response

import "github.com/gin-gonic/gin"

type Response struct {
    Status  int         `json:"status"`
    Message string      `json:"message"`
    Data    interface{} `json:"data,omitempty"`
    Error   string      `json:"error,omitempty"`
}

func Success(c *gin.Context, status int, data interface{}, message string) {
    c.JSON(status, Response{
        Status:  status,
        Message: message,
        Data:    data,
    })
}

func Error(c *gin.Context, status int, message string, err string) {
    c.JSON(status, Response{
        Status:  status,
        Message: message,
        Error:   err,
    })
}