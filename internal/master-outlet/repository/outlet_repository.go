package repository

import (
	"github.com/your-org/ums-bff-service-customer-360/internal/master-outlet/model"
)

type OutletRepository interface {
	GetAll(limit, offset int) ([]model.Outlet, error)
	GetByID(id int64) (*model.Outlet, error)
	Create(outlet *model.Outlet) error
	Update(outlet *model.Outlet) error
	Delete(id int64) error
}