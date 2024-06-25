package database


const CreatePeople = `
CREATE TABLE people (
id bigserioal NOT NULL,
first_name varchar(255) NOT NULL,
last_name varchar(255) NOT NULL,
age int NOT NULL
	);`

const InsertPeople = `
INSERT INTO people (first_name, last_name, age) VALUES
('Devan', 'Sisson', 34)
('Alyssa', 'Bateman', 32)
('Brooklyn', 'Sisson', 7);`

const AlterPeople = `ALTER TABLE books ADD PRIMARY KEY (id);`

const GetPeopleCount = `SELECT count(*) AS count FROM people;`
