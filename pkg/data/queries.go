package data

const (
	StmntLoadUsersTable = "CREATE TABLE users AS select id, account_id, email, department, " +
		"created_date, first_name, last_name, account_status, active_package_name, account_record_type " +
		"from read_parquet('%s') "
	QryAllUsers      = "SELECT id, account_id, email, department, created_date FROM users limit 100;"
	QryUsersByDomain = "SELECT id, account_id, email, department, created_date FROM users WHERE email LIKE ?"
	QryUsersByEmail  = "SELECT id, account_id, email, department, created_date FROM users WHERE email = ?"
)
