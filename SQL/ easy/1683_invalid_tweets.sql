SELECT
    tweet_id
FROM
    Tweets
WHERE
    CHARACTER_LENGTH(content) > 15;
