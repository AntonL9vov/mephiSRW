package handler

import (
	"github.com/gin-gonic/gin"
	"mephiSRW/pkg/model"
)

func (h *Handler) signUp(c *gin.Context) {
	var input model.User

	if err := c.BindJSON(&input); err != nil {
		
	}
}

func (h *Handler) signIn(c *gin.Context) {}
