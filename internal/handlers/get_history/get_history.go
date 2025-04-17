package gethistory

import (
	historymodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/history_model"
	"github.com/gin-gonic/gin"
)

func GetHistory(c *gin.Context) (int, []historymodel.History) {
	var history historymodel.History
	var histories []historymodel.History
	var code int

	code, histories = history.GetAllFromTable()
	if code != 200 {
		return code, []historymodel.History{}
	}
	return 200, histories
}
