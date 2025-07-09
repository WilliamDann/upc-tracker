package repository

import "reflect"

// definition
type InMemoryRepo[RecordType Identifiable] struct {
	data map[int64]RecordType
}

// constructor
func NewInMemoryRepo[RecordType Identifiable]() *InMemoryRepo[RecordType] {
	return &InMemoryRepo[RecordType]{map[int64]RecordType{}}
}

// operations
func (r *InMemoryRepo[RecordType]) GetAll() []RecordType {
	objs := []RecordType{}
	for _, val := range r.data {
		objs = append(objs, val)
	}
	return objs
}

func (r *InMemoryRepo[RecordType]) GetBy(filter map[string]any) []RecordType {
	var results []RecordType

	for _, record := range r.data {
		v := reflect.ValueOf(record)
		// If record is a pointer, get the element it points to
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}

		matches := true
		for key, val := range filter {
			fieldVal := v.FieldByName(key)
			if !fieldVal.IsValid() {
				// No such field, treat as no match
				matches = false
				break
			}

			// Compare the field's value with val
			// Use Interface() to get the underlying value
			if !reflect.DeepEqual(fieldVal.Interface(), val) {
				matches = false
				break
			}
		}

		if matches {
			results = append(results, record)
		}
	}

	return results
}

func (r *InMemoryRepo[RecordType]) Create(record RecordType) (*RecordType, error) {
	id := int64(len(r.data))
	record.SetID(id)
	r.data[id] = record
	return &record, nil
}

func (r *InMemoryRepo[RecordType]) Update(id int64, record RecordType) (*RecordType, error) {
	r.data[id] = record
	return &record, nil
}

func (r *InMemoryRepo[RecordType]) Delete(id int64) error {
	delete(r.data, id)
	return nil
}
