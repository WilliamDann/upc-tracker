package repository

type Repository[RecordType Identifiable] interface {
	GetAll() []RecordType
	GetBy(map[string]any) []RecordType

	Create(RecordType) (*RecordType, error)
	Update(int64, RecordType) (*RecordType, error)
	Delete(int64) error
}
