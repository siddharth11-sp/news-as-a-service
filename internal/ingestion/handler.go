package ingestion

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	service *Service
}

func NewHandler(
	service *Service,
) *Handler {

	return &Handler{
		service: service,
	}
}

func (h *Handler) Refresh(
	c *gin.Context,
) {

	entityID, err :=
		uuid.Parse(
			c.Param("id"),
		)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "invalid entity id",
			},
		)

		return
	}

	go h.service.IngestEntity(entityID)

	c.JSON(
		http.StatusAccepted,
		gin.H{
			"message": "refresh started",
		},
	)
}
