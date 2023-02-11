package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userIdCtx           = "userId"
	userRoleCtx         = "userRoleId"

	roleIDUser      = 0
	roleIDModerator = 1
	roleIDAdmin     = 2
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)

	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
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

func (h *Handler) checkRightsForEdit(c *gin.Context) {
	userID, roleID, err := getUser(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	userIDParam, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid param")
		return
	}

	if userID != userIDParam && roleID == roleIDUser {
		newErrorResponse(c, http.StatusForbidden, "you haven't access rights")
		return
	}
}

func (h *Handler) checkOnModerator(c *gin.Context) {
	_, roleID, err := getUser(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	if roleID == roleIDUser {
		newErrorResponse(c, http.StatusForbidden, "you are not moderator/admin")
		return
	}

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
