package handlers

import (
	"net/http"

	"github.com/eac0de/atlas/sentinel/internal/services"
	"github.com/eac0de/gophkeeper/shared/pkg/httperror"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthHandlers struct {
	authService    *services.AuthService
	sessionService *services.SessionService
}

func NewAuthHandlers(
	authService *services.AuthService,
	tokenService *services.SessionService,
) *AuthHandlers {
	return &AuthHandlers{
		authService:    authService,
		sessionService: tokenService,
	}
}

func (ah *AuthHandlers) GenerateEmailCodeHandler(c *gin.Context) {
	var requestData struct {
		Email *string `json:"email"`
	}
	if err := c.BindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": err.Error()})
		return
	}
	if requestData.Email == nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": "email is required"})
		return
	}
	emailCode, err := ah.authService.GenerateEmailCode(c.Request.Context(), *requestData.Email)
	if err != nil {
		msg, statusCode := httperror.GetMessageAndStatusCode(err)
		c.JSON(statusCode, gin.H{"detail": msg})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"email_code_id": emailCode.ID})
}

func (ah *AuthHandlers) NewVerifyEmailCodeHandler(rt_path string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestData struct {
			EmailCodeID *uuid.UUID `json:"email_code_id"`
			Code        *uint16    `json:"code"`
		}
		if err := c.BindJSON(&requestData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"detail": err.Error()})
			return
		}
		if requestData.EmailCodeID == nil || requestData.Code == nil {
			c.JSON(http.StatusBadRequest, gin.H{"detail": "email_code_id and code are required"})
			return
		}
		user, isNewUser, err := ah.authService.VerifyEmailCode(c.Request.Context(), *requestData.EmailCodeID, *requestData.Code)
		if err != nil {
			msg, statusCode := httperror.GetMessageAndStatusCode(err)
			c.JSON(statusCode, gin.H{"detail": msg})
			return
		}
		tokens, err := ah.sessionService.CreateSession(c.Request.Context(), user.ID, c.GetHeader("User-Agent"), c.ClientIP())
		if err != nil {
			msg, statusCode := httperror.GetMessageAndStatusCode(err)
			c.JSON(statusCode, gin.H{"detail": msg})
			return
		}
		statusCode := http.StatusOK
		if isNewUser {
			statusCode = http.StatusCreated
		}
		c.SetCookie("atlas_rt", tokens.RefreshToken, 7*24*60*60, rt_path, "", false, true)
		c.JSON(statusCode, gin.H{"access_token": tokens.AccessToken})
	}
}

func (ah *AuthHandlers) NewRefreshTokenHandler(rt_path string) gin.HandlerFunc {
	return func(c *gin.Context) {
		refreshToken, err := c.Cookie("atlas_rt")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"detail": "Refresh token is missing or invalid"})
			return
		}
		tokens, err := ah.sessionService.UpdateSession(c.Request.Context(), refreshToken, c.GetHeader("User-Agent"), c.ClientIP())
		if err != nil {
			msg, statusCode := httperror.GetMessageAndStatusCode(err)
			c.JSON(statusCode, gin.H{"detail": msg})
			return
		}
		c.SetCookie("atlas_rt", tokens.RefreshToken, 7*24*60*60, rt_path, "", false, true)
		c.JSON(http.StatusOK, gin.H{"access_token": tokens.AccessToken})
	}

}

func (ah *AuthHandlers) GetUserSessionsHandler(c *gin.Context) {
	userID := c.MustGet(gin.AuthUserKey).(uuid.UUID)
	sessions, err := ah.sessionService.GetSessionsList(c.Request.Context(), userID)
	if err != nil {
		msg, statusCode := httperror.GetMessageAndStatusCode(err)
		c.JSON(statusCode, gin.H{"detail": msg})
		return
	}
	c.JSON(http.StatusOK, sessions)
}

func (ah *AuthHandlers) DeleteSession(c *gin.Context) {
	sessionID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"detail": "Invalid session id"})
		return
	}
	err = ah.sessionService.DeleteSession(c.Request.Context(), sessionID)
	if err != nil {
		msg, statusCode := httperror.GetMessageAndStatusCode(err)
		c.JSON(statusCode, gin.H{"detail": msg})
		return
	}
	c.String(http.StatusNoContent, "")
}

func (ah *AuthHandlers) NewDeleteCurrentSession(rt_path string) gin.HandlerFunc {
	return func(c *gin.Context) {
		refreshToken, err := c.Cookie("atlas_rt")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"detail": "Refresh token is missing or invalid"})
			return
		}
		err = ah.sessionService.DeleteSessionByToken(c.Request.Context(), refreshToken)
		if err != nil {
			msg, statusCode := httperror.GetMessageAndStatusCode(err)
			c.JSON(statusCode, gin.H{"detail": msg})
			return
		}
		c.SetCookie("atlas_rt", "", -1, rt_path, "", false, true)
		c.String(http.StatusNoContent, "")
	}

}
