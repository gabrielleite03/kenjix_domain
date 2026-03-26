package model

type Category struct {
	ID          int64  `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Active      bool   `json:"active" db:"active"`
}

func (c *Category) IsActive() bool {
	return c.Active
}

func (c *Category) Deactivate() {
	c.Active = false
}
