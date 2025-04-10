package getmyprofile

import (
	profilemodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/profile_model"
	"github.com/gin-gonic/gin"
)

func GetMyProfile(c *gin.Context) (int, profilemodel.Profile) {
	id, exists := c.Get("id")
	if !exists {
		return 400, profilemodel.Profile{}
	}

	var profile profilemodel.Profile
	var code int

	profile.Id = id.(int)
	code = profile.GetFromTable()
	if code != 200 {
		return code, profilemodel.Profile{}
	}

	return 200, profile
}
