// Package usecase provides the use case implementations for user operations.
package usecase

import (
	"github.com/dawit_hopes/grpc_micro_service/internal/domain/models"
	interfaces "github.com/dawit_hopes/grpc_micro_service/pkg/v1"
)

type userUseCase struct {
	repo interfaces.RepoInterface
}

func NewUseCase(repo interfaces.RepoInterface) interfaces.UseCaseInterface {
	return &userUseCase{repo: repo}
}

func (u *userUseCase) Create(user models.User) (*models.User, error) {
	return u.repo.Create(user)
}
func (u *userUseCase) Get(id string) (models.User, error) {
	return u.repo.Get(id)
}
func (u *userUseCase) Update(id string, user models.User) error {
	return u.repo.Update(id, user)
}
func (u *userUseCase) Delete(id string) error {
	return u.repo.Delete(id)
}
