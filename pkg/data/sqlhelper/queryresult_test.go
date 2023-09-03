package sqlhelper

import (
	"database/sql"
	"testing"

	_ "github.com/marcboeker/go-duckdb"
	"github.com/stretchr/testify/assert"
)

func TestStringOr(t *testing.T) {
	// Test with non-nil string value

	def := "somevalue"
	result := StringOr(nil, def)
	if result != def {
		t.Errorf("StringOr returned incorrect result: got %v, want %v", result, def)
	}

	result = StringOr(4, def)
	if result != def {
		t.Errorf("StringOr returned incorrect result: got %v, want %v", result, def)
	}

	// Test with non-string value
	result = StringOr(def, "badresult")
	if result != def {
		t.Errorf("StringOr returned incorrect result: got %v, want %v", result, def)
	}
}

func createDuckDBTable(t *testing.T) *sql.DB {
	db, err := sql.Open("duckdb", "")
	if err != nil {
		t.Fatalf("Error opening duckdb: %v", err)
	}
	_, err = db.Exec("CREATE TABLE people ( name VARCHAR, age int4, city VARCHAR)")
	if err != nil {
		t.Fatalf("Error creating table: %v", err)
	}
	// Insert rows
	_, err = db.Exec(`
		INSERT INTO people(name, age, city) VALUES
		('John', 25, 'New York'),
		('Jane', 30, 'Los Angeles'),
		('Doe', 35, 'Chicago'),
		('Smith', 40, 'Houston'),
		('Davis', 45, 'Phoenix')
	`)
	if err != nil {
		t.Fatalf("Error inserting rows: %v", err)
	}
	return db
}

func TestResultSetFromRows(t *testing.T) {

	db := createDuckDBTable(t)
	defer db.Close()
	rows, err := db.Query("SELECT name, age, city FROM people")
	if err != nil {
		t.Fatalf("Error querying table: %v", err)
	}
	// Call the function being tested
	resultset, err := ResultSetFromRows(rows)
	// Check the results
	if err != nil {
		t.Errorf("ResultSetFromRows returned an error: %v", err)
	}
	expected := []map[string]interface{}{
		{"name": "John", "age": int32(25), "city": "New York"},
		{"name": "Jane", "age": int32(30), "city": "Los Angeles"},
		{"name": "Doe", "age": int32(35), "city": "Chicago"},
		{"name": "Smith", "age": int32(40), "city": "Houston"},
		{"name": "Davis", "age": int32(45), "city": "Phoenix"},
	}

	assert.Equal(t, expected, resultset)

}
