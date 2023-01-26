package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userIdCtx           = "userId"
	userRoleCtx         = "userRoleId"
)

func (h Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)

	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	fmt.Println(headerParts)
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userId, userRoleId, err := h.services.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userIdCtx, userId)
	c.Set(userRoleCtx, userRoleId)
}

func getUser(c *gin.Context) (int, int, error) {
	id, ok := c.Get(userIdCtx)
	if !ok {
		return 0, 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, 0, errors.New("user id is of invalid type")
	}

	roleId, ok := c.Get(userRoleCtx)
	if !ok {
		return 0, 0, errors.New("role id not found")
	}

	roleInt, ok := roleId.(int)
	if !ok {
		return 0, 0, errors.New("user role id is of invalid type")
	}

	return idInt, roleInt, nil
}
