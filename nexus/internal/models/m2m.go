package models

type WorkspaceUser struct {
	WorkspaceID int `json:"workspace_id"`
	UserID      int `json:"user_id"`
}

type ProjectUser struct {
	ProjectID int `json:"project_id"`
	UserID    int `json:"user_id"`
}

type TaskUser struct {
	TaskID int `json:"task_id"`
	UserID int `json:"user_id"`
}
