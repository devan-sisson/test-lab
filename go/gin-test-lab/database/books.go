package database

const CreateBooks = `
CREATE TABLE books (
isbn char(14) NOT NULL,
title varchar(255) NOT NULL,
author varchar(255) NOT NULL,
price decimal(5,2) NOT NULL
	);`

const InsertBooks = `
INSERT INTO books (isbn, title, author, price) VALUES
('978-1503261969', 'Emma', 'Jayne Austen', 9.44),
('978-1505255607', 'The Time Machine', 'H. G. Wells', 5.99),
('978-1503379640', 'The Prince', 'Niccolò Machiavelli', 6.99);`

const AlterBooks = `ALTER TABLE books ADD PRIMARY KEY (isbn);`

const GetBooksCount = `SELECT count(*) AS count FROM people;`
