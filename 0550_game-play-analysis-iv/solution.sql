SELECT
    ROUND(
        COUNT(DISTINCT IF(DATEDIFF(event_date, first_date) = 1, player_id, NULL)) / COUNT(DISTINCT player_id),
    2) AS fraction
FROM (
SELECT
    player_id,
    event_date,
    MIN(event_date) OVER (PARTITION BY player_id) AS first_date
    FROM Activity
) AS first;