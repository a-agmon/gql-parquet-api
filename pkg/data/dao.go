package data

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/a-agmon/gql-parquet-api/pkg/data/sqlhelper"
	"github.com/a-agmon/gql-parquet-api/pkg/model"
	"github.com/samber/lo"
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
func (d *DAO) GetUsersByEmail(email string) ([]*model.User, error) {
	rows, err := d.driver.QueryPreparedStatement(QryUsersByEmail, email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	resultset, err := sqlhelper.ResultSetFromRows(rows)
	if err != nil {
		return nil, err
	}
	users := lo.Map(resultset, func(row map[string]interface{}, index int) *model.User {
		return newUserFromRow(row)
	})
	return users, nil
}
func (s *DAO) GetUsersByEmailDomain(domain string) ([]*model.User, error) {
	formattedDomain := fmt.Sprintf("%%@%s%%", domain)
	rows, err := s.driver.QueryPreparedStatement(QryUsersByDomain, formattedDomain)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	resultset, err := sqlhelper.ResultSetFromRows(rows)
	if err != nil {
		return nil, err
	}
	//log.Printf("Got resultset: %+v", resultset)
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
	user.UserID = sqlhelper.StringOr(row["id"], "None")
	user.AccID = sqlhelper.StringOr(row["account_id"], "None")
	user.Email = sqlhelper.StringOr(row["email"], "None")
	user.Department = sqlhelper.StringOr(row["department"], "None")
	user.CreatedAt = sqlhelper.TimestampOr(row["created_date"], time.Time{})
	return user
}
