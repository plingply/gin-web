package models

import (
	"database/sql"
	"fmt"
	"reflect"
	"time"
)

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"deleted_at"`
}

func NewResult(data interface{}, c int, m ...string) *Result {
	r := &Result{Data: data, Code: c}

	if e, ok := data.(error); ok {
		if m == nil {
			r.Msg = e.Error()
		}
	} else {
		r.Msg = "SUCCESS"
	}
	if len(m) > 0 {
		r.Msg = m[0]
	}

	return r
}

type Params map[string]interface{}

func FormatResult(rows *sql.Rows) ([]Params, error) {
	columns, err := rows.Columns()

	if err != nil {
		return nil, err
	}

	length := len(columns)
	result := make([]Params, 0)

	for rows.Next() {
		current := makeResultReceiver(length)
		if err := rows.Scan(current...); err != nil {
			return result, err
		}
		value := make(map[string]interface{})
		for i := 0; i < length; i++ {
			key := columns[i]
			val := *(current[i]).(*interface{})
			if val == nil {
				value[key] = nil
				continue
			}
			vType := reflect.TypeOf(val)
			switch vType.String() {
			case "int":
				value[key] = val.(int)
			case "int32":
				value[key] = val.(int32)
			case "int64":
				value[key] = val.(int64)
			case "string":
				value[key] = val.(string)
			case "time.Time":
				value[key] = val.(time.Time)
			case "[]uint8":
				value[key] = string(val.([]uint8))
			default:
				fmt.Printf("unsupport data type '%s' now\n", vType)
			}
		}
		result = append(result, value)
	}
	return result, err
}

func makeResultReceiver(length int) []interface{} {
	result := make([]interface{}, 0, length)
	for i := 0; i < length; i++ {
		var current interface{}
		current = struct{}{}
		result = append(result, &current)
	}
	return result
}
