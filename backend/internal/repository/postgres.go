package repository

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
)

// stores information in a postgres table
type PostgresRepo[RecordType Identifiable] struct {
	table string
	db    *sqlx.DB
}

// constructor
func NewPostgresRepo[RecordType Identifiable](db *sqlx.DB, table string) *PostgresRepo[RecordType] {
	return &PostgresRepo[RecordType]{table, db}
}

// helpers

// generate the WHERE section of a sql query
func Where(filters map[string]any) (string, []any) {
	var whereClauses []string
	var values []interface{}
	i := 1
	for key, val := range filters {
		whereClauses = append(whereClauses, fmt.Sprintf("%s = $"+strconv.Itoa(i), key))
		values = append(values, val)
		i++
	}
	query := ""
	if len(whereClauses) > 0 {
		query = "WHERE " + strings.Join(whereClauses, " AND ")
	}
	return query, values
}

// generate an INSERT query
func Insert(table string, record any) (string, error) {
	v := reflect.ValueOf(record)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return "", fmt.Errorf("record must be a non-nil pointer to a struct")
	}
	v = v.Elem()
	t := v.Type()

	var columns []string
	var placeholders []string

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		tag := field.Tag.Get("json")
		if tag == "" {
			tag = strings.ToLower(field.Name)
		}
		if tag == "id" {
			continue
		}

		columns = append(columns, tag)
		placeholders = append(placeholders, ":"+tag)
	}

	return fmt.Sprintf(
		"insert into %s (%s) values (%s)",
		table,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
	), nil
}

// generate UPDATE query
func Update(table string, record any) (string, []any, error) {
	v := reflect.ValueOf(record)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	t := v.Type()
	var setClauses []string
	var args []any

	used := 1
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		dbTag := field.Tag.Get("json")
		if dbTag == "" || dbTag == "-" || dbTag == "id" {
			continue
		}
		setClauses = append(setClauses, fmt.Sprintf("%s = $"+strconv.Itoa(used), dbTag))
		args = append(args, v.Field(i).Interface())
		used += 1
	}

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $"+strconv.Itoa(used), table, strings.Join(setClauses, ", "))
	return query, args, nil
}

// operations
func (r *PostgresRepo[RecordType]) GetBy(filters map[string]any) []RecordType {
	results := []RecordType{}

	query, values := Where(filters)
	query = fmt.Sprintf("select * from %s %s", r.table, query)

	err := r.db.Select(&results, query, values...)
	if err != nil {
		log.Fatalln(err)
	}

	return results
}

func (r *PostgresRepo[RecordType]) GetAll() []RecordType {
	return r.GetBy(map[string]interface{}{})
}

func (r *PostgresRepo[RecordType]) Create(record RecordType) (*RecordType, error) {
	query, _ := Insert(r.table, record)
	fmt.Println(query)

	res, err := r.db.NamedExec(query, record)
	if err != nil {
		return nil, err
	}

	identifiable, ok := any(&record).(RecordType)
	if ok {
		id, err := res.LastInsertId()
		if err == nil {
			identifiable.SetID(id)
		}
	}

	return &identifiable, nil
}

func (r *PostgresRepo[RecordType]) Update(id int64, record RecordType) (*RecordType, error) {
	query, values, err := Update(r.table, record)
	if err != nil {
		return nil, err
	}

	// add id to values
	values = append(values, id)
	_, err = r.db.Exec(query, values...)

	if err != nil {
		return nil, err
	}
	record.SetID(id)

	return &record, nil
}

func (r *PostgresRepo[RecordType]) Delete(id int64) error {
	query := fmt.Sprintf("delete from %s where id=%v", r.table, id)

	_, err := r.db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
