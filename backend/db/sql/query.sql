-- name: GetEvent :one
SELECT *
FROM events
WHERE event_id = $1;

-- name: ListEventsByTypeAndTimeRange :many
SELECT *
FROM events
WHERE event_type = $1
  AND event_time BETWEEN $2 AND $3
ORDER BY event_time DESC;

-- name: GetFeedingEventWithDetails :one
SELECT
  e.*,
  f.*
FROM
  events e
JOIN
  feeding_events f ON e.event_id = f.event_id
WHERE
  e.event_type = 'Feeding' AND e.event_id = $1;

-- name: ListDiaperEventsWithPoop :many
SELECT
e.*,
d.*
FROM
events e
JOIN
diaper_events d ON e.event_id = d.event_id
WHERE
e.event_type = 'Diaper' AND d.poop = TRUE
ORDER BY e.event_time DESC;

-- name: CreateEvent :one
INSERT INTO events (event_type, event_time, notes, mood)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: CreateFeedingEvent :execresult
INSERT INTO feeding_events (event_id, amount, feed_type, spitup, start_time, end_time, notes)
VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: UpdateEventNotes :exec
UPDATE events
SET notes = $2
WHERE event_id = $1;


-- name: DeleteEvent :exec
DELETE FROM events
WHERE event_id = $1;
