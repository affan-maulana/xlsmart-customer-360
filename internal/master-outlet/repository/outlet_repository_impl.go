package repository

import (
	"database/sql"
	"fmt"

	"github.com/your-org/ums-bff-service-customer-360/internal/master-outlet/model"
)

type OutletRepositoryImpl struct {
	DB *sql.DB
}

func NewOutletRepositoryImpl(db *sql.DB) *OutletRepositoryImpl {
	return &OutletRepositoryImpl{DB: db}
}

func (r *OutletRepositoryImpl) GetAll(limit, offset int) ([]model.Outlet, error) {
	query := "SELECT id, name, code, address, city, province, is_active FROM master_outlet ORDER BY id LIMIT $1 OFFSET $2"
	rows, err := r.DB.Query(query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to query outlets: %w", err)
	}
	defer rows.Close()

	var outlets []model.Outlet
	for rows.Next() {
		var o model.Outlet
		if err := rows.Scan(&o.ID, &o.Name, &o.Code, &o.Address, &o.City, &o.Province, &o.IsActive); err != nil {
			return nil, fmt.Errorf("failed to scan outlet: %w", err)
		}
		outlets = append(outlets, o)
	}
	return outlets, nil
}

func (r *OutletRepositoryImpl) GetByID(id int64) (*model.Outlet, error) {
	query := "SELECT id, name, code, address, city, province, is_active FROM master_outlet WHERE id = $1"
	var o model.Outlet
	err := r.DB.QueryRow(query, id).Scan(&o.ID, &o.Name, &o.Code, &o.Address, &o.City, &o.Province, &o.IsActive)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to query outlet: %w", err)
	}
	return &o, nil
}

func (r *OutletRepositoryImpl) Create(outlet *model.Outlet) error {
	query := `INSERT INTO master_outlet (name, code, address, city, province, is_active)
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	return r.DB.QueryRow(query, outlet.Name, outlet.Code, outlet.Address, outlet.City, outlet.Province, outlet.IsActive).Scan(&outlet.ID)
}

func (r *OutletRepositoryImpl) Update(outlet *model.Outlet) error {
	query := `UPDATE master_outlet SET name=$1, code=$2, address=$3, city=$4, province=$5, is_active=$6 WHERE id=$7`
	_, err := r.DB.Exec(query, outlet.Name, outlet.Code, outlet.Address, outlet.City, outlet.Province, outlet.IsActive, outlet.ID)
	return err
}

func (r *OutletRepositoryImpl) Delete(id int64) error {
	_, err := r.DB.Exec("DELETE FROM master_outlet WHERE id = $1", id)
	return err
}