package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marciovasconcelosjr/gopportunities/schemas"
)

// @BasePath /api/v1

// @Summary Get openings
// @Description Get a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Sucess 200 {object} ListOpeningsResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSuccess(ctx, "list-openings", openings)
}
