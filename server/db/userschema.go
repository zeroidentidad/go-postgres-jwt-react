package db

//CreateUsersTable crea la tabla de usuarios sino fue creada con la BD
func CreateUsersTable() {
	DB.Query(`
		CREATE TABLE IF NOT EXISTS users( id serial PRIMARY KEY, name VARCHAR (100) NOT NULL, password VARCHAR (355) NOT NULL, email VARCHAR (355) UNIQUE NOT NULL, created_on TIMESTAMP NOT NULL default current_timestamp,updated_at TIMESTAMP NOT NULL default current_timestamp )`,
	)
}
