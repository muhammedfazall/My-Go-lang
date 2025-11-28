-- DDL: CREATE Database
CREATE DATABASE students_db;

-- DDL: CREATE Table (Connect to students_db first)
CREATE TABLE students (
    student_id SERIAL PRIMARY KEY,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    enrollment_year INT,
    major VARCHAR(100)
);

-- DML: INSERT Records
INSERT INTO students (first_name, last_name, enrollment_year, major) 
VALUES 
('alice', 'smith', 2023, 'Computer Science'),
('Bob', 'Jones', 2022, 'Mechanical Engineering'),
('Charlie', 'Brown', 2024, 'History'),
('flin', 'jon', 2020, 'phsycology'),
('jesse', 'pinkman', 2020, 'mathematics');

-- DML: UPDATE
UPDATE students SET major = 'Mathematics' WHERE first_name = 'alice';
UPDATE students SET major = 'Mathematics' WHERE first_name = 'jesse';

-- DML: FILTERING (Reading)
SELECT first_name, major FROM students WHERE enrollment_year < 2023;
SELECT first_name, major FROM students WHERE last_name LIKE 'S%';

-- DML: AGGREGATION & ORDERING
SELECT COUNT(*) AS total_students FROM students;
SELECT first_name, enrollment_year FROM students ORDER BY enrollment_year DESC;

-- DML: DELETE
DELETE FROM students WHERE first_name = 'Charlie';

-- DDL: CLEANUP
DROP TABLE students;