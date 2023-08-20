package data

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/a-agmon/gql-parquet-api/foundation/model"
)

type DataDriver interface {
	Execute(string) error
	Query(string) (*sql.Rows, error)
	Close() error
}

type Store struct {
	driver DataDriver
}

func NewStore(driver DataDriver) *Store {
	parquetTablePath, ok := os.LookupEnv("PARQUET_TABLE_PATH")
	if !ok {
		panic("PARQUET_TABLE_PATH env var is not set")
	}
	log.Printf("Loading parquet table from %s", parquetTablePath)
	initQuery := fmt.Sprintf(StmntLoadUsersTable, parquetTablePath)
	log.Printf("Executing init query")
	err := driver.Execute(initQuery)
	if err != nil {
		panic(err)
	}
	return &Store{driver: driver}
}

func (s *Store) GetUsers() ([]*model.User, error) {
	rows, err := s.driver.Query(QryAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	resultset, err := ResultSetFromRows(rows)
	if err != nil {
		return nil, err
	}
	users := make([]*model.User, 0)
	for _, row := range resultset {
		user := &model.User{}
		user.UserID = StringOr(row["user_id"], "None")
		user.AccID = StringOr(row["acc_id"], "None")
		user.Email = StringOr(row["email"], "None")
		user.DisplayName = StringOr(row["display_name"], "None")
		user.RoleName = StringOr(row["role_name"], "None")
		users = append(users, user)
	}
	return users, nil
}
