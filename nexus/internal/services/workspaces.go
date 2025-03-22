package services

import "github.com/eac0de/atlas/nexus/internal/models"

type WorkspacesService struct{}

func NewWorkspacesService() *WorkspacesService {
	return &WorkspacesService{}
}

func (s *WorkspacesService) GetWorkspacesList() ([]*models.Workspace, error) {
	return []*models.Workspace{}, nil
}

func (s *WorkspacesService) GetWorkspace(id int32) (*models.Workspace, error) {
	return &models.Workspace{}, nil
}

func (s *WorkspacesService) CreateWorkspace(workspace *models.Workspace) error {
	return nil
}

func (s *WorkspacesService) UpdateWorkspace(workspace *models.Workspace) error {
	return nil
}

func (s *WorkspacesService) DeleteWorkspace(id int32) error {
	return nil
}
