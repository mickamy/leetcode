WITH daily_stats AS (
    SELECT
        visited_on,
        SUM(amount) AS amount
    FROM Customer
    GROUP BY visited_on
),
     calculated AS (
         SELECT
             visited_on,
             SUM(amount) OVER (
                 ORDER BY visited_on
                 ROWS BETWEEN 6 PRECEDING AND CURRENT ROW
                 ) AS amount,
             ROUND(
                     AVG(amount) OVER (
                         ORDER BY visited_on
                         ROWS BETWEEN 6 PRECEDING AND CURRENT ROW
                         ), 2
             ) AS average_amount,
             MIN(visited_on) OVER () AS first_date
         FROM daily_stats
     )
SELECT
    visited_on,
    amount,
    average_amount
FROM calculated
WHERE DATEDIFF(visited_on, first_date) >= 6
ORDER BY visited_on
;
