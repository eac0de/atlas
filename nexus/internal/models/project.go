package models

import "database/sql"

type Project struct {
	ID          int            `json:"id"`
	Title       string         `json:"title"`
	Description sql.NullString `json:"description"`
	CreatedAt   string         `json:"created_at"`
	WorkspaceID int            `json:"workspace_id"`
}
