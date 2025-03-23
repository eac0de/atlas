package storage

import (
	"context"
	"fmt"
	"net/http"

	"github.com/eac0de/atlas/nexus/internal/models"
	"github.com/eac0de/atlas/shared/pkg/httperror"
	"github.com/gofrs/uuid"
)

func (storage *NexusStorage) GetWorkspacesList(ctx context.Context, userID uuid.UUID, limit int, offset int) ([]*models.Workspace, error) {
	query := "SELECT id, title, description, created_at, owner_id FROM workspaces WHERE user_id=$1"
	rows, err := storage.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	workspacesList := []*models.Workspace{}
	for rows.Next() {
		if rows.Err() != nil {
			return nil, rows.Err()
		}
		var workspace models.Workspace
		err = rows.Scan(&workspace.ID, &workspace.Title, &workspace.Description, &workspace.CreatedAt, &workspace.OwnerID)
		if err != nil {
			return nil, err
		}
		workspacesList = append(workspacesList, &workspace)
	}
	return workspacesList, nil
}

func (storage *NexusStorage) GetWorkspace(ctx context.Context, WorkspaceID uuid.UUID) (*models.Workspace, error) {
	query := "SELECT title, description, created_at, owner_id FROM Workspaces WHERE id=$1"
	row := storage.QueryRow(ctx, query, WorkspaceID)
	Workspace := models.Workspace{ID: WorkspaceID}
	err := row.Scan(&Workspace.Token, &Workspace.UserID, &Workspace.IP, &Workspace.Location, &Workspace.ClientInfo, &Workspace.LastLogin)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return nil, httperror.New(err, "Workspace not found", http.StatusNotFound)
		}
		return nil, err
	}
	return &Workspace, nil
}

func (storage *NexusStorage) InsertWorkspace(ctx context.Context, Workspace models.Workspace) error {
	query := "INSERT INTO Workspaces (id, token, user_id, ip, location, client_info, last_login) VALUES ($1, $2, $3, $4, $5, $6, $7)"
	_, err := storage.Exec(
		ctx,
		query,
		Workspace.ID,
		Workspace.Token,
		Workspace.UserID,
		Workspace.IP,
		Workspace.Location,
		Workspace.ClientInfo,
		Workspace.LastLogin,
	)
	if err != nil {
		return err
	}
	return nil
}

func (storage *NexusStorage) UpdateWorkspace(ctx context.Context, Workspace models.Workspace) error {
	if Workspace.ID == uuid.Nil {
		return fmt.Errorf("Workspace id is required for the update")
	}
	query := "UPDATE Workspaces SET token=$2, user_id=$3, ip=$4, location=$5, client_info=$6, last_login=$7 WHERE id=$1"
	_, err := storage.Exec(
		ctx,
		query,
		Workspace.ID,
		Workspace.Token,
		Workspace.UserID,
		Workspace.IP,
		Workspace.Location,
		Workspace.ClientInfo,
		Workspace.LastLogin,
	)
	if err != nil {
		return err
	}
	return nil
}

func (storage *NexusStorage) DeleteWorkspace(ctx context.Context, WorkspaceID uuid.UUID) error {
	query := "DELETE FROM Workspaces WHERE id=$1"
	_, err := storage.Exec(
		ctx,
		query,
		WorkspaceID,
	)
	if err != nil {
		return err
	}
	return nil
}
