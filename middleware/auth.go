package middleware

import (
	"github.com/gin-gonic/gin"
	"spike-mc-ops/config"
	"spike-mc-ops/response"
)

func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var params authParams
		if err := c.ShouldBindHeader(&params); err != nil {
			response.FailWithNoAuth("header: api_key is required", c)
			c.Abort()
		} else {
			if params.Token != config.Cfg.System.Token {
				response.FailWithNoAuth("header: api_key doesn't exist", c)
				c.Abort()
			}
		}
		c.Next()
	}
}

type authParams struct {
	Token string `header:"token" binding:"required"`
}
