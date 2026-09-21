select
    s.id,
    coalesce(
            if(
                    s.id % 2 = 0,
                    lag(s.student) over (order by s.id),
                    lead(s.student) over (order by s.id)
            )
    , s.student) student
from Seat s
;
