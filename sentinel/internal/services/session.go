package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/eac0de/atlas/sentinel/internal/models"
	"github.com/eac0de/gophkeeper/shared/pkg/httperror"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/mssola/user_agent"
)

type ISessionStore interface {
	GetSession(ctx context.Context, sessionID uuid.UUID) (*models.Session, error)
	InsertSession(ctx context.Context, session models.Session) error
	UpdateSession(ctx context.Context, session models.Session) error
	DeleteSession(ctx context.Context, sessionID uuid.UUID) error
	GetSessionsList(ctx context.Context, userID uuid.UUID) ([]*models.Session, error)
}

type SessionService struct {
	SecretKey    string
	AccessExp    time.Duration
	RefreshExp   time.Duration
	sessionStore ISessionStore
}

type Claims struct {
	jwt.RegisteredClaims
	UserID    uuid.UUID
	SessionID uuid.UUID
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewSessionService(
	secretKey string,
	accessExp time.Duration,
	refreshExp time.Duration,
	sessionStore ISessionStore,
) *SessionService {
	return &SessionService{
		SecretKey:    secretKey,
		AccessExp:    accessExp,
		RefreshExp:   refreshExp,
		sessionStore: sessionStore,
	}
}

func (s *SessionService) CreateSession(ctx context.Context, userID uuid.UUID, userAgent string, ip string) (*Tokens, error) {
	sessionID := uuid.New()
	accessToken, err := s.createToken(userID, sessionID, s.AccessExp)
	if err != nil {
		return nil, err
	}
	refreshToken, err := s.createToken(userID, sessionID, s.RefreshExp)
	if err != nil {
		return nil, err
	}
	session := models.Session{
		ID:         sessionID,
		Token:      refreshToken,
		IP:         net.ParseIP(ip),
		Location:   s.getLocation(ip),
		ClientInfo: s.getClientInfo(userAgent),
		LastLogin:  time.Now(),
		UserID:     userID,
	}
	err = s.sessionStore.InsertSession(ctx, session)
	if err != nil {
		return nil, err
	}
	return &Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *SessionService) UpdateSession(ctx context.Context, token string, userAgent string, ip string) (*Tokens, error) {
	claims, err := s.ParseToken(token)
	if err != nil {
		return nil, httperror.New(err, "Invalid token", http.StatusBadRequest)
	}
	session, err := s.sessionStore.GetSession(ctx, claims.SessionID)
	if err != nil {
		return nil, err
	}
	accessToken, err := s.createToken(session.UserID, session.ID, s.AccessExp)
	if err != nil {
		return nil, err
	}
	refreshToken, err := s.createToken(session.UserID, session.ID, s.RefreshExp)
	if err != nil {
		return nil, err
	}

	session.Token = refreshToken
	session.IP = net.ParseIP(ip)
	session.Location = s.getLocation(ip)
	session.ClientInfo = s.getClientInfo(userAgent)
	session.LastLogin = time.Now()

	err = s.sessionStore.UpdateSession(ctx, *session)
	if err != nil {
		return nil, err
	}
	return &Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *SessionService) GetSessionsList(ctx context.Context, userID uuid.UUID) ([]*models.Session, error) {
	return s.sessionStore.GetSessionsList(ctx, userID)
}

func (s *SessionService) DeleteSession(ctx context.Context, sessionID uuid.UUID) error {
	return s.sessionStore.DeleteSession(ctx, sessionID)
}

func (s *SessionService) DeleteSessionByToken(ctx context.Context, token string) error {
	claims, err := s.ParseToken(token)
	if err != nil {
		return err
	}
	return s.sessionStore.DeleteSession(ctx, claims.SessionID)
}

func (s *SessionService) createToken(userID uuid.UUID, sessionID uuid.UUID, exp time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(exp)),
		},
		UserID:    userID,
		SessionID: sessionID,
	})
	tokenString, err := token.SignedString([]byte(s.SecretKey))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}
	return tokenString, nil
}

func (s *SessionService) ParseToken(token string) (*Claims, error) {
	t, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.SecretKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}
	claims, ok := t.Claims.(*Claims)
	if !ok || !t.Valid {
		return nil, fmt.Errorf("invalid token claims")
	}
	if claims.ExpiresAt.Unix() < time.Now().Unix() {
		return nil, fmt.Errorf("token expired")
	}
	return claims, nil
}

func (s *SessionService) getClientInfo(userAgent string) string {
	const defaultClientInfo = "Unknown Client"

	if userAgent == "" {
		return defaultClientInfo
	}
	ua := user_agent.New(userAgent)
	browserName, _ := ua.Browser()
	platform := ""
	if ua.Platform() != "" {
		platform = fmt.Sprintf(", %s", ua.Platform())
	}
	os := ""
	if ua.OSInfo().FullName != "" {
		os = fmt.Sprintf(", %s", ua.OSInfo().FullName)
	}
	clientInfo := fmt.Sprintf("%s%s%s", browserName, platform, os)
	if clientInfo != "" {
		return clientInfo
	}
	return defaultClientInfo
}

func (s *SessionService) getLocation(ip string) string {
	const defaultLocation = "Unknown Location"

	var responseData struct {
		Country string `json:"country"`
		City    string `json:"city"`
		Message string `json:"message"`
	}

	resp, err := http.Get(fmt.Sprintf("http://ip-api.com/json/%s", ip))
	if err != nil {
		return defaultLocation
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		return defaultLocation
	}
	if responseData.Message == "private range" {
		return "Private Range"
	}
	location := fmt.Sprintf("%s%s", responseData.City, responseData.Country)
	if location == "" {
		return defaultLocation
	}
	return location
}
