-- Write your PostgreSQL query statement below
SELECT tweet_id
FROM TWEETS
WHERE CHAR_LENGTH(CONTENT) > 15
