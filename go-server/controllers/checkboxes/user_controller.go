package checkboxes

import (
	"net/http"
	"one-million-checkboxes/models"

	bitSets "one-million-checkboxes/services/bit_sets"

	"github.com/gin-gonic/gin"
)

// UpdateBitSet godoc
// @Summary Update the bitset at a given position
// @Description Update the bitset at a given position
// @ID update-bitset
// @Accept  json
// @Produce  json
// @Param data body models.UpdateRequest true "Update Request"
// @Success 200 {object} models.UpdateResponse
// @Router /api/v1/checkboxes/update [post]
func UpdateBitSet(c *gin.Context) {
	var req models.UpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bitSets.UpdateBitSet(req.Position, req.Value)
	c.JSON(http.StatusOK, models.UpdateResponse{Ok: true, StatusCode: http.StatusOK})
}

// GetCurrentBitSet godoc
// @Summary Get the current bitset
// @Description Get the current bitset
// @ID get-current-bitset
// @Produce  json
// @Success 200 {object} models.GetCurrentResponse
// @Router /api/v1/checkboxes/current [get]
func GetCurrentBitSet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"bitset": bitSets.GetCurrentBitSet()})
}

// WipeBitSet godoc
// @Summary Wipe the current bitset
// @Description Wipe the current bitset
// @ID wipe-bitset
// @Produce  json
// @Success 200 {object} models.GetCurrentResponse
// @Router /api/v1/checkboxes/wipe [get]
func WipeBitSet(c *gin.Context) {
	bitSets.WipeBitSet()
	c.JSON(http.StatusOK, gin.H{"bitset": bitSets.GetCurrentBitSet()})
}
