package services

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"regexp"
	"time"

	"github.com/eac0de/atlas/sentinel/internal/models"
	"github.com/eac0de/atlas/shared/pkg/httperror"
	"github.com/eac0de/gophkeeper/shared/pkg/emailsender"

	"github.com/google/uuid"
)

var emailRegexp = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

type AuthStore interface {
	InsertEmailCode(ctx context.Context, emailCode *models.EmailCode) error
	UpdateEmailCode(ctx context.Context, emailCode *models.EmailCode) error
	DeleteEmailCode(ctx context.Context, emailCodeID uuid.UUID) error
	GetEmailCodeByID(ctx context.Context, emailCodeID uuid.UUID) (*models.EmailCode, error)

	InsertUser(ctx context.Context, user *models.User) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
}

type AuthService struct {
	AuthStore   AuthStore
	emailSender emailsender.IEmailSender
}

func NewAuthService(authStore AuthStore, emailSender emailsender.IEmailSender) *AuthService {
	return &AuthService{
		AuthStore:   authStore,
		emailSender: emailSender,
	}
}

func (as *AuthService) GenerateEmailCode(
	ctx context.Context,
	email string,
) (*models.EmailCode, error) {
	is_valid := emailRegexp.MatchString(email)
	if !is_valid {
		return nil, httperror.New(nil, "Email is not valid", http.StatusBadRequest)
	}
	randInt := rand.Intn(10000)
	if randInt < 1000 {
		randInt = 10000 - randInt
	}
	emailCode := &models.EmailCode{
		ID:        uuid.New(),
		Email:     email,
		Code:      uint16(randInt),                  // Генерация 4-значного кода
		ExpiresAt: time.Now().Add(15 * time.Minute), // Время действия кода
	}
	err := as.AuthStore.InsertEmailCode(
		ctx,
		emailCode,
	)
	if err != nil {
		return nil, err
	}
	as.emailSender.Send(
		"Вход в приложение auth",
		fmt.Sprintf("Ваш код подтверждения в моб. приложении auth - %v", emailCode.Code),
		emailCode.Email,
	)
	return emailCode, nil
}

func (as *AuthService) VerifyEmailCode(
	ctx context.Context,
	emailCodeID uuid.UUID,
	code uint16,
) (*models.User, bool, error) {
	emailCode, err := as.AuthStore.GetEmailCodeByID(
		ctx,
		emailCodeID,
	)
	if err != nil {
		return nil, false, err
	}
	if emailCode.ExpiresAt.Before(time.Now()) || emailCode.NumberOfAttempts > 2 {
		as.AuthStore.DeleteEmailCode(ctx, emailCode.ID)
		return nil, false, httperror.New(
			nil,
			"Code is gone",
			http.StatusGone,
		)
	}
	if code != emailCode.Code {
		emailCode.NumberOfAttempts += 1
		err = as.AuthStore.UpdateEmailCode(ctx, emailCode)
		if err != nil {
			return nil, false, err
		}
		return nil, false, httperror.New(
			nil,
			"Incorrect code",
			http.StatusPreconditionFailed,
		)
	}
	var isNewUser bool
	user, err := as.AuthStore.GetUserByEmail(ctx, emailCode.Email)
	if err != nil {
		_, statusCode := httperror.GetMessageAndStatusCode(err)
		if statusCode != http.StatusNotFound {
			return nil, false, err
		}
		user, err = as.CreateUser(
			ctx,
			emailCode.Email,
		)
		if err != nil {
			return nil, false, err
		}
		isNewUser = true
	}
	as.AuthStore.DeleteEmailCode(ctx, emailCode.ID)
	return user, isNewUser, nil
}

func (as *AuthService) CreateUser(ctx context.Context, email string) (*models.User, error) {
	user := &models.User{
		ID:        uuid.New(),
		Email:     email,
		CreatedAt: time.Now(),
		IsSuper:   false,
	}
	err := as.AuthStore.InsertUser(
		ctx,
		user,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}
