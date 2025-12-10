SELECT
    username,
    toDate(created_at)        AS day,
    count()                   AS total_todos,
    sum(is_completed)         AS completed_todos,
    countIf(is_completed = 0) AS pending_todos
FROM default.todos
GROUP BY
    username,
    day
ORDER BY
    username,
    day
