-- Postgres 15 

---------------------------------------

DROP TABLE Users;

---------------------------------------

CREATE TABLE IF NOT EXISTS Users (
    UserID SERIAL PRIMARY KEY,
    Username TEXT NOT NULL UNIQUE,
    Password TEXT NOT NULL UNIQUE,
); 

---------------------------------------

SELECT * FROM Users

---------------------------------------

SELECT ? FROM Users

---------------------------------------

INSERT INTO Users (Username, Password)
VALUES (?, ?); 

---------------------------------------

UPDATE Users
SET Username = ?, Password = ?
WHERE Username = ?; 

---------------------------------------

DELETE FROM Users 
WHERE Username = ?;

---------------------------------------

