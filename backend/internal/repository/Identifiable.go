package repository

type Identifiable interface {
	GetID() int64
	SetID(str int64)
}
