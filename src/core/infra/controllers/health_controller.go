package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getting(c *gin.Context){
    c.JSONP(http.StatusOK, gin.H{
        "message": "200",
    })
}
