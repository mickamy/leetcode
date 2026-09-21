WITH checked_insurance AS (
    SELECT
        tiv_2016,
        COUNT(*) OVER (PARTITION BY tiv_2015) AS count_tiv_2015,
        COUNT(*) OVER (PARTITION BY lat, lon) AS count_lat_lon
    FROM Insurance
)
SELECT
    ROUND(SUM(tiv_2016), 2) AS tiv_2016
FROM checked_insurance
WHERE
    count_tiv_2015 > 1
  AND count_lat_lon = 1
;
