package service

import (
	"github.com/your-org/ums-bff-service-customer-360/internal/master-outlet/model"
	"github.com/your-org/ums-bff-service-customer-360/internal/master-outlet/repository"
)

type OutletService interface {
	GetAll(limit, offset int) ([]model.Outlet, error)
	GetByID(id int64) (*model.Outlet, error)
	Create(outlet *model.Outlet) error
	Update(outlet *model.Outlet) error
	Delete(id int64) error
}

type outletServiceImpl struct {
	repo repository.OutletRepository
}

func NewOutletService(repo repository.OutletRepository) OutletService {
	return &outletServiceImpl{repo: repo}
}

func (s *outletServiceImpl) GetAll(limit, offset int) ([]model.Outlet, error) {
	return s.repo.GetAll(limit, offset)
}

func (s *outletServiceImpl) GetByID(id int64) (*model.Outlet, error) {
	return s.repo.GetByID(id)
}

func (s *outletServiceImpl) Create(outlet *model.Outlet) error {
	return s.repo.Create(outlet)
}

func (s *outletServiceImpl) Update(outlet *model.Outlet) error {
	return s.repo.Update(outlet)
}

func (s *outletServiceImpl) Delete(id int64) error {
	return s.repo.Delete(id)
}