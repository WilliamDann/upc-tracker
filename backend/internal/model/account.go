package model

import "encoding/json"

type Account struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"` // Hidden by custom marshall
	Name     string `json:"name"`
}

// omit Password field when account is Marshalled
func (a Account) MarshalJSON() ([]byte, error) {
	// Create an alias to avoid infinite recursion
	type Alias Account
	return json.Marshal(&struct {
		Password *string `json:"password,omitempty"` // Explicitly omit password
		*Alias
	}{
		Password: nil, // Always hide password
		Alias:    (*Alias)(&a),
	})
}

func (a *Account) GetID() int64 {
	return a.ID
}
func (a *Account) SetID(id int64) {
	a.ID = id
}
