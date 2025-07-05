package model

type Product struct {
	ID   string `json:"id"`
	UPC  string `json:"upc"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

func (p *Product) GetID() string {
	return p.ID
}
func (p *Product) SetID(id string) {
	p.ID = id
}

// returns true if the product is valid
//
//	TODO parsing UPC code?
func (p *Product) Validate() bool {
	return p.Name != "" && p.UPC != ""
}
