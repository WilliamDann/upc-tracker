package model

import (
	"github.com/WilliamDann/upc-tracker/backend/pkg/upc"
)

type Product struct {
	ID    int64  `json:"id"`
	UPC   string `json:"upc"`
	Name  string `json:"name"`
	Descr string `json:"descr"`
}

func (p *Product) GetID() int64 {
	return p.ID
}
func (p *Product) SetID(id int64) {
	p.ID = id
}

// returns true if the product is valid
func (p *Product) Validate() bool {
	return p.Name != "" && upc.IsUPC(p.UPC)
}
