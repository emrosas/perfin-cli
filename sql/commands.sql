-- CREATE TABLE IF NOT EXISTS income (
--   id INTEGER PRIMARY KEY AUTOINCREMENT,
--   description TEXT NOT NULL,
--   amount INTEGER NOT NULL
-- );
--
-- CREATE TABLE IF NOT EXISTS expenses (
--   id INTEGER PRIMARY KEY AUTOINCREMENT,
--   description TEXT NOT NULL,
--   amount INTEGER NOT NULL
-- );

-- INSERT INTO income (description, amount) VALUES ('Salary', 50000);

-- ALTER TABLE income ADD COLUMN created_at DATETIME DEFAULT CURRENT_TIMESTAMP;
-- ALTER TABLE expenses ADD COLUMN created_at DATETIME DEFAULT CURRENT_TIMESTAMP;

-- DROP TABLE IF EXISTS income;
-- DROP TABLE IF EXISTS expenses;

-- CREATE TABLE IF NOT EXISTS income (
--   id INTEGER PRIMARY KEY AUTOINCREMENT,
--   description TEXT NOT NULL,
--   amount INTEGER NOT NULL,
--   created_at DATETIME DEFAULT CURRENT_TIMESTAMP
-- );
--
-- CREATE TABLE IF NOT EXISTS expenses (
--   id INTEGER PRIMARY KEY AUTOINCREMENT,
--   description TEXT NOT NULL,
--   amount INTEGER NOT NULL,
--   created_at DATETIME DEFAULT CURRENT_TIMESTAMP
-- );

-- INSERT INTO income (description, amount) VALUES ('Salary', 50000);
-- INSERT INTO income (description, amount) VALUES ('Rent', 10000);
-- INSERT INTO expenses (description, amount) VALUES ('Groceries', 2000);

-- CREATE VIEW total_income AS
--   SELECT SUM(amount) FROM income;
--
-- CREATE VIEW total_expenses AS
--   SELECT SUM(amount) FROM expenses;

-- DROP VIEW IF EXISTS total_balance;

-- CREATE VIEW total_balance AS
--   SELECT (SELECT SUM(amount) FROM income) - (SELECT SUM(amount) FROM expenses) AS balance;

SELECT * FROM total_income;
SELECT * FROM total_expenses;
SELECT * FROM total_balance;
