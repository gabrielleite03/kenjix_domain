package model

type Supplier struct {
	ID           int64   `json:"id" db:"id"`
	RazaoSocial  string  `json:"razaoSocial" db:"razao_social"`
	NomeFantasia string  `json:"nomeFantasia" db:"nome_fantasia"`
	CNPJ         string  `json:"cnpj" db:"cnpj"`
	IE           *string `json:"ie,omitempty" db:"ie"`
	Address      *string `json:"address,omitempty" db:"address"`
	Salesperson  *string `json:"salesperson,omitempty" db:"sales_person"`
	Email        *string `json:"email,omitempty" db:"email"`
	Phone        *string `json:"phone,omitempty" db:"phone"`
	Active       bool    `json:"active" db:"active"`
	CategoryID   *int64  `json:"category_id,omitempty" db:"category_id"`

	Category *Category `json:"category"`
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
