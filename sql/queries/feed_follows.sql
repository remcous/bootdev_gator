-- name: CreateFeedFollow :one
WITH inserted_feed_follows AS (
    INSERT INTO feed_follows (id,created_at,updated_at,feed_id,user_id)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING *
)
SELECT 
    inserted_feed_follows.*, 
    users.name as user_name, 
    feeds.name AS feed_name 
FROM inserted_feed_follows
INNER JOIN users ON user_id=users.id
INNER JOIN feeds ON feed_id=feeds.id;

-- name: GetFeedFollowsForUser :many
SELECT *, feeds.name AS feed_name, users.name AS user_name 
FROM feed_follows
INNER JOIN users ON user_id=users.id
INNER JOIN feeds ON feed_id=feeds.id
WHERE feed_follows.user_id=$1;