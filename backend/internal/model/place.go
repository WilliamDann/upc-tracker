package model

type Place struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
	City    string `json:"city"`
	State   string `json:"state"`
}

func (p *Place) GetID() int64 {
	return p.ID
}
func (p *Place) SetID(id int64) {
	p.ID = id
}
