-- Write your PostgreSQL query statement below
SELECT name, population, area
FROM World
WHERE area >= 3_000_000 OR population >= 25_000_000
