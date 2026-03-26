package model

type Warehouse struct {
	ID       int64  `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	Address  string `json:"address" db:"address"`
	Capacity *int64 `json:"capacity,omitempty" db:"capacity"`
	Active   bool   `json:"active" db:"active"`
}

func (w *Warehouse) IsActive() bool {
	return w.Active
}
func (w *Warehouse) Deactivate() {
	w.Active = false
}

type WarehousePlaceType struct {
	ID     int64  `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	Value  string `json:"value" db:"value"`
	Active bool   `json:"active" db:"active"`
}

func (wpt *WarehousePlaceType) IsActive() bool {
	return wpt.Active
}

func (wpt *WarehousePlaceType) Deactivate() {
	wpt.Active = false
}

type WarehousePlace struct {
	ID                   int64  `json:"id" db:"id"`
	Name                 string `json:"name" db:"name"`
	Active               bool   `json:"active" db:"active"`
	WarehousePlaceTypeID *int64 `json:"warehouse_place_type_id,omitempty" db:"warehouse_place_type_id"`
	WarehouseID          *int64 `json:"warehouse_id,omitempty" db:"warehouse_id"`

	WarehousePlaceType *WarehousePlaceType `json:"warehouse_place_type,omitempty"`
	Warehouse          *Warehouse          `json:"warehouse,omitempty"`
}

func (wp *WarehousePlace) IsActive() bool {
	return wp.Active
}

func (wp *WarehousePlace) Deactivate() {
	wp.Active = false
}
