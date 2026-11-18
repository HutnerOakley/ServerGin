package user_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *UserApi) UserList(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
