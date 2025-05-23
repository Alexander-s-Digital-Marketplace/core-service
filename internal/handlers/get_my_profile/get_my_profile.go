package getmyprofile

import (
	profilemodel "github.com/Alexander-s-Digital-Marketplace/core-service/internal/models/profile_model"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func GetMyProfile(c *gin.Context) (int, profilemodel.Profile) {
	id, exists := c.Get("id")
	if !exists {
		return 400, profilemodel.Profile{}
	}

	profile := profilemodel.Profile{
		AccountId: int(id.(int)),
	}

	code := profile.GetFromTableByAccountId()
	if code != 200 {
		return code, profilemodel.Profile{}
	}
	logrus.Infoln("profile", profile)

	return 200, profile
}
