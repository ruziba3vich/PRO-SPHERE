package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/prosphere/internal/items/repo"
	"github.com/ruziba3vich/prosphere/internal/items/storage/search"
)

type (
	SearchHandler struct {
		service repo.SearchRepository
		logger  *log.Logger
	}
)

func NewSearchHandler(service repo.SearchRepository, logger *log.Logger) *SearchHandler {
	return &SearchHandler{
		service: service,
		logger:  logger,
	}
}

func (h *SearchHandler) SearchItem(c *gin.Context) {
	var req search.SearchRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Println("-- ERROR OCCURED WHILE BINDING DATA --", err.Error())
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, err := h.service.SearchElement(c, &req)
	if err != nil {
		h.logger.Println("-- ERROR OCCURED WHILE SEARCHING ELEMENT --", err.Error())
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, response)
}
