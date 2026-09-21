SELECT
    distinct num ConsecutiveNums
FROM (
         SELECT
             num,
             ROW_NUMBER() OVER (ORDER BY id)
                 - ROW_NUMBER() OVER (PARTITION BY num ORDER BY id) AS grp
         FROM Logs
     ) AS sub
GROUP BY num, grp
HAVING COUNT(*) >= 3;
