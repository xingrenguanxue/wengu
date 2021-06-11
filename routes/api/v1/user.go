package v1

import (
	"github.com/gin-gonic/gin"
	"wengu/models"
	"net/http"
	log "github.com/sirupsen/logrus"
)

// SignUp 用户注册
func SignUp(c *gin.Context) {
	var user models.User
	// var msg string

	_ = c.ShouldBindJSON(&user)

	log.WithFields(log.Fields{
		"username": user.Username,
		"password": user.Password,
	}).Info("日志打印注册信息")

	code := models.CreateUser(&user)

	c.JSON (
		http.StatusOK, gin.H{
			"status":  code,
			"message": "新增成功",
		},
	)

	
}