package repository

import "strconv"

// definition
type InMemoryRepo[RecordType Identifiable] struct {
	data map[string]RecordType
}

// constructor
func NewInMemoryRepo[RecordType Identifiable]() *InMemoryRepo[RecordType] {
	return &InMemoryRepo[RecordType]{map[string]RecordType{}}
}

// operations
func (r *InMemoryRepo[RecordType]) GetAll() []RecordType {
	values := make([]RecordType, len(r.data))
	i := 0
	for _, value := range r.data {
		values[i] = value
		i++
	}
	return values
}

func (r *InMemoryRepo[RecordType]) GetById(id string) (*RecordType, bool) {
	val, ok := r.data[id]
	if !ok {
		return nil, false
	}
	return &val, true
}

func (r *InMemoryRepo[RecordType]) Create(record RecordType) RecordType {
	id := strconv.Itoa(len(r.data))
	record.SetID(id)
	r.data[id] = record
	return record
}

func (r *InMemoryRepo[RecordType]) Update(id string, record RecordType) (*RecordType, bool) {
	_, ok := r.data[id]
	if !ok {
		return nil, false
	}
	record.SetID(id)
	r.data[id] = record
	return &record, true
}

func (r *InMemoryRepo[RecordType]) Delete(id string) bool {
	_, ok := r.data[id]
	if !ok {
		return false
	}

	delete(r.data, id)
	return true
}
