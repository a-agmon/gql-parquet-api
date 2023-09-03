package sqlhelper

import (
	"database/sql"
	"time"
)

// StringOr returns the string value of v if v is a string, if its nil or if it cannot be casted to string it returns def
func StringOr(v interface{}, def string) string {
	if v == nil {
		return def
	}
	if s, ok := v.(string); ok {
		return s
	}
	return def
}

func IntOr(v interface{}, def int) int {
	if v == nil {
		return def
	}
	if s, ok := v.(int); ok {
		return s
	}
	return def
}

func TimestampOr(v interface{}, def time.Time) time.Time {
	if v == nil {
		return def
	}
	if s, ok := v.(time.Time); ok {
		return s
	}
	return def
}

func ResultSetFromRows(rows *sql.Rows) ([]map[string]interface{}, error) {
	resultset := make([]map[string]interface{}, 0)
	cols, _ := rows.Columns()
	columns := make([]interface{}, len(cols))
	columnPointers := make([]interface{}, len(cols))
	for i := range columns {
		columnPointers[i] = &columns[i]
	}
	for rows.Next() {
		row := make(map[string]interface{}, len(cols))
		if err := rows.Scan(columnPointers...); err != nil {
			return nil, err
		}
		// Create our map, and retrieve the value for each column from the pointers slice,
		// storing it in the map with the name of the column as the key.
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			row[colName] = *val
		}
		resultset = append(resultset, row)
	}
	return resultset, nil
}
