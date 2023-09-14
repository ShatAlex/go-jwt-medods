package handler

import (
	"net/http"

	tokens "github.com/ShatALex/TestTaskBackDev"
	"github.com/gin-gonic/gin"
)

// @Summary SignUp
// @Tags auth
// @Description endpoint for creating account
// @ID sign-up
// @Accept json
// @Produce json
// @Param input body tokens.SignUpUser true "account fields"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up [post]
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

// @Summary SignIn
// @Tags auth
// @Description endpoint for login
// @ID sign-in
// @Accept json
// @Produce json
// @Param input body tokens.SignInUser true "account fields"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-in [post]
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

// @Summary Refresh tokens
// @Security ApiKeyAuth
// @Tags Tokens
// @Description refresh tokens
// @ID refresh-tokens
// @Accept  json
// @Produce  json
// @Param input body tokens.SwaggerRefresh true "token"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /tokens/refresh [post]
func (h *Handler) refresh(c *gin.Context) {

	var input refreshInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	guid, err := getUserGuid(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, "errors when getting GUID")
		return
	}

	if ok := h.services.ValidateRefreshToken(c.Request.Context(), input.RefreshToken, guid); !ok {
		newErrorResponse(c, http.StatusInternalServerError, "invalid refreshToken")
		return
	}

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
