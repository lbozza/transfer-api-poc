package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"main/internal/core/domain"
	"main/internal/core/ports"
	"net/http"
)

type HTTPHandler struct {
	transferService ports.TransferService
}

func NewHTTPHandler(transferService ports.TransferService) *HTTPHandler {
	return &HTTPHandler{
		transferService: transferService,
	}
}

func (h *HTTPHandler) Create(c *gin.Context) {
	var req domain.TransferCreateParams

	if err := c.ShouldBindBodyWith(&req, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	transfer, err := h.transferService.Create(req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusOK, transfer)
}
