package models

import (
	"database/sql"
	"time"
)

type Task struct {
	ID          int32          `json:"id"`
	ProjectID   int32          `json:"project_id"`
	Status      int            `json:"status"`
	Title       string         `json:"title"`
	Description sql.NullString `json:"description"`
	Progress    int            `json:"progress"`
	CreatedAt   time.Time      `json:"created_at"`
	DeadlineAt  time.Time      `json:"deadline_at"`
	ParentID    sql.NullInt32  `json:"parent_id"`
}
