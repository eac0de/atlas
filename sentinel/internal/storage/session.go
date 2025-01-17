package storage

import (
	"context"
	"fmt"
	"net/http"

	"github.com/eac0de/atlas/sentinel/internal/models"
	"github.com/eac0de/gophkeeper/shared/pkg/httperror"

	"github.com/google/uuid"
)

func (storage *AuthStorage) GetSessionsList(ctx context.Context, userID uuid.UUID) ([]*models.Session, error) {
	query := "SELECT id, token, user_id, ip, location, client_info, last_login FROM sessions WHERE user_id=$1"
	rows, err := storage.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	sessionsList := []*models.Session{}
	for rows.Next() {
		if rows.Err() != nil {
			return nil, rows.Err()
		}
		var session models.Session
		err = rows.Scan(&session.ID, &session.Token, &session.UserID, &session.IP, &session.Location, &session.ClientInfo, &session.LastLogin)
		if err != nil {
			return nil, err
		}
		sessionsList = append(sessionsList, &session)
	}
	return sessionsList, nil
}

func (storage *AuthStorage) GetSession(ctx context.Context, sessionID uuid.UUID) (*models.Session, error) {
	query := "SELECT token, user_id, ip, location, client_info, last_login FROM sessions WHERE id=$1"
	row := storage.QueryRow(ctx, query, sessionID)
	session := models.Session{ID: sessionID}
	err := row.Scan(&session.Token, &session.UserID, &session.IP, &session.Location, &session.ClientInfo, &session.LastLogin)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return nil, httperror.New(err, "Session not found", http.StatusNotFound)
		}
		return nil, err
	}
	return &session, nil
}

func (storage *AuthStorage) InsertSession(ctx context.Context, session models.Session) error {
	query := "INSERT INTO sessions (id, token, user_id, ip, location, client_info, last_login) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	_, err := storage.Exec(
		ctx,
		query,
		session.ID,
		session.Token,
		session.UserID,
		session.IP,
		session.Location,
		session.ClientInfo,
		session.LastLogin,
	)
	if err != nil {
		return err
	}
	return nil
}

func (storage *AuthStorage) UpdateSession(ctx context.Context, session models.Session) error {
	if session.ID == uuid.Nil {
		return fmt.Errorf("session id is required for the update")
	}
	query := "UPDATE sessions SET token=$2, user_id=$3, ip=$4, location=$5, client_info=$6, last_login=$7 WHERE id=$1"
	_, err := storage.Exec(
		ctx,
		query,
		session.ID,
		session.Token,
		session.UserID,
		session.IP,
		session.Location,
		session.ClientInfo,
		session.LastLogin,
	)
	if err != nil {
		return err
	}
	return nil
}

func (storage *AuthStorage) DeleteSession(ctx context.Context, sessionID uuid.UUID) error {
	query := "DELETE FROM sessions WHERE id=$1"
	_, err := storage.Exec(
		ctx,
		query,
		sessionID,
	)
	if err != nil {
		return err
	}
	return nil
}
