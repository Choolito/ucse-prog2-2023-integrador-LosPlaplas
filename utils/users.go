package utils

import (
	"github.com/Choolito/ucse-prog2-2023-integrador-LosPlaplas/clients/responses"
	"github.com/gin-gonic/gin"
)

func SetUserInContext(c *gin.Context, user *responses.UserInfo) {
	c.Set("UserInfo", user)
}

func GetUserInfoFromContext(c *gin.Context) *responses.UserInfo {
	userInfo, _ := c.Get("UserInfo")

	user, _ := userInfo.(*responses.UserInfo)

	return user
}
