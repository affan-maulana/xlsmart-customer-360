package dto

type CreateOutletRequest struct {
	Name     string `json:"name"`
	Code     string `json:"code"`
	Address  string `json:"address,omitempty"`
	City     string `json:"city,omitempty"`
	Province string `json:"province,omitempty"`
}

type UpdateOutletRequest struct {
	Name     string `json:"name"`
	Code     string `json:"code"`
	Address  string `json:"address,omitempty"`
	City     string `json:"city,omitempty"`
	Province string `json:"province,omitempty"`
}

type OutletResponse struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Code     string `json:"code"`
	Address  string `json:"address,omitempty"`
	City     string `json:"city,omitempty"`
	Province string `json:"province,omitempty"`
	IsActive bool   `json:"is_active"`
}