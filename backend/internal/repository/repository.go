package repository

type Repository[RecordType Identifiable] interface {
	GetAll() []RecordType
	GetById(id string) (*RecordType, bool)
	Create(item RecordType) RecordType
	Update(id string, record RecordType) (*RecordType, bool)
	Delete(id string) bool
}
