package data

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/a-agmon/gql-parquet-api/pkg/data/sqlhelper"
	"github.com/a-agmon/gql-parquet-api/pkg/model"
)

type DataDriver interface {
	Execute(string) error
	Query(string) (*sql.Rows, error)
	QueryPreparedStatement(string, string) (*sql.Rows, error)
	Close() error
}

type DAO struct {
	driver DataDriver
}

func NewStore(driver DataDriver) *DAO {
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
	return &DAO{driver: driver}
}
func (s *DAO) GetUsersByEmailDomain(domain string) ([]*model.User, error) {
	rows, err := s.driver.QueryPreparedStatement(QryUsersByDomain, domain)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	resultset, err := sqlhelper.ResultSetFromRows(rows)
	if err != nil {
		return nil, err
	}
	log.Printf("Got resultset: %+v", resultset)
	users := make([]*model.User, 0)
	for _, row := range resultset {
		user := newUserFromRow(row)
		users = append(users, user)
	}
	return users, nil
}

func (s *DAO) GetUsers() ([]*model.User, error) {
	rows, err := s.driver.Query(QryAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	resultset, err := sqlhelper.ResultSetFromRows(rows)
	if err != nil {
		return nil, err
	}
	users := make([]*model.User, 0)
	for _, row := range resultset {
		user := newUserFromRow(row)
		users = append(users, user)
	}
	return users, nil
}

// user_id, acc_id, email, department, created_at
func newUserFromRow(row map[string]interface{}) *model.User {
	user := &model.User{}
	user.UserID = sqlhelper.StringOr(row["user_id"], "None")
	user.AccID = sqlhelper.StringOr(row["acc_id"], "None")
	user.Email = sqlhelper.StringOr(row["email"], "None")
	user.Department = sqlhelper.StringOr(row["department"], "None")
	user.CreatedAt = sqlhelper.TimestampOr(row["created_at"], time.Time{})
	return user
}
