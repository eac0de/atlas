package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/eac0de/atlas/sentinel/internal/api/handlers"
	"github.com/eac0de/atlas/sentinel/internal/api/inmiddlewares"
	"github.com/eac0de/atlas/sentinel/internal/config"
	"github.com/eac0de/atlas/sentinel/internal/grpcserver"
	"github.com/eac0de/atlas/sentinel/internal/services"
	"github.com/eac0de/atlas/sentinel/internal/storage"
	"github.com/eac0de/gophkeeper/shared/pkg/emailsender"

	"github.com/gin-gonic/gin"
)

func setupRouter(
	sessionService *services.SessionService,
	authService *services.AuthService,
) *gin.Engine {
	router := gin.Default()
	rootGroup := router.Group("api/auth")

	authHandlers := handlers.NewAuthHandlers(
		authService,
		sessionService,
	)

	rootGroup.POST("code/generate/", authHandlers.GenerateEmailCodeHandler)
	rootGroup.POST("code/verify/", authHandlers.NewVerifyEmailCodeHandler("/api/auth/token/"))
	rootGroup.POST("token/", authHandlers.NewRefreshTokenHandler("/api/auth/token/"))
	rootGroup.DELETE("token/", authHandlers.NewDeleteCurrentSession("/api/auth/token/"))

	authenticatedGroup := rootGroup.Group("/", inmiddlewares.NewAuthMiddleware(sessionService))
	authenticatedGroup.GET("/sessions/", authHandlers.GetUserSessionsHandler)
	authenticatedGroup.DELETE("/sessions/:id/", authHandlers.DeleteSession)
	return router
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg := config.MustLoad()

	if !cfg.IsDev {
		gin.SetMode(gin.ReleaseMode)
	}

	authStorage, err := storage.NewAuthStorage(
		ctx,
		cfg.PSQLHost,
		cfg.PSQLPort,
		cfg.PSQLUsername,
		cfg.PSQLPassword,
		cfg.PSQLDBName,
	)
	if err != nil {
		panic(err)
	}
	err = authStorage.Migrate(ctx, "./migrations", false)
	if err != nil {
		panic(err)
	}
	defer authStorage.Close()

	var emailSender emailsender.IEmailSender
	if cfg.IsDev {
		emailSender = emailsender.NewMock()
	} else {
		emailSender = emailsender.New(cfg.SMTPHost, cfg.SMTPPort, cfg.SMTPUsername, cfg.SMTPPassword)
	}

	sessionService := services.NewSessionService(cfg.JWTSecretKey, cfg.JWTAccessExp, cfg.JWTRefreshExp, authStorage)
	authService := services.NewAuthService(authStorage, emailSender)

	gprcAuthServer := grpcserver.NewAuthGRPCServer(cfg.GPRCServerAddress, sessionService)
	go gprcAuthServer.Run()

	r := setupRouter(sessionService, authService)
	go r.Run(cfg.ServerAddress)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)
	<-sigChan
}
