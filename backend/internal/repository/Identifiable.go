package repository

type Identifiable interface {
	GetID() string
	SetID(str string)
}
