package model

type Supplier struct {
	ID         int64   `json:"id" db:"id"`
	Name       string  `json:"name" db:"name"`
	CNPJ       string  `json:"cnpj" db:"cnpj"`
	Country    *string `json:"country,omitempty" db:"country"`
	Email      *string `json:"email,omitempty" db:"email"`
	Phone      *string `json:"phone,omitempty" db:"phone"`
	Seller     *string `json:"seller,omitempty" db:"seller"`
	SellerFone *string `json:"seller_fone,omitempty" db:"seller_fone"`
	Active     bool    `json:"active" db:"active"`
	CategoryID *int64  `json:"category_id,omitempty" db:"category_id"`

	Category *Category `json:"category,omitempty"`
}

func (s *Supplier) IsActive() bool {
	return s.Active
}

func (s *Supplier) Deactivate() {
	s.Active = false
}

func (s *Supplier) HasCNPJ() bool {
	return s.CNPJ != ""
}
