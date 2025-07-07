package repository

type Repository[RecordType Identifiable] interface {
	GetAll() []RecordType

	GetMatches(check func(RecordType) bool) []RecordType
	GetMatch(check func(RecordType) bool) (*RecordType, bool)

	GetById(id string) (*RecordType, bool)

	Create(item RecordType) RecordType
	Update(id string, record RecordType) (*RecordType, bool)
	Delete(id string) bool
}
