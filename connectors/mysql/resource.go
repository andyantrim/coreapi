package mysql

import (
	"time"

	"github.com/andyantrim/apiutil/messages"
	"github.com/andyantrim/coreapi/entity"
)

var (
	R1 = entity.Resource{
		Base: entity.Base{
			CreatedBy: 1,
			CreatedAt: time.Now(),
			ID:        1,
		},
		Name: "First",
		Args: map[string]interface{}{
			"example": 1234,
		},
	}

	R2 = entity.Resource{
		Base: entity.Base{
			CreatedBy: 1,
			CreatedAt: time.Now(),
			ID:        2,
		},
		Name: "Second",
		Args: map[string]interface{}{},
	}

	R3 = entity.Resource{
		Base: entity.Base{
			CreatedBy: 2,
			CreatedAt: time.Now(),
			ID:        3,
		},
		Name: "Second",
		Args: map[string]interface{}{
			"example": 1234,
			"public":  true,
		},
	}
)

type ResourceRepo struct{}

func NewResourceRepo() *ResourceRepo {
	return &ResourceRepo{}
}

func (r *ResourceRepo) List() (*entity.Resources, error) {
	return &entity.Resources{R1, R2, R3}, nil
}

func (r *ResourceRepo) Add(_ *entity.Resource) error {
	return nil
}

func (r *ResourceRepo) Get(id int64) (*entity.Resource, error) {
	switch id {
	case 1:
		return &R1, nil
	case 2:
		return &R2, nil
	case 3:
		return &R3, nil
	default:
		return nil, messages.NewNotFound("not found in DB")
	}
}
