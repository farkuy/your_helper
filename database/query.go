package database

const (
	createUserTable = `CREATE TABLE IF NOT EXISTS users
	(
    	id SERIAL PRIMARY KEY,
    	location TEXT
	)`
)
