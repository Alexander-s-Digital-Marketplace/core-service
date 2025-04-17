package getprofilebyid

import (
	"strconv"

	profilemodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/profile_model"
	"github.com/gin-gonic/gin"
)

func GetProfileById(c *gin.Context) (int, profilemodel.Profile) {
	id := c.Query("profile_id")
	if id == "" {
		return 400, profilemodel.Profile{}
	}

	var profile profilemodel.Profile
	var code int
	var err error

	profile.Id, err = strconv.Atoi(id)
	if err != nil {
		return 400, profilemodel.Profile{}
	}
	code = profile.GetFromTable()
	if code != 200 {
		return code, profilemodel.Profile{}
	}

	return 200, profile
}
