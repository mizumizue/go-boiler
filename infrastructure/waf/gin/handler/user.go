package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func FetchUsers(gin *gin.Context) {
	gin.PureJSON(http.StatusOK, "Hoge Hoge Hoge")
}
