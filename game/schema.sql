/* Create the hangman database and tables.*/

CREATE DATABASE hangman;

CREATE TABLE games (
    uuid    varchar(100) CONSTRAINT uuid PRIMARY KEY,
    turns_left integer NOT NULL,
    word    varchar(100) NOT NULL,
    used    varchar(100),
    available_hints integer NOT NULL
)
