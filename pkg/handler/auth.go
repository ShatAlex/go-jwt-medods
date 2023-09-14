package handler

import (
	"net/http"

	tokens "github.com/ShatALex/TestTaskBackDev"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input tokens.SignUpUser

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	guid, err := h.services.CreateUser(c.Request.Context(), input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"guid": guid,
	})
}

func (h *Handler) signIn(c *gin.Context) {

	var input tokens.SignInUser

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	accessToken, refreshToken, err := h.services.GenerateTokens(c.Request.Context(), input.Guid)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

type refreshInput struct {
	RefreshToken string `json:"refresh_token" bson:"refreshtoken"`
}

func (h *Handler) refresh(c *gin.Context) {

	var input refreshInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	guid, err := h.services.TakeGuidByRefToken(c.Request.Context(), input.RefreshToken)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	accessToken, refreshToken, err := h.services.GenerateTokens(c.Request.Context(), guid)

	if err != nil {

		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
