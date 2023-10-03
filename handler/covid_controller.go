package covid

import (
	"covid/internal/response"
	covid "covid/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCovidSummary(c *gin.Context) {
	result, err := covid.GetCovidSummaryService()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, response.OK(result))
}
