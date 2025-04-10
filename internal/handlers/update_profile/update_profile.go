package updateprofile

import (
	profilemodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/profile_model"
	"github.com/gin-gonic/gin"
)

func UpdateProfile(c *gin.Context) (int, string) {
	id, exists := c.Get("id")
	if !exists {
		return 400, "Error get id"
	}

	var profile profilemodel.Profile
	var code int
	code = profile.DecodeFromContext(c)
	if code != 200 {
		return code, "Error decode JSON"
	}
	profile.Id = id.(int)

	code = profile.UpdateInTable()
	if code != 200 {
		return code, "Error update in table"
	}

	return 200, "Success update profile"
}
