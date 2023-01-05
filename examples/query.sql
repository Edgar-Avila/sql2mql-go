CREATE TABLE students (
    id INT PRIMARY KEY,
    name VARCHAR(100),
    graduated BOOL,
    score INT
);

INSERT INTO students
(name, graduated, score)
VALUES
("Andrew", TRUE, 100),
("Tom", TRUE, 75),
("Tobby", TRUE, 80),
("Johann", FALSE, 77),
("Amadeus", FALSE, 55)

SELECT * FROM students 
ORDER BY score DESC
LIMIT 3;

DROP TABLE students;
