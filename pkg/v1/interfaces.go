// Package v1 defines the interfaces for the v1 API.
package v1

import "github.com/dawit_hopes/grpc_micro_service/internal/domain/models"

// RepoInterface defines the interface for repository methods.
type RepoInterface interface {
	Create(models.User) error
	Get(id string) (models.User, error)
	Update(models.User) error
	Delete(id string) error
}

// UseCaseInterface defines the interface for usecase methods.
type UseCaseInterface interface {
	Create(models.User) error
	Get(id string) (models.User, error)
	Update(models.User) error
	Delete(id string) error
}
