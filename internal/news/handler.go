package news

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func (h *Handler) ListNews(
	c *gin.Context,
) {

	entityID, err :=
		uuid.Parse(
			c.Param("id"),
		)

	if err != nil {

		c.JSON(
			400,
			gin.H{
				"error": "invalid entity id",
			},
		)

		return
	}

	page, _ :=
		strconv.Atoi(
			c.DefaultQuery(
				"page",
				"1",
			),
		)

	pageSize, _ :=
		strconv.Atoi(
			c.DefaultQuery(
				"page_size",
				"20",
			),
		)

	resp, err :=
		h.service.ListNews(
			entityID,
			c.Query("sentiment"),
			c.Query("source"),
			page,
			pageSize,
		)

	if err != nil {

		c.JSON(
			500,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.JSON(
		200,
		resp,
	)
}
