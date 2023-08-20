package data

const (
	StmntLoadUsersTable = "CREATE TABLE users AS select user_id, acc_id, email, display_name, role_name from read_parquet('%s') "
	QryAllUsers         = "SELECT user_id, acc_id, email, display_name, role_name FROM users;"
)
