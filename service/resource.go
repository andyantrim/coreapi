package service

import (
	"github.com/andyantrim/apiutil/messages"
	"github.com/andyantrim/coreapi/entity"
)

type ResourceInterface interface {
	List() (*entity.Resources, error)
	Add(*entity.Resource) error
	Get(id int64) (*entity.Resource, error)
}

type ResourceService struct {
	repo ResourceInterface
}

func NewResourceService(repo ResourceInterface) *ResourceService {
	return &ResourceService{
		repo: repo,
	}
}

func (r *ResourceService) List() (*entity.Resources, error) {
	return r.repo.List()
}

func (r *ResourceService) Add(obj *entity.Resource) error {
	if obj.Base.IsEmpty() {
		return messages.NewValidationError("cannot supply id on create request")
	}
	err := r.repo.Add(obj)
	if err != nil {
		return messages.NewServerError(err.Error())
	}
	obj, err = r.repo.Get(obj.ID)
	if err != nil {
		return messages.NewServerError(err.Error())
	}

	return nil
}

func (r *ResourceService) Get(id int64) (*entity.Resource, error) {
	return r.repo.Get(id)
}
