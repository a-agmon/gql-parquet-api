package data

// user data - user_id, acc_id, email, department, created_at
const (
	StmntLoadUsersTable = "CREATE TABLE users AS select user_id, acc_id, email, department, created_at from read_parquet('%s') "
	QryAllUsers         = "SELECT user_id, acc_id, email, department, created_at FROM users limit 100;"
	QryUsersByDomain    = "SELECT user_id, acc_id, email, department, created_at FROM users WHERE contains(email, ?);"
)
