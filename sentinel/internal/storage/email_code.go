package storage

import (
	"context"
	"fmt"
	"net/http"

	"github.com/eac0de/atlas/sentinel/internal/models"
	"github.com/eac0de/gophkeeper/shared/pkg/httperror"

	"github.com/google/uuid"
)

func (storage *AuthStorage) GetEmailCodeByID(ctx context.Context, emailCodeID uuid.UUID) (*models.EmailCode, error) {
	query := "SELECT id, email, code, expires_at, number_of_attempts FROM email_codes WHERE id=$1"
	row := storage.QueryRow(ctx, query, emailCodeID)
	emailCode := models.EmailCode{}
	err := row.Scan(
		&emailCode.ID,
		&emailCode.Email,
		&emailCode.Code,
		&emailCode.ExpiresAt,
		&emailCode.NumberOfAttempts,
	)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return nil, httperror.New(err, "EmailCode not found", http.StatusNotFound)
		}
		println(err.Error())
		return nil, err
	}
	return &emailCode, nil
}

func (storage *AuthStorage) UpdateEmailCode(ctx context.Context, emailCode *models.EmailCode) error {
	if emailCode.ID == uuid.Nil {
		return fmt.Errorf("failed to update emailCode: ID cannot be empty")
	}
	query := "UPDATE email_codes SET email=$2, code=$3, expires_at=$4, number_of_attempts=$5 WHERE id=$1"
	_, err := storage.Exec(
		ctx,
		query,
		emailCode.ID,
		emailCode.Email,
		emailCode.Code,
		emailCode.ExpiresAt,
		emailCode.NumberOfAttempts,
	)
	if err != nil {
		return err
	}
	return nil
}

func (storage *AuthStorage) DeleteEmailCode(ctx context.Context, emailCodeID uuid.UUID) error {
	query := "DELETE FROM email_codes WHERE id=$1"
	_, err := storage.Exec(ctx, query, emailCodeID)
	if err != nil {
		return err
	}
	return nil
}

func (storage *AuthStorage) InsertEmailCode(ctx context.Context, emailCode *models.EmailCode) error {
	query := "INSERT INTO email_codes (id, email, code, expires_at, number_of_attempts) VALUES($1,$2,$3,$4,$5)"
	_, err := storage.Exec(
		ctx,
		query,
		emailCode.ID,
		emailCode.Email,
		emailCode.Code,
		emailCode.ExpiresAt,
		emailCode.NumberOfAttempts,
	)
	if err != nil {
		return err
	}
	return nil
}
