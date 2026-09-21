select max(sub.num) as num
from (
    select n.num
    from MyNumbers n
    group by n.num
    having count(n.num) = 1
) sub
;
