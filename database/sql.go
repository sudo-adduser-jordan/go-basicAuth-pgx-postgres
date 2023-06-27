package database

// pgx v5

var CREATE_USER_TABLE = `
	CREATE TABLE IF NOT EXISTS Users (
		UserID SERIAL,
		Username TEXT NOT NULL UNIQUE,
		Password TEXT NOT NULL UNIQUE,
		PRIMARY KEY(UserID)
	);
`

var DROP_TABLE = `
	DROP TABLE Users;
`

var INSERT_USER = `
	INSERT INTO Users (Username, Password)
	VALUES ($1, $2); 
`

var SELECT_USER = `
	SELECT * 
	FROM users
	WHERE username = $1
	ORDER BY username ASC
	LIMIT 1
`

var UPDATE_USER = `
	UPDATE Users
	SET Username = $1, Password = $2
	WHERE Username = $3; 
`

var DELETE_USER = `
	DELETE FROM Users 
	WHERE Username = $1;
`
