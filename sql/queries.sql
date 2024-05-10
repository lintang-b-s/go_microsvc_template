-- name: InsertUser :one
INSERT INTO users(
     username, email, dob, gender, password
)VALUES(
    $1, $2, $3, $4, $5
) RETURNING id;

-- name: InsertSession :exec
INSERT INTO sessions(
     ref_token_id, username, refresh_token, expires_at
)VALUES(
    $1, $2, $3, $4
);


-- -- name: InsertLikes :exec
-- INSERT INTO likes(
--     id, user_id, tweets_id
-- ) VALUES(
--     $1, $2, $3
-- );

-- -- name: DeleteLikes :exec
-- DELETE FROM likes
-- WHERE user_id=$1 AND tweets_id=$2;


-- -- name: InsertTweets :exec
-- INSERT INTO tweets(
--     id, user_id, hashtag, content
-- ) VALUES(
--     $1, $2, $3, $4
-- );

-- -- name: UpdateTweets :exec
-- UPDATE tweets
-- SET
--     hashtag=$2,
--     content=$3,
--     updated_time=$4
-- WHERE id=$1;


-- -- name: InsertVideos :exec
-- INSERT INTO videos (
--     id, name, url, length, size
-- ) VALUES (
--     $1, $2, $3, $4, $5
-- );

-- -- name: InsertImages :exec
-- INSERT INTO images(
--     id, name, url, size
-- ) VALUES(
--     $1, $2, $3, $4
-- );

-- name: GetUser :one
SELECT id, username, email, dob, gender, created_time, updated_time 
FROM users
WHERE id=$1;

-- name: GetUserByEmail :one
SELECT id, username, email, password, dob, gender, created_time, updated_time 
FROM users
WHERE email=$1;

-- name: GetSession :one
SELECT id, username, refresh_token, expires_at, created_at
FROM sessions
WHERE id=$1;



-- name: DeleteSession :exec
DELETE FROM sessions
WHERE id=$1;


-- -- name: SelectTweets :many
-- SELECT t.id, t.user_id, t.hashtag, t.content, t.created_time, t.updated_time, v.url, i.url, l.id, u.id, u.username
-- FROM tweets t
-- LEFT JOIN likes l ON l.tweets_id = t.id
-- LEFT JOIN users u ON u.id = t.user_id
-- LEFT JOIN videos v ON t.id = v.tweets_id
-- LEFT JOIN images i ON t.id = i.tweets_id
-- WHERE id = ANY($1::UUID[]);











