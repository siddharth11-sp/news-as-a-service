package entity

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(
	service Service,
) *Handler {

	return &Handler{
		service: service,
	}
}

func (h *Handler) CreateEntity(
	c *gin.Context,
) {

	var req CreateEntityRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	entity, err :=
		h.service.CreateEntity(req)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.JSON(
		http.StatusCreated,
		entity,
	)
}

func (h *Handler) ListEntities(
	c *gin.Context,
) {

	entities, err :=
		h.service.ListEntities()

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		entities,
	)
}

func (h *Handler) GetEntity(
	c *gin.Context,
) {

	id := c.Param("id")

	entity, err :=
		h.service.GetEntity(id)

	if err != nil {

		c.JSON(
			http.StatusNotFound,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.JSON(
		http.StatusOK,
		entity,
	)
}
