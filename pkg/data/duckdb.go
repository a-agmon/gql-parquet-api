package data

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"github.com/a-agmon/gql-parquet-api/pkg/aws"
	"github.com/marcboeker/go-duckdb"
	"log"
)

type DuckDBDriver struct {
	db *sql.DB
}

func getBootQueries(awsAccessKeyID string, awsSecretAccessKey string, awsSessionToken string, awsRegion string) []string {
	bootQueries := []string{
		"INSTALL json;",
		"LOAD json;",
		"INSTALL parquet;",
		"LOAD parquet;",
		"INSTALL httpfs;",
		"LOAD httpfs;",
	}
	if awsAccessKeyID != "" {
		bootQueries = append(bootQueries, "SET s3_access_key_id='"+awsAccessKeyID+"';")
	}
	if awsSecretAccessKey != "" {
		bootQueries = append(bootQueries, "SET s3_secret_access_key='"+awsSecretAccessKey+"';")
	}
	if awsSessionToken != "" {
		bootQueries = append(bootQueries, "SET s3_session_token='"+awsSessionToken+"';")
	}
	if awsRegion != "" {
		bootQueries = append(bootQueries, "SET s3_region='"+awsRegion+"';")
	}
	return bootQueries
}

func initializeDB(awsAccessKeyID string, awsSecretAccessKey string, awsSessionToken string, awsRegion string) (*sql.DB, error) {
	bootQueries := getBootQueries(awsAccessKeyID, awsSecretAccessKey, awsSessionToken, awsRegion)
	connector, err := duckdb.NewConnector("", func(execer driver.ExecerContext) error {
		for _, qry := range bootQueries {
			_, err := execer.ExecContext(context.Background(), qry, make([]driver.NamedValue, 0))
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	db := sql.OpenDB(connector)
	return db, nil
}

func NewDuckDBDriver(cred aws.AWSCred) *DuckDBDriver {
	db, err := initializeDB(cred.AccessKeyID, cred.SecretAccessKey, cred.SessionToken, cred.Region)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	log.Print("connected to duckdb in memory db")
	return &DuckDBDriver{
		db: db,
	}
}

func (d *DuckDBDriver) Execute(statement string) error {
	log.Printf("Executing statement: \n--\n %s \n---\n", statement)
	_, err := d.db.Exec(statement)
	if err != nil {
		log.Printf("Error executing statement: %v", err)
		return err
	}
	return nil
}

func (d *DuckDBDriver) Query(statement string) (*sql.Rows, error) {
	log.Printf("Executing query: \n--\n %s \n---\n", statement)
	rows, err := d.db.Query(statement)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (d *DuckDBDriver) QueryPreparedStatement(query string, v string) (*sql.Rows, error) {
	log.Printf("Executing statement: \n--\n %s \n---\n", query)
	stmt, err := d.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	log.Printf("Executing statement with args: %+v", v)
	rows, err := stmt.Query(v)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	return rows, nil
}

func (d *DuckDBDriver) Close() error {
	return d.db.Close()
}
