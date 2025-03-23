package models

type Workspace struct {
	ID          int32   `json:"id"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	CreatedAt   string  `json:"created_at"`
	OwnerID     int32   `json:"owner_id"`
}
