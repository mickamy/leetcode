select sub.person_name
from (
    select
        *,
        sum(weight) over (order by turn) total
    from Queue
) sub
where sub.total <= 1000
order by sub.turn desc limit 1
;
